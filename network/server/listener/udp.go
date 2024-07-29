package listener

import (
	"encoding/json"
	"goms/network/server/api"
	"goms/util"
	"log/slog"
	"net"
)

type UDPListener struct {
	addr       string
	xorKey     []byte
	handleFunc UDPHandleFunc
	lis        *net.UDPConn
}

func NewUDPListener(addr string, xorKey string, handleFunc UDPHandleFunc) Listener {
	l := &UDPListener{
		addr:       addr,
		xorKey:     []byte(xorKey),
		handleFunc: handleFunc,
	}
	return l
}

// Start implements Listener.
func (l *UDPListener) Start() {
	udpAddr, err := net.ResolveUDPAddr("udp", l.addr)
	if err != nil {
		slog.Error("Resolve udp addr failed", "err", err)
		return
	}
	lis, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		slog.Error("Create UDPListener failed", "err", err)
		return
	}
	l.lis = lis
	slog.Info("UDPListener is starting at " + lis.LocalAddr().String())
	for {
		if l.lis == nil {
			slog.Debug("UDPListener is empty")
			break
		}
		buf := make([]byte, 1024)
		n, addr, err := l.lis.ReadFromUDP(buf)
		if err != nil {
			slog.Error("Read packet length failed", "err", err)
			continue
		}
		slog.Info("New client udp packet came from", "addr", addr)
		if n < 2 {
			slog.Warn("UDP Packet length incorrect")
			continue
		}
		data := buf[0:n]
		util.SimpleXOR(data, l.xorKey)
		var msg api.Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			slog.Error("Failed to decode message", "err", err, "addr", addr)
			continue
		}
		l.handleFunc(&msg)
		buf, _ = json.Marshal(msg)
		util.SimpleXOR(buf, l.xorKey)
		_, err = l.lis.WriteToUDP(buf, addr)
		if err != nil {
			slog.Error("Fail to reply message", "addr", addr)
			continue
		}
		slog.Debug("Handle api message ok", "addr", addr)
	}
}

// Stop implements Listener.
func (l *UDPListener) Stop() {
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
