package client

import (
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/mongodb/model/social"
	"goms/network"
	"goms/network/manager"
	"goms/opcode"
	"goms/packet/object"
	"goms/util"
	"log/slog"
	"net"

	"github.com/dop251/goja"
)

type channelServerHandler func(c network.ChannelClient, data []byte)

type channelClient struct {
	network.Client
	handlerFuncMap            map[uint16]channelServerHandler
	damageSeed                [3]uint32
	char                      *character.Character
	inv                       *inventory.Inventory
	soc                       *social.Social
	screenWidth, screenHeight int
	vm                        *goja.Runtime
	sm                        *object.ScriptMessage
}

func NewChannelClient(conn net.Conn, recvIV, sendIV [4]byte, s network.ChannelServer) network.ChannelClient {
	c := &channelClient{}
	c.Client = NewBaseClient(conn, recvIV, sendIV, c, s)
	c.damageSeed = [3]uint32{util.GetRandomSeed(), util.GetRandomSeed(), util.GetRandomSeed()}
	c.initScriptManager()
	c.handlerFuncMap = make(map[uint16]channelServerHandler)
	c.handlerFuncMap[opcode.CClientSocket_MigrateIn] = s.RecvMigrateIn
	c.handlerFuncMap[opcode.CUserLocal_SendClientResolution] = s.RecvClientResolution
	c.handlerFuncMap[opcode.CWvsContext_RequestInstanceTable] = s.RecvRequestInstanceTable
	c.handlerFuncMap[opcode.CQuickslotKeyMappedMan_SaveQuickslotKeyMap] = s.RecvSaveQuickslotKeyMap
	c.handlerFuncMap[opcode.CUserLocal_TalkToNpc] = s.RecvTalkToNpc
	c.handlerFuncMap[opcode.CUserLocal_TalkToNpcStep] = s.RecvTalkToNpcStep
	return c
}

func (c *channelClient) initScriptManager() {
	c.vm = goja.New()
	c.vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	c.vm.Set(string(manager.NPCManagerName), manager.NewNPCManager(c))
	c.vm.Set(string(manager.QuestManagerName), manager.NewQuestManager(c))
}

// Disconnect implements network.ChannelClient.
// Subtle: this method shadows the method (Client).Disconnect of channelClient.Client.
func (c *channelClient) Disconnect() {
	c.Client.Disconnect()
}

// GetClientIP implements network.ChannelClient.
// Subtle: this method shadows the method (Client).GetClientIP of channelClient.Client.
func (c *channelClient) GetClientIP() string {
	return c.Client.GetClientIP()
}

// IsDisconnected implements network.ChannelClient.
// Subtle: this method shadows the method (Client).IsDisconnected of channelClient.Client.
func (c *channelClient) IsDisconnected() bool {
	return c.Client.IsDisconnected()
}

// RecvPacket implements network.ChannelClient.
// Subtle: this method shadows the method (Client).RecvPacket of channelClient.Client.
func (c *channelClient) RecvPacket() {
	c.Client.RecvPacket()
}

// SendPacket implements network.ChannelClient.
// Subtle: this method shadows the method (Client).SendPacket of channelClient.Client.
func (c *channelClient) SendPacket(buf []byte) {
	c.Client.SendPacket(buf)
}

// handlePacket implements IChild.
func (c *channelClient) handlePacket(op uint16, data []byte) {
	handlerFunc, ok := c.handlerFuncMap[op]
	if !ok {
		slog.Debug("Not found handler function", "opcode", op, "field", opcode.InMap[op])
		return
	}
	handlerFunc(c, data)
}

// BindGameData implements network.ChannelClient.
func (c *channelClient) BindGameData(char *character.Character, inv *inventory.Inventory, soc *social.Social) {
	c.char = char
	c.inv = inv
	c.soc = soc
}

// GetChar implements network.ChannelClient.
func (c *channelClient) GetChar() *character.Character {
	return c.char
}

// GetInv implements network.ChannelClient.
func (c *channelClient) GetInv() *inventory.Inventory {
	return c.inv
}

// GetSoc implements network.ChannelClient.
func (c *channelClient) GetSoc() *social.Social {
	return c.soc
}

// GetDamageSeed implements network.ChannelClient.
func (c *channelClient) GetDamageSeed() [3]uint32 {
	return c.damageSeed
}

// GetResolution implements network.ChannelClient.
func (c *channelClient) GetResolution() (int, int) {
	return c.screenWidth, c.screenHeight
}

// SetResolution implements network.ChannelClient.
func (c *channelClient) SetResolution(screenWidth int, screenHeight int) {
	c.screenWidth = screenWidth
	c.screenHeight = screenHeight
}

// SetScriptMessage implements network.ChannelClient.
func (c *channelClient) SetScriptMessage(sm *object.ScriptMessage) {
	c.sm = sm
}

// GetScriptMessage implements network.ChannelClient.
func (c *channelClient) GetScriptMessage() *object.ScriptMessage {
	return c.sm
}

// StartNPCScript implements network.ChannelClient.
func (c *channelClient) StartNPCScript(p *goja.Program) {
	c.runScript(p, "start")
}

func (c *channelClient) runScript(p *goja.Program, initFuncName string) {
	_, err := c.vm.RunProgram(p)
	if err != nil {
		slog.Error("Failed to run script", "err", err, "charID", c.char.ID)
		return
	}
	initFunc, ok := goja.AssertFunction(c.vm.Get(initFuncName))
	if !ok {
		slog.Error("start function not found in script", "charID", c.char.ID)
		return
	}
	_, err = initFunc(goja.Undefined())
	if err != nil {
		slog.Error("Failed to call function", "err", err, "funcName", initFuncName)
		return
	}
	slog.Debug("Run script ok", "charID", c.char.ID)
}

// ResumeNPCScript implements network.ChannelClient.
func (c *channelClient) ResumeNPCScript(typ, mode, selection int8) {
	actionFunc, ok := goja.AssertFunction(c.vm.Get("action"))
	if !ok {
		slog.Error("start function not found in script", "charID", c.char.ID)
		return
	}
	_, err := actionFunc(goja.Undefined(), c.vm.ToValue(mode), c.vm.ToValue(typ), c.vm.ToValue(selection))
	if err != nil {
		slog.Error("Failed to call function", "err", err, "funcName", "action")
		return
	}
	slog.Debug("Run script ok", "charID", c.char.ID)
}
