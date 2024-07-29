package outpacket

import (
	"goms/maple"
	"goms/opcode"
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

func NewConnect(recvIV, sendIV []byte) []byte {
	p := &connect{
		outPacket:    newOutPacket(0),
		version:      maple.Version,
		minorVersion: maple.MinorVersion,
		recvIV:       recvIV,
		sendIV:       sendIV,
		locale:       uint8(maple.Region),
		isTestServer: false,
	}

	// The first packet not contain opcode
	p.EncodeUint16(p.version)
	p.EncodeStr(p.minorVersion)
	p.EncodeBuffer(p.recvIV)
	p.EncodeBuffer(p.sendIV)
	if p.locale == uint8(maple.CN) {
		p.EncodeUint16(uint16(p.locale))
		p.buf[0] = 16 // packet length
	} else {
		p.EncodeByte(p.locale)
		p.buf[0] = 15 // packet length
	}
	p.EncodeBool(p.isTestServer)
	return p.buf
}

// CClientSocket::OnReceiveHotfix
func NewReceiveHotfix() []byte {
	p := newOutPacket(opcode.CClientSocket_OnReceiveHotfix)
	p.EncodeByte(0)
	return p.buf
}

// CClientSocket::OnAliveReq
func NewAliveReq() []byte {
	p := newOutPacket(opcode.CClientSocket_OnAliveReq)
	return p.buf
}

// CClientSocket::OnCheckAliveAck
func NewCheckAliveAck() []byte {
	p := newOutPacket(opcode.CClientSocket_OnCheckAliveAck)
	return p.buf
}

// CClientSocket::OnPingCheckResult
func NewPingCheckResult() []byte {
	p := newOutPacket(opcode.CClientSocket_OnPingCheckResult)
	p.Fill(5)
	return p.buf
}

// CClientSocket::OnAuthenMessage
func NewAuthenMessage() []byte {
	p := newOutPacket(opcode.CClientSocket_OnAuthenMessage)
	p.EncodeInt32(0)
	return p.buf
}
