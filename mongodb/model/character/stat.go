package character

type CharacterStat struct {
	Level        uint8        `bson:"level"`
	Exp          uint64       `bson:"exp"`
	Pop          uint32       `bson:"pop"`
	Guild        string       `bson:"guild"`
	Alliance     string       `bson:"aliiance"`
	HP           uint32       `bson:"hp"`
	MaxHP        uint32       `bson:"max_hp"`
	MP           uint32       `bson:"mp"`
	MaxMP        uint32       `bson:"max_mp"`
	AP           uint16       `bson:"ap"`
	Str          uint16       `bson:"str"`
	Dex          uint16       `bson:"dex"`
	Int          uint16       `bson:"int"`
	Luk          uint16       `bson:"luk"`
	HyperStat    HyperStat    `bson:"hyper_stat"`
	DetailedStat DetailedStat `bson:"detailed_stat"`
}

type HyperStat struct {
}

type DetailedStat struct {
}
