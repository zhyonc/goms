package client

import (
	"goms/opcode"
	"goms/transport"
	"net"
)

type channelClient struct {
	transport.GameClient
	server transport.ChannelServer
}

func NewChannelClient(conn net.Conn, server transport.ChannelServer) transport.ChannelClient {
	c := &channelClient{
		server: server,
	}
	c.GameClient = NewGameClient(conn, c)
	return c
}

// handlePacket implements IChild.
func (c *channelClient) handlePacket(op uint16, data []byte) {
	switch op {
	case uint16(opcode.SendSecurityPacket):
		c.server.RecvSecurityPacket(c)
	case uint16(opcode.MigrateIn):
		c.server.RecvMigrateIn(c, data)
	}
}
