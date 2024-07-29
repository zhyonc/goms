package listener

import (
	"log/slog"
	"net"
)

type TCPListener struct {
	addr       string
	lis        *net.TCPListener
	handleFunc TCPHandleFunc
}

func NewTCPListener(addr string, handleFunc TCPHandleFunc) Listener {
	l := &TCPListener{
		addr:       addr,
		handleFunc: handleFunc,
	}
	return l
}

// Start implements Listener.
func (l *TCPListener) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", l.addr)
	if err != nil {
		slog.Error("Resolve tcp addr failed", "err", err, "addr", l.addr)
		return
	}
	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		slog.Error("Failed to listen", "err", err, "port", tcpAddr.Port)
		return
	}
	l.lis = lis
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
		go l.handleFunc(conn)
	}
}

// Stop implements Listener.
func (l *TCPListener) Stop() {
	if l.lis == nil {
		slog.Debug("Listener is empty")
		return
	}
	err := l.lis.Close()
	if err != nil {
		slog.Error("Failed to close listener", "err", err, "addr", l.addr)
		return
	}
	slog.Info("Listener was closed", "addr", l.addr)
	l.lis = nil
}
