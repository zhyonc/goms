package social

// Call by CUserRemote::Init
type Guild struct {
	Name        string `bson:"name"`
	MarkBg      uint16 `bson:"mark_bg"`
	MarkBgColor uint8  `bson:"mark_bg_color"`
	Mark        uint16 `bson:"mark"`
	MarkColor   uint8  `bson:"mark_color"`
}
