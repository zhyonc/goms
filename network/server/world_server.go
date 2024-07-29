package server

import (
	"context"
	"encoding/json"
	"goms/config"
	"goms/mongodb"
	"goms/network"
	"goms/network/server/api"
	"goms/network/server/listener"
	"net"
	"strconv"
	"sync"
	"time"
)

type worldServer struct {
	conf              *config.Config
	channelServerList []network.ChannelServer
	udpListener       listener.Listener
	authorizedIPs     sync.Map
	cancel            context.CancelFunc
}

func NewWorldServer(conf *config.Config, dbClient *mongodb.DBClient) network.WorldServer {
	s := &worldServer{
		conf: conf,
	}
	ip, _, _ := net.SplitHostPort(conf.WorldServer.Addr)
	ports := conf.WorldServer.ChannelPorts
	for index, port := range ports {
		addr := net.JoinHostPort(ip, strconv.Itoa(port))
		channelServer := NewChannelServer(s, conf.WorldServer.WorldID, uint8(index), addr, dbClient)
		s.channelServerList = append(s.channelServerList, channelServer)
	}
	s.udpListener = listener.NewUDPListener(conf.WorldServer.Addr, conf.WorldServer.UDPXORKey, s.HandleUDPMessage)
	return s
}

// Run implements network.WorldServer.
func (s *worldServer) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	var wg sync.WaitGroup
	wg.Add(1)
	// Clear expire authorized ip
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				s.authorizedIPs.Range(func(key, value any) bool {
					expireTime, ok := value.(time.Time)
					if !ok || time.Since(expireTime) > 60*time.Second {
						ip, ok := key.(string)
						if ok {
							s.KickClient(ip)
						}
					}
					return true
				})
			case <-ctx.Done():
				return
			}
		}
	}()
	// Run channel server
	for _, channelServer := range s.channelServerList {
		wg.Add(1)
		go func(channelServer network.ChannelServer) {
			defer wg.Done()
			channelServer.Run()
		}(channelServer)
	}
	// Start udp listener
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.udpListener.Start()
	}()
	wg.Wait()
}

// Stop implements network.WorldServer.
func (s *worldServer) Stop() {
	s.cancel()
	for _, channelServer := range s.channelServerList {
		channelServer.Stop()
	}
	s.udpListener.Stop()
}

// KickClient implements network.WorldServer.
func (s *worldServer) KickClient(ip string) {
	s.authorizedIPs.Delete(ip)
}

// HandleUDPMessage implements network.WorldServer.
func (s *worldServer) HandleUDPMessage(msg *api.Message) {
	switch msg.APICode {
	case api.QueryGaugePx:
		gaugePx := make([]uint32, len(s.channelServerList))
		for index, channelServer := range s.channelServerList {
			res := channelServer.GetClientCount() * 64 / s.conf.WorldServer.OnlineLimitPerChannel
			if res < 1 {
				res = 1
			}
			gaugePx[index] = uint32(res)
		}
		resp := api.QueryGaugePxResponse{
			GaugePx: gaugePx,
		}
		msg.Content, _ = json.Marshal(resp)
		return
	case api.PrepareMigrate:
		s.authorizedIPs.Store(msg.ClientIP, time.Now())
	default:
		msg.Status = "Unknown api code"
	}
}

// CheckAuthorizedIP implements network.WorldServer.
func (s *worldServer) CheckAuthorizedIP(ip string) bool {
	_, ok := s.authorizedIPs.Load(ip)
	if ok {
		s.authorizedIPs.Delete(ip)
	}
	return ok
}
