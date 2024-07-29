package inpacket

type SaveQuickslotKeyMap struct {
	inPacket
	Keycode [28]uint32
}

// CQuickslotKeyMappedMan::SaveQuickslotKeyMap
func NewSaveQuickslotKeyMap(data []byte) *SaveQuickslotKeyMap {
	p := &SaveQuickslotKeyMap{inPacket: newInPacket(data)}
	for i := 0; i < len(p.Keycode); i++ {
		p.Keycode[i] = p.DecodeUint32()
	}
	return p
}
