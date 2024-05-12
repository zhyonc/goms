package test

import (
	"goms/packet/outpacket"
	"goms/transport/client"
	"testing"
)

const (
	loginServerAddr   string = "127.0.0.1:8484"
	channelServerAddr string = "127.0.0.1:8500"
	timeout           uint8  = 3
)

var (
	loginDialer   *client.UDPClient = client.NewUDPClient(loginServerAddr, timeout)
	channelDialer *client.UDPClient = client.NewUDPClient(channelServerAddr, timeout)
)

func TestCheckAliveAck(t *testing.T) {
	loginDialer.SendPacket(outpacket.NewAliveReq())
}
