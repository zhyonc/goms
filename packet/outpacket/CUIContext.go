package outpacket

import "goms/opcode"

func checkFieldID(fieldID uint32) bool {
	return fieldID-340000000 <= 8 || fieldID == 340000009 || fieldID/10000000 != 34
}

// CUIContext::OnGrowthHelper
func NewGrowthHelper(fieldID uint32, v3 uint16) []byte {
	p := newOutPacket(opcode.CUIContext_OnGrowthHelper)
	if checkFieldID(fieldID) {
		p.EncodeUint16(uint16(v3))
		switch v3 {
		case 1:
			// CGrowthHelperMan::SetCheckCount
			p.EncodeUint32(0) // nCntNormal
			p.EncodeUint32(0) // nCntHard
			count := 0
			p.EncodeUint32(uint32(count))
			count = 0
			p.EncodeUint32(uint32(count))
		case 2:
		case 3:
			// CGrowthHelperMan::OnRecommendItemList
			count := 0
			p.EncodeUint32(uint32(count))
			p.EncodeUint32(0)
		case 4:
			// CUIGrowthHelper::ShowAlarmEffect
		}
	}
	return p.buf
}

// CUIContext::OnContentsMap
func NewContentsMap(fieldID uint32) []byte {
	p := newOutPacket(opcode.CUIContext_OnContentsMap)
	var condition uint16 = 1
	p.EncodeUint16(condition)
	if checkFieldID(fieldID) && condition == 1 {
		// CContentsMapMan::OnReceiveFieldContentRewardData
		count := 0
		p.EncodeUint32(uint32(count))
		p.EncodeUint32(0)
	}
	return p.buf
}
