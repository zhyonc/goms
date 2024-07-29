package outpacket

import "goms/opcode"

// CFuncKeyMappedMan::OnInit
func NewCFuncKeyMappedMan() []byte {
	p := newOutPacket(opcode.CFuncKeyMappedMan_OnInit)
	condition := true
	p.EncodeBool(condition)
	if !condition {
		for i := 0; i < 89; i++ {
			FunckeyMappedEncode(&p)
		}
	}
	return p.buf
}

// Call by CFuncKeyMappedMan::OnInit
// FUNCKEY_MAPPED::Decode
func FunckeyMappedEncode(p *outPacket) {
	p.EncodeByte(0)
	p.EncodeUint32(0)
}

// CFuncKeyMappedMan::OnPacket
func NewPetConsumeItemID() []byte {
	p := newOutPacket(opcode.CFuncKeyMappedMan_OnPetConsumeItemID)
	p.EncodeUint32(0) // m_nPetConsumeItemID
	return p.buf
}

// CFuncKeyMappedMan::OnPacket
func NewPetConsumeMPItemID() []byte {
	p := newOutPacket(opcode.CFuncKeyMappedMan_OnPetConsumeMPItemID)
	p.EncodeUint32(0) // m_nPetConsumeMPItemID
	return p.buf
}

// CFuncKeyMappedMan::OnPacket
func NewPetConsumeSkillID() []byte {
	p := newOutPacket(opcode.CFuncKeyMappedMan_OnPetConsumeSkillID)
	p.EncodeUint32(0)
	return p.buf
}

// CFuncKeyMappedMan::OnPacket
func NewPetConsumeUNKID() []byte {
	p := newOutPacket(opcode.CFuncKeyMappedMan_OnPetConsumeUNKID)
	p.EncodeUint32(0)
	return p.buf
}
