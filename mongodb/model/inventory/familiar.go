package inventory

type FamiliarInventory struct {
	Cards        []Familiar `bson:"cards"`
	SummonedCard Familiar   `bson:"summoned_card"`
}

type Familiar struct {
	ID      uint32 `bson:"id"`
	Index   uint32 `bson:"index"`
	Name    string `bson:"name"`
	Level   uint16 `bson:"level"`
	Exp     uint32 `bson:"exp"`
	Skill   uint16 `bson:"skill"`
	Option1 uint16 `bson:"option1"`
	Option2 uint16 `bson:"option2"`
	Option3 uint16 `bson:"option3"`
	Grade   uint8  `bson:"grade"`
}
