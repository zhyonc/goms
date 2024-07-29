package network

import (
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/mongodb/model/social"
	"goms/network/server/api"
	"goms/packet/object"
	"net"

	"github.com/dop251/goja"
)

type Server interface {
	Run()
	Stop()
	KickClient(ip string)
}

type LoginServer interface {
	Server
	HandleTCPConn(conn net.Conn)
	HandleUDPMessage(msg *api.Message)
	RecvPermissonRequest(c LoginClient, data []byte)
	RecvApplyRSAKey(c LoginClient, data []byte)
	RecvApplyHotfix(c LoginClient, data []byte)
	RecvCheckLoginAuthInfo(c LoginClient, data []byte)
	RecvGenderSetRequest(c LoginClient, data []byte)
	RecvWorldStatusCheck(c LoginClient, data []byte)
	RecvSelectWorldRequest(c LoginClient, data []byte)
	RecvGotoWorldSelect(c LoginClient, data []byte)
	RecvCreateCharStep(c LoginClient, data []byte)
	RecvCheckDuplicateIDPacket(c LoginClient, data []byte)
	RecvNewCharPacket(c LoginClient, data []byte)
	RecvChangeCharOrderRequest(c LoginClient, data []byte)
	RecvDeleteCharPacket(c LoginClient, data []byte)
	RecvReservedDeleteCharacterConfirmStep(c LoginClient, data []byte)
	RecvReservedDeleteCharacterCancelStep(c LoginClient, data []byte)
	RecvSelectCharacterRequest(c LoginClient, data []byte)
}

type WorldServer interface {
	Server
	HandleUDPMessage(msg *api.Message)
	CheckAuthorizedIP(ip string) bool
}

type ChannelServer interface {
	Server
	HandleTCPConn(conn net.Conn)
	GetClientCount() uint32
	SaveGameData(c ChannelClient)
	RecvMigrateIn(c ChannelClient, data []byte)
	RecvClientResolution(c ChannelClient, data []byte)
	RecvRequestInstanceTable(c ChannelClient, data []byte)
	RecvSaveQuickslotKeyMap(c ChannelClient, data []byte)
	RecvTalkToNpc(c ChannelClient, data []byte)
	RecvTalkToNpcStep(c ChannelClient, data []byte)
}

type Client interface {
	Disconnect()
	IsDisconnected() bool
	RecvPacket()
	SendPacket(buf []byte)
	GetClientIP() string
}

type LoginClient interface {
	Client
	SetAccountID(accountID uint32)
	GetAccountID() uint32
	SetWorldID(id uint8)
	GetWorldID() uint8
	SetChannelIndex(index uint8)
	GetChannelIndex() uint8
}

type ChannelClient interface {
	Client
	BindGameData(char *character.Character, inv *inventory.Inventory, soc *social.Social)
	GetChar() *character.Character
	GetInv() *inventory.Inventory
	GetSoc() *social.Social
	GetDamageSeed() [3]uint32
	SetResolution(screenWidth int, screenHeight int)
	GetResolution() (int, int)
	SetScriptMessage(sm *object.ScriptMessage)
	GetScriptMessage() *object.ScriptMessage
	StartNPCScript(p *goja.Program)
	ResumeNPCScript(typ, mode, selection int8)
}
