package outpacket

import "goms/packet/object"

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnSay
func CScriptManOnSay(p *outPacket, sm *object.ScriptMessage) {
	if sm.Param&4 != 0 {
		p.EncodeInt32(sm.SpeakerTemplateID) // nSpeakerTemplateID
	}
	p.EncodeLocalStr(sm.Text)
	p.EncodeBool(sm.Prev)  // bPrev
	p.EncodeBool(sm.Next)  // bNext
	p.EncodeInt32(sm.Wait) // tWait
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnSayUNK
func CScriptManOnSayUNK(p *outPacket) {
	p.EncodeLocalStr("")
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeInt32(0)
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnSayImage
func CScriptManOnSayImage(p *outPacket) {
	count := 0
	p.EncodeByte(byte(count))
	for i := 0; i < count; i++ {
		p.EncodeLocalStr("")
	}
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnAskYesNo
func CScriptManOnAskYesNo(p *outPacket, param byte) {
	if param&4 != 0 {
		p.EncodeInt32(0) // nSpeakerTemplateID
	}
	p.EncodeLocalStr("")
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnAskText
func CScriptManOnAskText(p *outPacket, param byte) {
	if param&4 != 0 {
		p.EncodeInt32(0) // nSpeakerTemplateID
	}
	p.EncodeLocalStr("")
	p.EncodeLocalStr("")
	p.EncodeInt16(0) // nLenMin
	p.EncodeInt16(0) // nLenMax
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnAskNumber
func CScriptManOnAskNumber(p *outPacket) {
	p.EncodeLocalStr("")
	p.EncodeInt32(0) // nDef
	p.EncodeInt32(0) // nMin
	p.EncodeInt32(0) // nMax
}

// Call by CScriptMan::OnScriptMessage
// CScriptMan::OnAskMenu
func CScriptManOnAskMenu(p *outPacket, param byte) {
	if param&4 != 0 {
		p.EncodeInt32(0) // nSpeakerTemplateID
	}
	p.EncodeLocalStr("")
}
