package object

import "goms/maple"

type ScriptMessage struct {
	SpeakerTypeID     byte
	SpeakerTemplateID int32
	MsgType           maple.ScriptMsgType
	Param             byte
	Color             byte
	Text              string
	Prev              bool
	Next              bool
	Wait              int32
}
