package character

type CharacterLook struct {
	SkinColor        uint8  `bson:"skin_color"`
	Skin             uint32 `bson:"skin"`
	Face             uint32 `bson:"face"`
	Hair             uint32 `bson:"hair"`
	SpecialFace      uint32 `bson:"special_face"`
	MixBaseHairColor uint8  `bson:"mix_base_hair_color"`
	MixAddHairColor  uint8  `bson:"mix_add_hair_color"`
	MixHairBaseProb  uint8  `bson:"mix_hair_base_prob"`
	MixedHairColor   uint8  `bson:"mixed_hair_color"`
	MixHairPercent   uint8  `bson:"mix_hair_percent"`
}
