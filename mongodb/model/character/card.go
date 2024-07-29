package character

type CharacterCard struct {
	CharacterID    uint32 `bson:"character_id"`
	CharacterLevel uint8  `bson:"character_level"`
	CharacterClass uint32 `bson:"character_class"`
}
