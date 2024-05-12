package outpacket

import (
	"goms/util"
	"io"
	"strings"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type outPacket struct {
	buf     []byte
	encoder *encoding.Encoder
}

func newOutPacket(opcode uint16) outPacket {
	var p outPacket
	p.buf = make([]byte, 0)
	p.EncodeUint16(opcode)
	p.encoder = traditionalchinese.Big5.NewEncoder() //MS949
	return p
}

func (p *outPacket) EncodeString(s string) {
	bytes := []byte(s) // ASCII Code
	p.EncodeInt16(int16(len(bytes)))
	p.buf = append(p.buf, bytes...)
}

func (p *outPacket) EncodeLocalString(s string) {
	reader := strings.NewReader(s)
	transformer := transform.NewReader(reader, p.encoder)
	bytes, err := io.ReadAll(transformer)
	if err != nil {
		p.EncodeString("")
		return
	}
	p.EncodeInt16(int16(len(bytes)))
	p.buf = append(p.buf, bytes...)
}

func (p *outPacket) EncodeLocalStringBuf(s string, length int) {
	reader := strings.NewReader(s)
	transformer := transform.NewReader(reader, p.encoder)
	bytes, err := io.ReadAll(transformer)
	if err != nil {
		p.EncodeBuf(make([]byte, length))
		return
	}
	zeroLength := length - len(bytes)
	if zeroLength <= 0 {
		bytes = bytes[0:length]
	} else {
		buf := make([]byte, zeroLength)
		bytes = append(bytes, buf...)
	}
	p.buf = append(p.buf, bytes...)
}

func (p *outPacket) EncodeBool(v bool) {
	if v {
		p.buf = append(p.buf, 1)
	} else {
		p.buf = append(p.buf, 0)
	}
}

func (p *outPacket) EncodeBuf(bytes []byte) {
	p.buf = append(p.buf, bytes...)
}

func (p *outPacket) EncodeByte(b byte) {
	p.buf = append(p.buf, b)
}

func (p *outPacket) EncodeUint16(v uint16) {
	p.buf = append(p.buf, byte(v), byte(v>>8))
}

func (p *outPacket) EncodeUint32(v uint32) {
	bytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		bytes[i] = byte(v >> (i * 8))
	}
	p.buf = append(p.buf, bytes...)
}

func (p *outPacket) EncodeUint64(v uint64) {
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(v >> (i * 8))
	}
	p.buf = append(p.buf, bytes...)
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
	p.EncodeInt64(util.UnixToFT(time))
}
