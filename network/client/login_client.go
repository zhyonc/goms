package client

import (
	"goms/network"
	"goms/opcode"
	"log/slog"
	"net"
)

type loginServerHandler func(c network.LoginClient, data []byte)

type loginClient struct {
	network.Client
	accountID      uint32
	worldID        uint8
	channelIndex   uint8
	handlerFuncMap map[uint16]loginServerHandler
}

func NewLoginClient(conn net.Conn, recvIV, sendIV [4]byte, s network.LoginServer) network.LoginClient {
	c := &loginClient{}
	c.Client = NewBaseClient(conn, recvIV, sendIV, c, s)
	c.handlerFuncMap = make(map[uint16]loginServerHandler)
	c.handlerFuncMap[opcode.CClientSocket_SendPermissionRequest] = s.RecvPermissonRequest
	c.handlerFuncMap[opcode.CLogin_ApplyRSAKey] = s.RecvApplyRSAKey
	c.handlerFuncMap[opcode.CClientSocket_ApplyHotfix] = s.RecvApplyHotfix
	c.handlerFuncMap[opcode.CLogin_CheckLoginAuthInfo] = s.RecvCheckLoginAuthInfo
	c.handlerFuncMap[opcode.CLogin_SendGenderSetRequest] = s.RecvGenderSetRequest
	c.handlerFuncMap[opcode.CLogin_SendWorldStatusCheck] = s.RecvWorldStatusCheck
	c.handlerFuncMap[opcode.CLogin_SendSelectWorldRequest] = s.RecvSelectWorldRequest
	c.handlerFuncMap[opcode.CLogin_GotoWorldSelect] = s.RecvGotoWorldSelect
	c.handlerFuncMap[opcode.CLogin_OnCreateCharStep_Callback] = s.RecvCreateCharStep
	c.handlerFuncMap[opcode.CLogin_SendCheckDuplicateIDPacket] = s.RecvCheckDuplicateIDPacket
	c.handlerFuncMap[opcode.CLogin_SendNewCharPacket] = s.RecvNewCharPacket
	c.handlerFuncMap[opcode.CLogin_DirectGoToField] = s.RecvSelectCharacterRequest // After OnCreateNewCharacterResult, the same as SelectCharacterRequest
	c.handlerFuncMap[opcode.CLogin_SendChangeCharOrderRequest] = s.RecvChangeCharOrderRequest
	c.handlerFuncMap[opcode.CLogin_SendDeleteCharPacket] = s.RecvDeleteCharPacket
	c.handlerFuncMap[opcode.CLogin_SendReservedDeleteCharacterCancelStep] = s.RecvReservedDeleteCharacterCancelStep
	c.handlerFuncMap[opcode.CLogin_SendReservedDeleteCharacterConfirmStep] = s.RecvReservedDeleteCharacterConfirmStep
	c.handlerFuncMap[opcode.CLogin_SendSelectCharacterRequest] = s.RecvSelectCharacterRequest
	return c
}

// handlePacket implements IChild.
func (c *loginClient) handlePacket(op uint16, data []byte) {
	handlerFunc, ok := c.handlerFuncMap[op]
	if !ok {
		slog.Debug("Not found handler function", "opcode", op, "field", opcode.InMap[op])
		return
	}
	handlerFunc(c, data)
}

// SetAccountID implements network.LoginClient.
func (c *loginClient) SetAccountID(accountID uint32) {
	c.accountID = accountID
}

// GetAccountID implements network.LoginClient.
func (c *loginClient) GetAccountID() uint32 {
	return c.accountID
}

// SetWorldID implements network.LoginClient.
func (c *loginClient) SetWorldID(id uint8) {
	c.worldID = id
}

// GetWorldID implements network.LoginClient.
func (c *loginClient) GetWorldID() uint8 {
	return c.worldID
}

// GetChannelIndex implements network.LoginClient.
func (c *loginClient) GetChannelIndex() uint8 {
	return c.channelIndex
}

// SetChannelIndex implements network.LoginClient.
func (c *loginClient) SetChannelIndex(index uint8) {
	c.channelIndex = index
}
