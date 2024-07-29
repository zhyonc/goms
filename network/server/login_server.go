package server

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"goms/config"
	"goms/maple"
	"goms/mongodb"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/mongodb/model/social"
	"goms/network"
	"goms/network/client"
	"goms/network/server/api"
	"goms/network/server/listener"
	"goms/packet/inpacket"
	"goms/packet/outpacket"
	"goms/util"
	"log/slog"
	"net"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type loginServer struct {
	conf              *config.Config
	dbClient          *mongodb.DBClient
	tcpListener       listener.Listener
	udpListener       listener.Listener
	clientMap         sync.Map
	cancel            context.CancelFunc
	worldServerXORKey []byte
	gaugePx           []uint32
}

func NewLoginServer(conf *config.Config, dbClient *mongodb.DBClient) network.LoginServer {
	s := &loginServer{
		conf:              conf,
		dbClient:          dbClient,
		worldServerXORKey: []byte(conf.WorldServer.UDPXORKey),
		gaugePx:           make([]uint32, len(conf.WorldServer.ChannelPorts)),
	}
	s.tcpListener = listener.NewTCPListener(conf.LoginServer.TCPAddr, s.HandleTCPConn)
	s.udpListener = listener.NewUDPListener(conf.LoginServer.UDPAddr, conf.LoginServer.UDPXORKey, s.HandleUDPMessage)
	return s
}

// Run implements network.LoginServer.
func (s *loginServer) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	var wg sync.WaitGroup
	// Get channel server gauge px from world server
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(300 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				msg := api.NewMessage("", api.QueryGaugePx, nil)
				err := msg.Send(s.conf.WorldServer.Addr, s.worldServerXORKey)
				if err != nil {
					slog.Error("Failed to send message", "err", err)
					continue
				}
				var resp api.QueryGaugePxResponse
				err = json.Unmarshal(msg.Content, &resp)
				if err != nil {
					slog.Error("Failed to decode message", "err", err)
					continue
				}
				s.gaugePx = resp.GaugePx
				slog.Debug("Update gauge px ok")
			case <-ctx.Done():
				return
			}
		}
	}()
	// Start tcp listener
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.tcpListener.Start()
	}()
	// Start udp listener
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.udpListener.Start()
	}()
	wg.Wait()
}

// Stop implements network.LoginServer.
func (s *loginServer) Stop() {
	s.cancel()
	s.tcpListener.Stop()
	s.udpListener.Stop()
}

// KickClient implements network.LoginServer.
func (s *loginServer) KickClient(ip string) {
	temp, ok := s.clientMap.Load(ip)
	if !ok {
		return
	}
	_, ok = temp.(network.LoginClient)
	if !ok {
		return
	}
	s.clientMap.Delete(ip)
}

// HandleTCPConn implements network.LoginServer.
func (s *loginServer) HandleTCPConn(conn net.Conn) {
	slog.Info("New client connected", "addr", conn.RemoteAddr())
	var recvIV [4]byte
	var sendIV [4]byte
	rand.Read(recvIV[:])
	rand.Read(sendIV[:])
	c := client.NewLoginClient(conn, recvIV, sendIV, s)
	go c.RecvPacket()
	_, err := conn.Write(outpacket.NewConnect(recvIV[:], sendIV[:]))
	if err != nil {
		slog.Error("Failed to send connect packet", "addr", conn.RemoteAddr())
		c.Disconnect()
		return
	}
	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	s.KickClient(ip)
	s.clientMap.Store(ip, c)
}

