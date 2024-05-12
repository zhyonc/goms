package server

import (
	"context"
	"goms/config"
	"goms/mongodb/model/character"
	"goms/mongodb/usecase"
	"goms/packet/inpacket"
	"goms/packet/outpacket"
	"goms/transport"
	"goms/transport/client"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"time"
)

type channelServer struct {
	BaseServer
	conf             *config.ChannelServerConfig
	characterUsecase usecase.CharacterUsecase
	testGameClient   transport.GameClient
	characters       sync.Map
}

func NewChannelServer(conf config.ChannelConfig) transport.ChannelServer {
	s := &channelServer{
		conf: &conf.ChannelServer,
	}
	s.BaseServer = NewBaseServer(conf.Logger, conf.DB, s)
	s.characterUsecase = usecase.NewCharacterUsecase(s.BaseServer.GetDB())
	return s
}

// HandleTCPConn implements IChild.
func (s *channelServer) HandleTCPConn(conn net.Conn) {
	slog.Info("New client connected", "addr", conn.RemoteAddr().String())
	c := client.NewChannelClient(conn, s)
	go c.RecvPacket()
	conn.Write(outpacket.NewConnect())
	if s.conf.UDPAddr != "" {
		s.testGameClient = c
	}
}

// HandleUDPData implements IChild.
func (s *channelServer) HandleUDPData(data []byte) {
	if s.testGameClient == nil {
		slog.Error("Not found active test game client")
		return
	}
	s.testGameClient.SendPacket(data)
}

// Run implements Server.
func (s *channelServer) Run() {
	var wg sync.WaitGroup
	for _, port := range s.conf.Ports {
		addr := net.JoinHostPort(s.conf.IP, strconv.Itoa(port))
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.BaseServer.StartTCPListen(addr)
		}()
	}
	if s.conf.UDPAddr != "" {
		// For testing game packet. Don't set udp addr on production environment!
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.BaseServer.StartUDPListen(s.conf.UDPAddr)
		}()
	}
	wg.Wait()
}

// Stop implements Server.
func (s *channelServer) Stop() {
	s.BaseServer.Stop()
}

// RecvSecurityPacket implements transport.ChannelServer.
func (s *channelServer) RecvSecurityPacket(c transport.ChannelClient) {
	c.SendPacket(outpacket.NewSecurityPacket())
}

// RecvMigrateIn implements transport.ChannelServer.
func (s *channelServer) RecvMigrateIn(c transport.ChannelClient, data []byte) {
	in := inpacket.NewMigrateIn(data)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var char *character.Character
	temp, ok := s.characters.Load(in.CharacterID)
	if ok {
		char, _ = temp.(*character.Character)
	} else {
		char = s.characterUsecase.GetCharacterByID(ctx, in.CharacterID)
	}
	if char == nil {
		c.Disconnect()
		return
	}
	s.characters.Store(in.CharacterID, char)
	c.SendPacket(outpacket.NewOpcodeEncryption())
	// Todo
}
