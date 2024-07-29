package inpacket

import (
	"bytes"
	"goms/packet"
	"io"

	"golang.org/x/text/transform"
)

type inPacket struct {
	pos  int
	data []byte
	size int
}

func newInPacket(data []byte) inPacket {
	var p inPacket
	p.pos = 0
	p.data = data
	p.size = len(p.data)
	return p
}

func (p *inPacket) Skip(offset int) {
	if offset <= 0 {
		return
	}
	if p.pos+int(offset) > p.size {
		p.pos = p.size - 1
		return
	}
	p.pos += offset
}

func (p *inPacket) DecodeStr() string {
	length := int(p.DecodeUint16())
	if p.pos+length > p.size {
		return ""
	}
	str := string(p.data[p.pos : p.pos+length])
	p.pos += length
	return str
}

func (p *inPacket) DecodeLocalStr() string {
	length := int(p.DecodeUint16())
	buf := p.DecodeBuffer(length)
	reader := bytes.NewReader(buf)
	transformer := transform.NewReader(reader, packet.LgDecoder)
	buf, err := io.ReadAll(transformer)
	if err != nil {
		return ""
	}
	return string(buf)
}

func (p *inPacket) DecodeBuffer(count int) []byte {
	if p.pos+count > p.size {
		return nil
	}
	buf := make([]byte, count)
	for i := 0; i < count; i++ {
		buf[i] = p.DecodeByte()
	}
	return buf
}

func (p *inPacket) DecodeByte() byte {
	if p.pos+1 > p.size {
		return 0
	}
	b := p.data[p.pos]
	p.pos += 1
	return b
}

func (p *inPacket) DecodeBool() bool {
	if p.pos+1 > p.size {
		return false
	}
	b := p.data[p.pos]
	p.pos += 1
	return b == 1
}

func (p *inPacket) DecodeUint16() uint16 {
	if p.pos+2 > p.size {
		return 0
	}
	number := uint16(p.data[p.pos]) | uint16(p.data[p.pos+1])<<8
	p.pos += 2
	return number
}

func (p *inPacket) DecodeUint32() uint32 {
	if p.pos+4 > p.size {
		return 0
	}
	var number uint32 = 0
	for i := 0; i < 4; i++ {
		index := p.pos + i
		number |= uint32(p.data[index]) << (i * 8)
	}
	p.pos += 4
	return number
}

func (p *inPacket) DecodeUint64() uint64 {
	if p.pos+4 > p.size {
		return 0
	}
	var number uint64 = 0
	for i := 0; i < 8; i++ {
		index := p.pos + i
		number |= uint64(p.data[index]) << (i * 8)
	}
	p.pos += 8
	return number
}

func (p *inPacket) DecodeInt8() int8 {
	return int8(p.DecodeByte())
}

func (p *inPacket) DecodeInt16() int16 {
	return int16(p.DecodeUint16())
}

func (p *inPacket) DecodeInt32() int32 {
	return int32(p.DecodeUint32())
}

func (p *inPacket) DecodeInt64() int64 {
	return int64(p.DecodeUint64())
}
