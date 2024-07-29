package inpacket

type RequestInstanceTable struct {
	inPacket
	TableName string
	Col       int32
	Row       int32
}

// CWvsContext::RequestInstanceTable
func NewRequestInstanceTable(data []byte) *RequestInstanceTable {
	p := &RequestInstanceTable{inPacket: newInPacket(data)}
	p.TableName = p.DecodeStr()
	p.Col = p.DecodeInt32()
	p.Row = p.DecodeInt32()
	return p
}

// CWvsContext::OnPartyResult callback
type PartyResultCallback struct {
	inPacket
	IsFadeWndExist bool
	PartyID        uint32
}

func NewPartyResultCallback(data []byte) *PartyResultCallback {
	p := &PartyResultCallback{inPacket: newInPacket(data)}
	p.IsFadeWndExist = p.DecodeBool()
	p.PartyID = p.DecodeUint32()
	return p
}
