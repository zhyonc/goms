package outpacket

import (
	"goms/opcode"
)

// Call by CStage::OnSetField
// CFieldCustom::Decode
func CFieldCustomEncode(p *outPacket) {
	p.EncodeUint32(0) // nPartyBonusExpRate
	p.EncodeStr("")   // sBgm
	p.EncodeUint32(0) // dwBgFieldID
}

// CFiled::SetQuestClear
func NewSetQuestClear() []byte {
	p := newOutPacket(opcode.CField_OnSetQuestClear)
	return p.buf
}

// CField::OnMouseMove call CUIWorldMapSearch::OnMouseMove
func NewMouseMove() []byte {
	p := newOutPacket(opcode.CField_OnMouseMove)
	return p.buf
}

// CQuickslotKeyMappedMan::OnInit
func NewCQuickslotKeyMappedMan(quickSlotKeys [28]uint32) []byte {
	p := newOutPacket(opcode.CQuickslotKeyMappedMan_OnInit)
	keyLength := len(quickSlotKeys)
	condition := keyLength > 0
	p.EncodeBool(condition)
	if condition {
		// CInPacket::DecodeBuffer(0x70u);
		for i := 0; i < keyLength; i++ {
			p.EncodeUint32(quickSlotKeys[i])
		}
	}
	return p.buf
}

// CField::OnStarPlanetBurningTimeInfo
func NewStarPlanetBurningTimeInfo() []byte {
	p := newOutPacket(opcode.CField_OnStarPlanetBurningTimeInfo)
	p.EncodeUint32(0)
	p.EncodeStr("")
	return p.buf
}

// CField::OnAdminResult
func NewAdminResult() []byte {
	p := newOutPacket(opcode.CField_OnAdminResult)
	p.EncodeUint32(51) // LABEL_13
	p.EncodeByte(1)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	return p.buf
}

// CField::OnSetQuickMoveInfo
func NewSetQuickMoveInfo() []byte {
	p := newOutPacket(opcode.CField_OnSetQuickMoveInfo)
	count := 0
	p.EncodeByte(byte(count))
	for i := 0; i < count; i++ {
		QuickMoveInfoEncode(&p)
	}
	return p.buf
}

// CField::OnMomentAreaOnOffAll
func NewMomentAreaOnOffAll() []byte {
	p := newOutPacket(opcode.CField_OnMomentAreaOnOffAll)
	// CAttrField::DecodeMomentAreaActivateList Start
	condition := false
	p.EncodeBool(condition)
	if condition {
		GeometryCMomentAreaEncodeActivateList(&p)
	}
	condition = false
	p.EncodeBool(condition)
	if condition {
		GeometryCMomentAreaEncodeActivateList(&p)
	}
	condition = false
	p.EncodeBool(condition)
	if condition {
		GeometryCMomentAreaEncodeActivateList(&p)
	}
	return p.buf
}

// Geometry::CMomentArea::DecodeActivateList
func GeometryCMomentAreaEncodeActivateList(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeStr("")
	}
}
