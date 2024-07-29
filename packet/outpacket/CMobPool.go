package outpacket

import "goms/opcode"

// CMobPool::OnMobCrcKeyChanged
func NewMobCrcKeyChanged() []byte {
	p := newOutPacket(opcode.CMobPool_OnMobCrcKeyChanged)
	p.EncodeUint32(0) // m_dwMobCrcKey
	return p.buf
}
