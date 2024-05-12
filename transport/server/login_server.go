package server

import (
	"context"
	"goms/config"
	"goms/maple"
	"goms/maple/tip"
	"goms/maple/world"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
	"goms/mongodb/usecase"
	"goms/packet/inpacket"
	"goms/packet/outpacket"
	"goms/transport"
	"goms/transport/client"
	"goms/util"
	"log/slog"
	"net"
	"sync"
	"time"
)

type loginServer struct {
	BaseServer
	conf             config.LoginServerConfig
	worldList        []*config.WorldConfig
	counterUsecase   usecase.CounterUsecase
	accountUsecase   usecase.AccountUsecase
	characterUsecase usecase.CharacterUsecase
	testGameClient   transport.GameClient
}

func NewLoginServer(conf config.LoginConfig) transport.LoginServer {
	s := &loginServer{}
	s.conf = conf.LoginServer
	s.worldList = conf.WorldList
	s.BaseServer = NewBaseServer(conf.Logger, conf.DB, s)
	s.counterUsecase = usecase.NewCounterUsecase(s.BaseServer.GetDB())
	s.accountUsecase = usecase.NewAccountUsecase(s.BaseServer.GetDB())
	s.characterUsecase = usecase.NewCharacterUsecase(s.BaseServer.GetDB())
	return s
}

func (s *loginServer) GetWorldConf(worldID world.WorldID) *config.WorldConfig {
	for _, world := range s.worldList {
		if world.ID == worldID {
			return world
		}
	}
	return nil
}

// HandleTCPConn implements IChild.
func (s *loginServer) HandleTCPConn(conn net.Conn) {
	slog.Info("New client connected", "addr", conn.RemoteAddr().String())
	c := client.NewLoginClient(conn, s)
	go c.RecvPacket()
	conn.Write(outpacket.NewConnect())
	if s.conf.UDPAddr != "" {
		s.testGameClient = c
	}
}

// HandleUDPData implements IChild.
func (s *loginServer) HandleUDPData(data []byte) {
	if s.testGameClient == nil {
		slog.Error("Not found active test game client")
		return
	}
	s.testGameClient.SendPacket(data)
}

// Run implements Server.
func (s *loginServer) Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.BaseServer.StartTCPListen(s.conf.TCPAddr)
	}()
	if s.conf.UDPAddr != "" {
		// For testing game packet. Don't set udp addr on production environment!
		wg.Add(1)
		go func() {
			wg.Done()
			s.BaseServer.StartUDPListen(s.conf.UDPAddr)
		}()
	}
	wg.Wait()
}

// Stop implements Server.
func (s *loginServer) Stop() {
	done := make(chan bool, 1)
	go func() {
		s.BaseServer.Stop()
		done <- true
	}()
	<-done
}

// RecvSecurityPacket implements transport.LoginServer.
func (s *loginServer) RecvSecurityPacket(c transport.LoginClient) {
	c.SendPacket(outpacket.NewSecurityPacket())
}

// RecvPermissonRequest implements transport.LoginServer.
func (s *loginServer) RecvPermissonRequest(c transport.LoginClient, data []byte) {
	in := inpacket.NewSecurityPacket(data)
	if in.Region != maple.Region || in.Version != maple.Version {
		c.Disconnect()
	}
	// Dont print anything here, someone could just prank us and spam this packet
	// Use mina sessionCreated if you want to print xxx connected
}

// RecvApplyHotfix implements transport.LoginServer.
func (s *loginServer) RecvApplyHotfix(c transport.LoginClient) {
	c.SendPacket(outpacket.NewReceiveHotfix())
}

// RecvCheckLoginAuthInfo implements transport.LoginServer.
func (s *loginServer) RecvCheckLoginAuthInfo(c transport.LoginClient, data []byte) {
	in := inpacket.NewCheckLoginAuthInfo(data)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	account := s.accountUsecase.FindAccountByUsername(ctx, in.Username)
	// Auto register
	if account == nil {
		if s.conf.EnableAutoRegister && s.register(ctx, in) {
			c.SendPacket(outpacket.NewChooseGender())
			return
		}
		c.SendPacket(outpacket.NewLoginAuthFailed(tip.NotRegistered))
		return
	}
	// Auth password
	if !s.comparePassword(account.Password, in.Password) {
		c.SendPacket(outpacket.NewLoginAuthFailed(tip.IncorrectPassword))
		return
	}
	// Auth banned
	if account.IsForeverBanned || account.TempBannedExpireTime.After(time.Now()) {
		c.SendPacket(outpacket.NewLoginAuthBanned(account))
		return
	}
	// Auth second password
	if account.SecondPassword == "" {
		c.SendPacket(outpacket.NewChooseGender())
		return
	}
	// Login Success
	s.accountUsecase.UpdateLoginTime(ctx, account.ID, in.MAC)
	c.SetAccountID(account.ID)
	c.SendPacket(outpacket.NewLoginAuthSuccess(account))
	s.SendWorldInformation(c)
}

func (s *loginServer) register(ctx context.Context, in *inpacket.CheckLoginAuthInfo) bool {
	if s.conf.EnableBcryptPassword {
		bPassword, err := util.Bcrypt(in.Password)
		if err != nil {
			slog.Error("Failed to bcrypt password")
			return false
		}
		in.Password = bPassword
	}
	accountID := s.counterUsecase.GetAccountID(ctx)
	account := model.NewAccount(accountID, in.Username, in.Password, in.MAC)
	return s.accountUsecase.CreateNewAccount(ctx, account)
}

func (s *loginServer) comparePassword(dbPassword, reqPassword string) bool {
	if s.conf.EnableBcryptPassword {
		ok := util.ComparePassword(dbPassword, reqPassword)
		if !ok {
			return false
		}
	} else if dbPassword != reqPassword {
		return false
	}
	return true
}

