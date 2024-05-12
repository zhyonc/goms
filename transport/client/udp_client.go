package client

import (
	"log/slog"
	"net"
	"time"
)

// Send packet emulator
type UDPClient struct {
	serverAddr string
	timeout    time.Duration
}

func NewUDPClient(serverAddr string, timeout uint8) *UDPClient {
	c := &UDPClient{
		serverAddr: serverAddr,
		timeout:    time.Duration(timeout) * time.Second,
	}
	return c
}

func (c *UDPClient) SendPacket(data []byte) {
	conn, err := net.Dial("udp", c.serverAddr)
	if err != nil {
		slog.Error("Dial dst server failed", "err", err, "addr", c.serverAddr)
		return
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		slog.Error("Send packet to dst server failed", "err", err, "addr", c.serverAddr)
		return
	}
	slog.Info("Send udp packet ok")
}
