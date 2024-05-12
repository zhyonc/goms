package transport

import (
	"goms/maple/world"
)

type Server interface {
	Run()
	Stop()
}

type LoginServer interface {
	Server
	RecvSecurityPacket(c LoginClient)
	RecvPermissonRequest(c LoginClient, data []byte)
	RecvApplyHotfix(c LoginClient)
	RecvCheckLoginAuthInfo(c LoginClient, data []byte)
	RecvGenderSetRequest(c LoginClient, data []byte)
	SendWorldInformation(c LoginClient)
	RecvSelectWorldButton(c LoginClient, data []byte)
	RecvSelectWorldRequest(c LoginClient, data []byte)
	RecvGotoWorldSelect(c LoginClient)
	RecvCheckSPWExistPacket(c LoginClient)
	RecvCheckDuplicateIDPacket(c LoginClient, data []byte)
	RecvNewCharPacket(c LoginClient, data []byte)
	RecvDeleteCharPacket(c LoginClient, data []byte)
	RecvChangeCharOrderRequest(data []byte)
	RecvSelectCharacterRequest(c LoginClient, data []byte)
}

type ChannelServer interface {
	Server
	RecvSecurityPacket(c ChannelClient)
	RecvMigrateIn(c ChannelClient, data []byte)
}

type Client interface {
	Disconnect()
}

type GameClient interface {
	Client
	RecvPacket()
	SendPacket(buf []byte)
}

type LoginClient interface {
	GameClient
	SetAccountID(accountID uint32)
	GetAccountID() uint32
	SetWorldID(id uint8)
	GetWorldID() world.WorldID
	SetChannelIndex(index uint8)
	GetChannelIndex() uint8
}

type ChannelClient interface {
	GameClient
}
