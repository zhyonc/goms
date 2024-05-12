package inpacket

import (
	"io"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type inPacket struct {
	pos     int
	data    []byte
	size    int
	decoder *encoding.Decoder
}

func newInPacket(data []byte) inPacket {
	var p inPacket
	p.pos = 0
	p.data = data
	p.size = len(p.data)
	p.decoder = traditionalchinese.Big5.NewDecoder() //MS949 code
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

func (p *inPacket) DecodeString() string {
	length := int(p.DecodeUint16())
	if p.pos+length > p.size {
		return ""
	}
	str := string(p.data[p.pos : p.pos+length])
	p.pos += length
	return str
}

func (p *inPacket) DecodeLocalString() string {
	str := p.DecodeString()
	reader := strings.NewReader(str)
	transformer := transform.NewReader(reader, p.decoder)
	bytes, err := io.ReadAll(transformer)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p *inPacket) DecodeBytes(count int) []byte {
	if p.pos+count > p.size {
		return nil
	}
	bytes := make([]byte, count)
	for i := 0; i < count; i++ {
		bytes[i] = p.DecodeByte()
	}
	return bytes
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