// RecvGenderSetRequest implements transport.LoginServer.
func (s *loginServer) RecvGenderSetRequest(c transport.LoginClient, data []byte) {
	in := inpacket.NewGenderSetRequest(data)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	isSet := s.accountUsecase.UpdateGenderAndSecondPassword(ctx, in.Username, in.SecondPassword, in.Gender)
	c.SendPacket(outpacket.NewGenderSetResult(isSet))
}

// RecvSelectWorldButton implements transport.LoginServer.
func (s *loginServer) RecvSelectWorldButton(c transport.LoginClient, data []byte) {
	in := inpacket.NewSelectWorldButton(data)
	c.SendPacket(outpacket.NewSelectWorldButton(tip.Success, in.WorldID))
}

// RecvSelectWorldRequest implements transport.LoginServer.
func (s *loginServer) RecvSelectWorldRequest(c transport.LoginClient, data []byte) {
	in := inpacket.NewSelectWorldRequest(data)
	c.SetWorldID(in.WorldID)
	c.SetChannelIndex(in.ChannelIndex)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chars := s.characterUsecase.GetCharactersByAccountID(ctx, c.GetAccountID(), uint8(c.GetWorldID()))
	worldConf := s.GetWorldConf(c.GetWorldID())
	res := tip.Success
	if worldConf == nil {
		res = tip.Unknown
		return
	}
	c.SendPacket(outpacket.NewSelectWorldResult(res, chars, worldConf.ID, worldConf.DisableCreateChar, worldConf.RenameCharEventStartTime, worldConf.RenameCharEventEndTime))
}

// RecvGotoWorldSelect implements transport.LoginServer.
func (s *loginServer) RecvGotoWorldSelect(c transport.LoginClient) {
	c.SetWorldID(255)
	c.SetChannelIndex(255)
}

// SendWorldInformation implements transport.LoginServer.
func (s *loginServer) SendWorldInformation(c transport.LoginClient) {
	for _, worldConf := range s.worldList {
		out := outpacket.NewWorldInformation(
			worldConf.ID,
			worldConf.State,
			worldConf.Desc,
			worldConf.ChannelPorts,
			worldConf.Ballons,
		)
		c.SendPacket(out)
	}
	c.SendPacket(outpacket.NewWorldInformationEnd())
}

// RecvCheckSPWExistPacket implements transport.LoginServer.
func (s *loginServer) RecvCheckSPWExistPacket(c transport.LoginClient) {
	c.SendPacket(outpacket.NewCheckSPWExistResult())
}

// RecvCheckDuplicateIDPacket implements transport.LoginServer.
func (s *loginServer) RecvCheckDuplicateIDPacket(c transport.LoginClient, data []byte) {
	in := inpacket.NewCheckDuplicateIDPacket(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	isDuplicate := s.characterUsecase.FindCharacterName(ctx, in.CharacterName)
	c.SendPacket(outpacket.NewCheckDuplicatedIDResult(in.CharacterName, isDuplicate))
}

// RecvNewCharPacket implements transport.LoginServer.
func (s *loginServer) RecvNewCharPacket(c transport.LoginClient, data []byte) {
	in := inpacket.NewCharPacket(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	characterID := s.counterUsecase.GetCharacterID(ctx)
	char := character.NewCharacter(characterID, c.GetAccountID(), c.GetWorldID(), in)
	ok := s.characterUsecase.CreateNewCharacter(ctx, char)
	if !ok {
		c.SendPacket(outpacket.NewCreateNewCharacterResult(tip.ServerBusy, nil))
		return
	}
	c.SendPacket(outpacket.NewCreateNewCharacterResult(tip.Success, char))
}

// RecvDeleteCharPacket implements transport.LoginServer.
func (s *loginServer) RecvDeleteCharPacket(c transport.LoginClient, data []byte) {
	in := inpacket.NewDeleteCharPacket(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var res tip.Tip
	account := s.accountUsecase.FindAccountByID(ctx, c.GetAccountID())
	if account == nil {
		res = tip.ServerBusy
	} else if account.SecondPassword != in.SecondPassword {
		res = tip.IncorrectSecondPassword
	} else {
		ok := s.characterUsecase.DeleteCharacter(ctx, in.CharacterID)
		if !ok {
			res = tip.ServerBusy
		} else {
			res = tip.Success
		}
	}
	c.SendPacket(outpacket.NewDeleteCharacterResult(in.CharacterID, res))
}

// RecvChangeCharOrderRequest implements transport.LoginServer.
func (s *loginServer) RecvChangeCharOrderRequest(data []byte) {
	_ = inpacket.NewChangeCharOrderRequest(data)
	// Todo
}

// RecvSelectCharacterRequest implements transport.LoginServer.
func (s *loginServer) RecvSelectCharacterRequest(c transport.LoginClient, data []byte) {
	in := inpacket.NewSelectCharacterRequest(data)
	worldConf := s.GetWorldConf(c.GetWorldID())
	if worldConf == nil || int(c.GetChannelIndex()) > len(worldConf.ChannelPorts) {
		c.SendPacket(outpacket.NewSelectCharacterResult(tip.JoinGameUNK, nil, 0, 0, false))
		return
	}
	ip4 := net.ParseIP(worldConf.ChannelIP).To4()
	if ip4 == nil {
		ip4 = []byte{127, 0, 0, 1}
	}
	port := worldConf.ChannelPorts[c.GetChannelIndex()]
	out := outpacket.NewSelectCharacterResult(tip.Success, ip4, port, in.CharacterID, in.IsInvisibleOnline)
	c.SendPacket(out)
}
