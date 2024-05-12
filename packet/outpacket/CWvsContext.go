package outpacket

import "goms/opcode"

func NewHourChanged(n1, n2 uint16) []byte {
	p := newOutPacket(uint16(opcode.HourChanged))
	p.EncodeUint16(n1)
	p.EncodeUint16(n2)
	return p.buf
}
