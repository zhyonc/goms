package record

type Record struct {
	ID     uint32  `bson:"id"`
	Quests []Quest `bson:"quests"`
}
