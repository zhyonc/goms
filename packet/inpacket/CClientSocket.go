package inpacket

type PermissionRequest struct {
	inPacket
	Region  uint8
	Version uint16
}

// CClientSocket::SendPermissionRequest
func NewPermissionRequest(data []byte) *PermissionRequest {
	p := &PermissionRequest{inPacket: newInPacket(data)}
	p.Region = p.DecodeByte()
	p.Version = p.DecodeUint16()
	return p
}

type OnAliveReqCallback struct {
	inPacket
	IV []byte
}

// CClientSocket::OnAliveReq
func NewAliveReqCallback(data []byte) *OnAliveReqCallback {
	p := &OnAliveReqCallback{inPacket: newInPacket(data)}
	p.IV = p.DecodeBuffer(4)
	return p
}

// MachineId(16)=MAC(6)+UNK(10)
type MigrateIn struct {
	inPacket
	WorldID     uint32
	CharacterID uint32
	MachineID   []byte
}

func NewMigrateIn(data []byte) *MigrateIn {
	p := &MigrateIn{inPacket: newInPacket(data)}
	p.WorldID = p.DecodeUint32()
	p.CharacterID = p.DecodeUint32()
	p.MachineID = p.DecodeBuffer(16)
	return p
}
