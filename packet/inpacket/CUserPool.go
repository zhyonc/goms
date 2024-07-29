package inpacket

type ClientResolution struct {
	inPacket
	ResolutionType int16
}

func NewClientResolution(data []byte) *ClientResolution {
	p := &ClientResolution{inPacket: newInPacket(data)}
	p.ResolutionType = p.DecodeInt16()
	return p
}

type TalkToNpc struct {
	inPacket
	ObjectID int32
	CharPosX int16
	CharPosY int16
}

// CUserLocal::TalkToNpc
func NewTalkToNpc(data []byte) *TalkToNpc {
	p := &TalkToNpc{inPacket: newInPacket(data)}
	p.ObjectID = p.DecodeInt32()
	p.CharPosX = p.DecodeInt16()
	p.CharPosY = p.DecodeInt16()
	return p
}

// Found in CScriptMan::OnScriptMessage
type TalkToNpcAction struct {
	inPacket
	Type      int8
	Mode      int8
	Selection int8
}

func NewTalkToNpcAction(data []byte) *TalkToNpcAction {
	p := &TalkToNpcAction{inPacket: newInPacket(data)}
	length := len(data)
	p.Type = p.DecodeInt8()
	if length > 0 {
		if length == 2 {
			p.Mode = p.DecodeInt8()
		}
		if length == 3 {
			p.Selection = p.DecodeInt8()
		}
	}
	return p
}
