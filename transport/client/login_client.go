package client

import (
	"goms/maple/world"
	"goms/opcode"
	"goms/packet/outpacket"
	"goms/transport"
	"net"
	"time"
)

const (
	pingInterval uint16 = 0
	pingTimeout  uint16 = 60
)

type loginClient struct {
	transport.GameClient
	server       transport.LoginServer
	pingInterval time.Duration
	pingTimeout  time.Duration
	pingAckTime  time.Time
	accountID    uint32
	worldID      world.WorldID
	channelIndex uint8
}

func NewLoginClient(conn net.Conn, server transport.LoginServer) transport.LoginClient {
	c := &loginClient{
		server:       server,
		pingInterval: time.Duration(pingInterval),
		pingTimeout:  time.Duration(pingTimeout),
	}
	c.GameClient = NewGameClient(conn, c)
	return c
}

func (c *loginClient) ClientPingServer() {
	if c.pingInterval == 0 {
		return
	}
	defer c.GameClient.Disconnect()
	ticker := time.NewTicker(c.pingInterval)
	defer ticker.Stop()
	for range ticker.C {
		if c == nil {
			break
		}
		c.SendPacket(outpacket.NewAliveReq())
		since := time.Since(c.pingAckTime)
		if since > c.pingTimeout {
			break
		}
	}
}

func (c *loginClient) ServerPingClient() {
	if c.pingInterval == 0 {
		return
	}
	c.pingAckTime = time.Now()
}

// handlePacket implements IChild.
func (c *loginClient) handlePacket(op uint16, data []byte) {
	switch op {
	case uint16(opcode.SendSecurityPacket):
		c.server.RecvSecurityPacket(c)
	case uint16(opcode.SendPermissionRequest):
		c.server.RecvPermissonRequest(c, data)
	case uint16(opcode.ApplyHotfix):
		c.server.RecvApplyHotfix(c)
	case uint16(opcode.LoginScreenLoaded):
		go c.ServerPingClient()
	case uint16(opcode.OnAliveReq_Callback):
		c.ClientPingServer()
	case uint16(opcode.CheckLoginAuthInfo):
		c.server.RecvCheckLoginAuthInfo(c, data)
	case uint16(opcode.SendGenderSetRequest):
		c.server.RecvGenderSetRequest(c, data)
	case uint16(opcode.SendSelectWorldButton):
		c.server.RecvSelectWorldButton(c, data)
	case uint16(opcode.SendSelectWorldRequest):
		c.server.RecvSelectWorldRequest(c, data)
	case uint16(opcode.GotoWorldSelect):
		c.server.RecvGotoWorldSelect(c)
	case uint16(opcode.SendWorldInformation):
		c.server.SendWorldInformation(c)
	case uint16(opcode.SendCheckSPWExistPacket):
		c.server.RecvCheckSPWExistPacket(c)
	case uint16(opcode.SendCheckDuplicateIDPacket):
		c.server.RecvCheckDuplicateIDPacket(c, data)
	case uint16(opcode.SendNewCharPacket):
		c.server.RecvNewCharPacket(c, data)
	case uint16(opcode.SendNewCharPacket_ClientToGame):
		c.server.RecvSelectCharacterRequest(c, data)
	case uint16(opcode.SendDeleteCharPacket):
		c.server.RecvDeleteCharPacket(c, data)
	case uint16(opcode.SendChangeCharOrderRequest):
		c.server.RecvChangeCharOrderRequest(data)
	case uint16(opcode.SendSelectCharacterRequest):
		c.server.RecvSelectCharacterRequest(c, data)
	}
}

// SetAccountID implements transport.LoginClient.
func (c *loginClient) SetAccountID(accountID uint32) {
	c.accountID = accountID
}

// GetAccountID implements transport.LoginClient.
func (c *loginClient) GetAccountID() uint32 {
	return c.accountID
}

// SetWorldID implements transport.LoginClient.
func (c *loginClient) SetWorldID(id uint8) {
	c.worldID = world.WorldID(id)
}

// GetWorldID implements transport.LoginClient.
func (c *loginClient) GetWorldID() world.WorldID {
	return c.worldID
}

// GetChannelIndex implements transport.LoginClient.
func (c *loginClient) GetChannelIndex() uint8 {
	return c.channelIndex
}

// SetChannelIndex implements transport.LoginClient.
func (c *loginClient) SetChannelIndex(index uint8) {
	c.channelIndex = index
}
