package inpacket

import "goms/util"

type MigrateIn struct {
	inPacket
	WorldID     uint32
	CharacterID uint32
	MAC         string
}

func NewMigrateIn(data []byte) *MigrateIn {
	p := &MigrateIn{inPacket: newInPacket(data)}
	p.WorldID = p.DecodeUint32()
	p.CharacterID = p.DecodeUint32()
	p.MAC = util.Bytes2MAC(p.DecodeBytes(6))
	return p
}
