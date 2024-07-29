package outpacket

import (
	"goms/packet"
	"goms/util"
	"io"
	"strings"
	"time"

	"golang.org/x/text/transform"
)

type outPacket struct {
	buf []byte
}

func newOutPacket(opcode uint16) outPacket {
	var p outPacket
	p.buf = make([]byte, 0)
	p.EncodeUint16(opcode)
	return p
}

func newEmptyOutPacket() outPacket {
	var p outPacket
	p.buf = make([]byte, 0)
	return p
}

func (p *outPacket) Fill(length int) {
	buf := make([]byte, length)
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeStr(s string) {
	buf := []byte(s) // ASCII Code
	p.EncodeInt16(int16(len(buf)))
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeLocalStr(s string) {
	reader := strings.NewReader(s)
	transformer := transform.NewReader(reader, packet.LgEncoder)
	buf, err := io.ReadAll(transformer)
	if err != nil {
		p.EncodeStr("")
		return
	}
	p.EncodeInt16(int16(len(buf)))
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeLocalName(s string, limitLength int) {
	reader := strings.NewReader(s)
	transformer := transform.NewReader(reader, packet.LgEncoder)
	buf, err := io.ReadAll(transformer)
	if err != nil {
		p.Fill(limitLength)
		return
	}
	bufLength := len(buf)
	if bufLength > limitLength {
		buf = buf[0:limitLength]
	} else if bufLength < limitLength {
		zeroBuf := make([]byte, limitLength-bufLength)
		buf = append(buf, zeroBuf...) // Fill zero
	}
	p.EncodeBuffer(buf)
}

func (p *outPacket) EncodeBool(v bool) {
	if v {
		p.buf = append(p.buf, 1)
	} else {
		p.buf = append(p.buf, 0)
	}
}

func (p *outPacket) EncodeBuffer(buf []byte) {
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeByte(b byte) {
	p.buf = append(p.buf, b)
}

func (p *outPacket) EncodeUint16(v uint16) {
	p.buf = append(p.buf, byte(v), byte(v>>8))
}

func (p *outPacket) EncodeUint32(v uint32) {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		buf[i] = byte(v >> (i * 8))
	}
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeUint64(v uint64) {
	buf := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buf[i] = byte(v >> (i * 8))
	}
	p.buf = append(p.buf, buf...)
}

func (p *outPacket) EncodeInt8(v int8) {
	p.EncodeByte(byte(v))
}

func (p *outPacket) EncodeInt16(v int16) {
	p.EncodeUint16(uint16(v))
}

func (p *outPacket) EncodeInt32(v int32) {
	p.EncodeUint32(uint32(v))
}

func (p *outPacket) EncodeInt64(v int64) {
	p.EncodeUint64(uint64(v))
}

func (p *outPacket) EncodeFT(time time.Time) {
	p.EncodeInt64(util.Unix2FT(time))
}