// HandleUDPMessage implements network.LoginServer.
func (s *loginServer) HandleUDPMessage(msg *api.Message) {
	temp, ok := s.clientMap.Load(msg.ClientIP)
	if !ok {
		msg.Status = "No active login client found"
		return
	}
	c, ok := temp.(network.LoginClient)
	if !ok {
		msg.Status = "Failed to assert network.LoginClient"
		return
	}
	switch msg.APICode {
	case api.TestPacket:
		c.SendPacket(msg.Content)
		msg.Status = "TestPacket OK"
	case api.SkipSDOAuth:
		var req api.SkipSDOAuthRequest
		err := json.Unmarshal(msg.Content, &req)
		if err != nil {
			msg.Status = "Bad request"
			return
		}
		if c.GetAccountID() > 0 {
			msg.Status = "Already skipped sdo auth"
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		account := s.dbClient.AccountUsecase.FindAccountByID(ctx, req.AccountID)
		if account == nil {
			msg.Status = "Account does not exist"
			return
		}
		ok := s.NextAuth(c, account)
		if !ok {
			msg.Status = "Auth failed"
			return
		}
		msg.Status = "Skip sdo auth ok"
	case api.KickGameClient:
		s.clientMap.Delete(c.GetClientIP())
		c.Disconnect()
		msg.Status = "Kick login client ok"
	default:
		msg.Status = "Unknown api code"
	}
}

// RecvPermissonRequest implements network.LoginServer.
func (s *loginServer) RecvPermissonRequest(c network.LoginClient, data []byte) {
	in := inpacket.NewPermissionRequest(data)
	if in.Region != uint8(maple.Region) || in.Version != maple.Version {
		c.Disconnect()
	}
	// Dont print anything here, someone could just prank us and spam this packet
	// Use mina sessionCreated if you want to print xxx connected
}

// RecvApplyRSAKey implements network.LoginServer.
func (s *loginServer) RecvApplyRSAKey(c network.LoginClient, data []byte) {
	c.SendPacket(outpacket.NewRSAKey())
}

// RecvApplyHotfix implements network.LoginServer.
func (s *loginServer) RecvApplyHotfix(c network.LoginClient, data []byte) {
	c.SendPacket(outpacket.NewReceiveHotfix())
}

// RecvCheckLoginAuthInfo implements network.LoginServer.
func (s *loginServer) RecvCheckLoginAuthInfo(c network.LoginClient, data []byte) {
	in := inpacket.NewCheckLoginAuthInfo(data)
	mac := util.Bytes2MAC(in.MachineID[0:6])
	slog.Info("New login auth info from game", "username", in.Username, "password", in.Password)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	account := s.dbClient.AccountUsecase.FindAccountByUsername(ctx, in.Username)
	if account == nil {
		if s.conf.LoginServer.IsAutoRegister && s.AutoRegister(in.Username, in.Password, c.GetClientIP(), mac) {
			c.SendPacket(outpacket.NewChooseGender(in.Username))
			return
		}
		c.SendPacket(outpacket.NewCheckPasswordResult(maple.TipUsernameNotFound, nil))
		return
	}
	// Auth password
	// SDO Login use random token as password every time that no need to compare
	// ok := util.ComparePassword(false, account.Password, in.Password)
	// if !ok {
	// 	c.SendPacket(outpacket.NewCheckPasswordResult(tip.IncorrectPassword, nil))
	// 	return
	// }
	account.LoginMAC = mac
	_ = s.NextAuth(c, account)
}

// From RecvCheckLoginAuthInfo
func (s *loginServer) AutoRegister(username, password, ip, mac string) bool {
	done := make(chan bool, 1)
	go s.dbClient.WithTransaction(func(ctx mongo.SessionContext) (any, error) {
		accountID := s.dbClient.CounterUsecase.GetAccountID(ctx)
		account := model.NewAccount(accountID, username, password, ip, mac)
		res := s.dbClient.AccountUsecase.CreateNewAccount(ctx, account)
		if !res {
			done <- false
			return nil, fmt.Errorf("failed to create new account")
		}
		done <- true
		return nil, nil
	})
	return <-done
}

func (s *loginServer) NextAuth(c network.LoginClient, account *model.Account) bool {
	// Auth banned
	if account.IsForeverBanned || account.TempBannedExpireDate.After(time.Now()) {
		c.SendPacket(outpacket.NewCheckPasswordResultBanned(account))
		return false
	}
	// Auth second password
	if account.SecondPassword == "" {
		c.SendPacket(outpacket.NewChooseGender(account.Username))
		return false
	}
	// Update login record
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	s.dbClient.AccountUsecase.UpdateLoginRecord(ctx, account.ID, c.GetClientIP(), account.LoginMAC)
	// Login Success
	c.SetAccountID(account.ID)
	c.SendPacket(outpacket.NewCheckPasswordResult(maple.TipSuccess, account))
	s.sendWorldInformation(c)
	return true
}

// RecvGenderSetRequest implements network.LoginServer.
func (s *loginServer) RecvGenderSetRequest(c network.LoginClient, data []byte) {
	in := inpacket.NewGenderSetRequest(data)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok := s.dbClient.AccountUsecase.UpdateGenderAndSecondPassword(ctx, in.Username, in.SecondPassword, in.Gender)
	c.SendPacket(outpacket.NewGenderSetResult(ok))
	s.sendWorldInformation(c)
}

func (s *loginServer) sendWorldInformation(c network.LoginClient) {
	worldCount := 1
	for i := 0; i < worldCount; i++ {
		out := outpacket.NewWorldInformation(
			s.conf.WorldServer.WorldID,
			s.conf.WorldServer.WorldTag,
			s.conf.WorldServer.ChannelPorts,
			s.conf.WorldServer.OnlineLimitPerChannel,
			s.gaugePx,
			s.conf.WorldServer.WorldBallons,
		)
		c.SendPacket(out)
	}
	c.SendPacket(outpacket.NewWorldInformationEnd())
	c.SendPacket(outpacket.NewRecommendWorldMessage(
		worldCount,
		s.conf.WorldServer.WorldID,
		s.conf.WorldServer.WorldRecommendMsg,
	))
}

// RecvWorldStatusCheck implements network.LoginServer.
func (s *loginServer) RecvWorldStatusCheck(c network.LoginClient, data []byte) {
	c.SendPacket(outpacket.NewWorldStatus(maple.WorldIdle))
}

// RecvSelectWorldRequest implements network.LoginServer.
func (s *loginServer) RecvSelectWorldRequest(c network.LoginClient, data []byte) {
	in := inpacket.NewSelectWorldRequest(data)
	if in.WorldID != uint8(s.conf.WorldServer.WorldID) {
		c.SendPacket(outpacket.NewSelectWorldResultFailed(maple.TipUnknown))
		return
	}
	c.SetWorldID(in.WorldID)
	c.SetChannelIndex(in.ChannelIndex)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Find character order ids
	charOrder := s.dbClient.CharacterOrderUsecase.FindCharacterOrder(ctx, c.GetAccountID(), in.WorldID)
	if charOrder == nil {
		charOrder = model.NewCharacterOrder(c.GetAccountID(), in.WorldID)
		ok := s.dbClient.CharacterOrderUsecase.CreateNewCharacterOrder(ctx, charOrder)
		if !ok {
			c.SendPacket(outpacket.NewSelectWorldResultFailed(maple.TipDBFail))
			return
		}
	}
	charOrderIDs := charOrder.CharacterIDs
	// Find all characters
	dbChars := s.dbClient.CharacterUsecase.FindCharactersByAccountID(ctx, c.GetAccountID(), in.WorldID)
	if dbChars == nil {
		c.SendPacket(outpacket.NewSelectWorldResultFailed(maple.TipServerBusy))
		return
	}
	// Separate characters type
	reservedChars := make([]*character.Character, 0)
	chars := make([]*character.Character, 0)
	invMap := make(map[uint32]*inventory.Inventory)
	renamedCharIDs := make([]uint32, 0)
	burningCharLength := 0
	for _, char := range dbChars {
		if char.IsDeleted {
			continue
		}
		// Find character's inventory
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		temp := s.dbClient.InventoryUsecase.FindModelByID(ctx, char.ID)
		if temp == nil {
			slog.Error("Failed to find inventory", "characterID", char.ID)
			continue
		}
		inv, ok := temp.(*inventory.Inventory)
		if !ok {
			slog.Error("Failed to assert type *inventory.Inventory")
			continue
		}
		invMap[char.ID] = inv
		// Character Statistics
		if char.IsReserved {
			reservedChars = append(reservedChars, char)
		}
		chars = append(chars, char)
		if char.IsRenamed {
			renamedCharIDs = append(renamedCharIDs, char.ID)
		}
		if char.IsBurning {
			burningCharLength++
		}
	}
	c.SendPacket(outpacket.NewSetClientKey())
	c.SendPacket(outpacket.NewSetPhysicalWorldID(s.conf.WorldServer.WorldID))
	c.SendPacket(outpacket.NewSelectWorldResultSuccess(s.conf.WorldServer.WorldID, charOrderIDs, reservedChars, chars, invMap,
		renamedCharIDs, s.conf.WorldServer.RenameCharEventStartDate, s.conf.WorldServer.RenameCharEventEndDate, burningCharLength))
}

// RecvGotoWorldSelect implements network.LoginServer.
func (s *loginServer) RecvGotoWorldSelect(c network.LoginClient, data []byte) {
	c.SetWorldID(255)
	c.SetChannelIndex(255)
	s.sendWorldInformation(c)
}

// RecvCreateCharStep implements network.LoginServer.
func (s *loginServer) RecvCreateCharStep(c network.LoginClient, data []byte) {
	c.SendPacket(outpacket.NewCreateCharStep())
}

// RecvCheckDuplicateIDPacket implements network.LoginServer.
func (s *loginServer) RecvCheckDuplicateIDPacket(c network.LoginClient, data []byte) {
	in := inpacket.NewCheckDuplicateIDPacket(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	isDuplicate := s.dbClient.CharacterUsecase.FindCharacterName(ctx, in.CharacterName)
	c.SendPacket(outpacket.NewCheckDuplicatedIDResult(in.CharacterName, isDuplicate))
}

// RecvNewCharPacket implements network.LoginServer.
func (s *loginServer) RecvNewCharPacket(c network.LoginClient, data []byte) {
	in := inpacket.NewCharPacket(data)
	s.dbClient.WithTransaction(func(ctx mongo.SessionContext) (any, error) {
		characterID := s.dbClient.CounterUsecase.GetCharacterID(ctx)
		charOrder := s.dbClient.CharacterOrderUsecase.FindCharacterOrder(ctx, c.GetAccountID(), uint8(c.GetWorldID()))
		charOrder.CharacterIDs = append(charOrder.CharacterIDs, characterID)
		char := character.NewCharacter(characterID, c.GetAccountID(), c.GetWorldID(), in)
		inv := inventory.NewInventory(characterID, in)
		soc := social.NewSocial(characterID)
		if !s.dbClient.CharacterOrderUsecase.UpdateCharacterOrder(ctx, charOrder) ||
			!s.dbClient.CharacterUsecase.CreateNewModel(ctx, char) ||
			!s.dbClient.InventoryUsecase.CreateNewModel(ctx, inv) ||
			!s.dbClient.SocialUsecase.CreateNewModel(ctx, soc) {
			c.SendPacket(outpacket.NewCreateNewCharacterResult(maple.TipServerBusy, nil, nil))
			return nil, fmt.Errorf("failed to create new char doc")
		}
		c.SendPacket(outpacket.NewCreateNewCharacterResult(maple.TipSuccess, char, inv))
		return nil, nil
	})
}

// RecvChangeCharOrderRequest implements network.LoginServer.
func (s *loginServer) RecvChangeCharOrderRequest(c network.LoginClient, data []byte) {
	in := inpacket.NewChangeCharOrderRequest(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	charOrder := s.dbClient.CharacterOrderUsecase.FindCharacterOrder(ctx, c.GetAccountID(), uint8(c.GetWorldID()))
	charOrder.CharacterIDs = in.CharacterIDs
	ok := s.dbClient.CharacterOrderUsecase.UpdateCharacterOrder(ctx, charOrder)
	if !ok {
		slog.Error("Failed to change char order")
	}
	// No need to reply
}

// RecvDeleteCharPacket implements network.LoginServer.
func (s *loginServer) RecvDeleteCharPacket(c network.LoginClient, data []byte) {
	// DeleteChar -> 24h reserverd time
	in := inpacket.NewDeleteCharPacket(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var res maple.Tip
	account := s.dbClient.AccountUsecase.FindAccountByID(ctx, c.GetAccountID())
	reservedDate := time.Now()
	if account == nil {
		res = maple.TipServerBusy
	} else if account.SecondPassword != in.SecondPassword {
		res = maple.TipDBFail
	} else {
		ok := s.dbClient.CharacterUsecase.ReserveCharacter(ctx, in.CharacterID, reservedDate)
		if !ok {
			res = maple.TipServerBusy
		} else {
			res = maple.TipSuccess
		}
	}
	c.SendPacket(outpacket.NewReservedDeleteCharacterResult(in.CharacterID, res, reservedDate))
}

// RecvReservedDeleteCharacterCancelStep implements network.LoginServer.
func (s *loginServer) RecvReservedDeleteCharacterCancelStep(c network.LoginClient, data []byte) {
	in := inpacket.NewReservedDeleteCharacterCancelStep(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var res maple.Tip
	ok := s.dbClient.CharacterUsecase.RestoreCharacter(ctx, in.CharacterID)
	if !ok {
		res = maple.TipServerBusy
	} else {
		res = maple.TipSuccess
	}
	c.SendPacket(outpacket.NewReservedDeleteCharacterCancelResult(in.CharacterID, res))
}

// RecvReservedDeleteCharacterConfirmStep implements network.LoginServer.
func (s *loginServer) RecvReservedDeleteCharacterConfirmStep(c network.LoginClient, data []byte) {
	// Character can be deleted after 24h reserverd time
	in := inpacket.NewReservedDeleteCharacterConfirmStep(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var res maple.Tip
	account := s.dbClient.AccountUsecase.FindAccountByID(ctx, c.GetAccountID())
	if account == nil {
		res = maple.TipServerBusy
	} else if account.SecondPassword != in.SecondPassword {
		res = maple.TipDBFail
	} else {
		s.dbClient.WithTransaction(func(ctx mongo.SessionContext) (any, error) {
			charOrder := s.dbClient.CharacterOrderUsecase.FindCharacterOrder(ctx, c.GetAccountID(), uint8(c.GetWorldID()))
			for index, charID := range charOrder.CharacterIDs {
				if charID == in.CharacterID {
					if index == len(charOrder.CharacterIDs)-1 {
						charOrder.CharacterIDs = charOrder.CharacterIDs[:index]
					} else {
						charOrder.CharacterIDs = append(charOrder.CharacterIDs[:index], charOrder.CharacterIDs[index+1:]...)
					}
					break
				}
			}
			if !s.dbClient.CharacterOrderUsecase.UpdateCharacterOrder(ctx, charOrder) ||
				!s.dbClient.CharacterUsecase.DeleteCharacter(ctx, in.CharacterID) {
				res = maple.TipServerBusy
				return nil, fmt.Errorf("failed to delete char doc")
			} else {
				res = maple.TipSuccess
				return nil, nil
			}
		})
	}
	c.SendPacket(outpacket.NewDeleteCharacterResult(in.CharacterID, res))
}

// RecvSelectCharacterRequest implements network.LoginServer.
func (s *loginServer) RecvSelectCharacterRequest(c network.LoginClient, data []byte) {
	in := inpacket.NewSelectCharacterRequest(data)
	worldConf := s.conf.WorldServer
	if int(c.GetChannelIndex()) > len(worldConf.ChannelPorts) {
		c.SendPacket(outpacket.NewSelectCharacterResult(maple.TipUnknown, nil, 0, 0))
		return
	}
	// Send authorized ip to world server
	msg := api.NewMessage(c.GetClientIP(), api.PrepareMigrate, nil)
	err := msg.Send(s.conf.WorldServer.Addr, s.worldServerXORKey)
	if err != nil {
		slog.Error("Failed to send api.PrepareMigrate message", "err", err)
		c.SendPacket(outpacket.NewSelectCharacterResult(maple.TipServerBusy, nil, 0, 0))
		return
	}
	// Send select character result
	ip4 := net.ParseIP(maple.NexonIP).To4()
	port := worldConf.ChannelPorts[c.GetChannelIndex()]
	out := outpacket.NewSelectCharacterResult(maple.TipSuccess, ip4, port, in.CharacterID)
	c.SendPacket(out)
}
