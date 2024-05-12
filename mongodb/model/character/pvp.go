package character

type CharacterPVP struct {
	Exp        uint32 `bson:"exp"`
	Rank       uint8  `bson:"rank"`
	Point      uint32 `bson:"point"`
	ModeLevel  uint8  `bson:"mode_level"`
	ModeType   uint8  `bson:"mode_type"`
	EventPoint uint32 `bson:"event_point"`
}
