package inpacket

type SecurityPacket struct {
	inPacket
	Region  uint8
	Version uint16
}

func NewSecurityPacket(data []byte) *SecurityPacket {
	p := &SecurityPacket{inPacket: newInPacket(data)}
	p.Region = p.DecodeByte()
	p.Version = p.DecodeUint16()
	return p
}
