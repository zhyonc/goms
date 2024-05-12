package outpacket

import (
	"goms/maple"
	"goms/opcode"
	"goms/packet/crypt"
)

type connect struct {
	outPacket
	version      uint16
	minorVersion string
	sendIV       []byte
	recvIV       []byte
	locale       uint8
	isTestServer bool
}

func NewConnect() []byte {
	p := &connect{
		outPacket:    newOutPacket(0),
		version:      maple.Version,
		minorVersion: maple.MinorVersion,
		recvIV:       crypt.RecvIV[:],
		sendIV:       crypt.SendIV[:],
		locale:       maple.Region,
		isTestServer: false,
	}
	// The first packet not contain opcode
	// And the first packet length is 15
	p.buf[0] = 15
	p.EncodeUint16(p.version)
	p.EncodeString(p.minorVersion)
	p.EncodeBuf(p.recvIV)
	p.EncodeBuf(p.sendIV)
	p.EncodeByte(p.locale)
	p.EncodeBool(p.isTestServer)
	return p.buf
}

// Tell the client the IP of the new channel.
func NewMigrateCommand(ok bool, ip4 []byte, port uint16) []byte {
	p := newOutPacket(uint16(opcode.OnMigrateCommand))
	p.EncodeBool(ok)
	p.EncodeBuf(ip4)
	p.EncodeUint16(port)
	return p.buf
}

func NewAliveReq() []byte {
	p := newOutPacket(uint16(opcode.OnAliveReq))
	return p.buf
}

func NewCheckAliveAck() []byte {
	p := newOutPacket(uint16(opcode.OnChannelAliveReq))
	return p.buf
}

func NewAuthCodeChanged() []byte {
	p := newOutPacket(uint16(opcode.OnAuthenCodeChanged))
	p.EncodeByte(0)
	p.EncodeInt32(0)
	return p.buf
}

func NewSecurityPacket() []byte {
	p := newOutPacket(uint16(opcode.OnSecurityPacket))
	p.EncodeByte(1)
	p.EncodeBool(false)
	return p.buf
}

func NewPrivateServerAuth() []byte {
	p := newOutPacket(uint16(opcode.OnPrivateServerAuth))
	p.EncodeInt32(0)
	return p.buf
}

func NewReceiveHotfix() []byte {
	p := newOutPacket(uint16(opcode.OnReceiveHotfix))
	p.EncodeByte(0)
	return p.buf
}

func NewOpcodeEncryption() []byte {
	p := newOutPacket(uint16(opcode.OpcodeEncryption))
	p.EncodeUint32(uint32(crypt.RandNumStrLen))
	p.EncodeUint32(uint32(len(crypt.EncryptContent)))
	p.EncodeBuf(crypt.EncryptContent)
	return p.buf
}
