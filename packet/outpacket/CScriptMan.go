package outpacket

import (
	"goms/maple"
	"goms/opcode"
	"goms/packet/object"
)

// CScriptMan::OnScriptMessage
func NewScriptMessage(sm *object.ScriptMessage) []byte {
	p := newOutPacket(opcode.CScriptMan_OnScriptMessage)
	p.EncodeByte(sm.SpeakerTypeID)      // nSpeakerTypeID
	p.EncodeInt32(sm.SpeakerTemplateID) // nSpeakerTemplateID
	condition := false
	p.EncodeBool(condition)
	if condition {
		p.EncodeInt32(0)
	}
	p.EncodeByte(byte(sm.MsgType)) // nMsgType
	p.EncodeByte(sm.Param)         // bParam
	p.EncodeByte(sm.Color)         // eColor
	switch sm.MsgType {
	case maple.Say:
		CScriptManOnSay(&p, sm)
	case maple.SayUNK:
		CScriptManOnSayUNK(&p)
	case maple.SayImage:
		CScriptManOnSayImage(&p)
	case maple.AskYesNo:
		CScriptManOnAskYesNo(&p, 0)
	case maple.AskText:
		CScriptManOnAskText(&p, 0)
	case maple.AskNumber:
		CScriptManOnAskNumber(&p)
	case maple.AskMenu:
		CScriptManOnAskMenu(&p, 0)
	}
	return p.buf
}
