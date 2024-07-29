package inpacket

type BackupPacket struct {
	inPacket
	CallType         uint16
	ErrorCode        uint32
	BackupBufferSize uint16
	BackupBuffer     []byte
}

// CWvsApp::SendBackupPacket
func NewBackupPacket(data []byte) *BackupPacket {
	p := &BackupPacket{inPacket: newInPacket(data)}
	p.CallType = p.DecodeUint16()
	p.ErrorCode = p.DecodeUint32()
	p.BackupBufferSize = p.DecodeUint16()
	p.BackupBuffer = p.DecodeBuffer(int(p.BackupBufferSize))
	return p
}
