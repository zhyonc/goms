package server

import (
	"goms/config"
	"goms/logger"
	"goms/mongodb"
	"log/slog"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
)

type IChild interface {
	HandleTCPConn(conn net.Conn)
	HandleUDPData(data []byte)
}

type BaseServer interface {
	GetDB() *mongo.Database
	StartTCPListen(addr string)
	StartUDPListen(addr string)
	Stop()
}

type baseServer struct {
	dbClient     *mongodb.DBClient
	db           *mongo.Database
	tcpListeners []*net.TCPListener
	udpListener  *net.UDPConn
	IChild
}

func NewBaseServer(loggerConf config.LoggerConfig, dbConf config.DBConfig, child IChild) BaseServer {
	slog.SetDefault(logger.NewLogger(loggerConf.LogLevel))
	var s baseServer
	s.dbClient = mongodb.NewDBClient(dbConf.DBURI, dbConf.DBTimeout)
	s.tcpListeners = make([]*net.TCPListener, 0)
	s.udpListener = nil
	s.db = s.dbClient.SelectDB(dbConf.DBName)
	s.IChild = child
	return &s
}

// GetDB implements BaseServer.
func (s *baseServer) GetDB() *mongo.Database {
	return s.db
}

// StartTCPListen implements BaseServer.
func (s *baseServer) StartTCPListen(addr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		slog.Error("Resolve tcp addr failed", "err", err, "addr", addr)
		return
	}
	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		slog.Error("Failed to listen", "err", err, "port", tcpAddr.Port)
		return
	}
	s.tcpListeners = append(s.tcpListeners, lis)
	slog.Info("TCPListener is starting at " + lis.Addr().String())
	for {
		if lis == nil {
			slog.Debug("TCPListener is empty")
			break
		}
		conn, err := lis.AcceptTCP()
		if err != nil {
			slog.Error("Failed to accept conn ", "err", err)
			continue
		}
		go s.IChild.HandleTCPConn(conn)
	}
}

// StartUDPListen implements BaseServer.
func (s *baseServer) StartUDPListen(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		slog.Error("Resolve udp addr failed", "err", err)
		return
	}
	lis, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		slog.Error("Create UDPListener failed", "err", err)
		return
	}
	s.udpListener = lis
	slog.Info("UDPListener is starting at " + lis.LocalAddr().String())
	for {
		if s.udpListener == nil {
			slog.Debug("UDPListener is empty")
			break
		}
		buf := make([]byte, 2048)
		n, addr, err := s.udpListener.ReadFromUDP(buf)
		if err != nil {
			slog.Error("Read packet length failed", "err", err)
			continue
		}
		udpClientAddr := addr.String()
		data := buf[0:n]
		slog.Info("New client udp packet came from", "addr", udpClientAddr)
		go s.IChild.HandleUDPData(data)
	}
}

// Stop implements BaseServer.
func (s *baseServer) Stop() {
	for _, lis := range s.tcpListeners {
		s.stopTCPListen(lis)
	}
	s.stopUDPListen()
	s.dbClient.Disconnect()
}

func (s *baseServer) stopTCPListen(lis *net.TCPListener) {
	if lis == nil {
		slog.Debug("Listener is empty")
		return
	}
	err := lis.Close()
	if err != nil {
		slog.Error("Failed to close listener", "err", err, "addr", lis.Addr())
		return
	}
	slog.Info("Listener was closed", "addr", lis.Addr())
	lis = nil
}

func (s *baseServer) stopUDPListen() {
	if s.udpListener == nil {
		slog.Debug("Listener is empty")
		return
	}
	err := s.udpListener.Close()
	if err != nil {
		slog.Error("Failed to close listener", "err", err, "addr", s.udpListener.LocalAddr())
		return
	}
	slog.Info("Listener was closed", "addr", s.udpListener.LocalAddr())
	s.udpListener = nil
}
