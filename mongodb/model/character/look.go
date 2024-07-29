package character

type CharacterLook struct {
	SkinColor        uint8  `bson:"skin_color"`
	Skin             uint32 `bson:"skin"`
	Face             uint32 `bson:"face"`
	Hair             uint32 `bson:"hair"`
	DefFaceAcc       uint32 `bson:"def_face_acc"`
	Ear              uint32 `bson:"ear"`
	Tail             uint32 `bson:"tail"`
	MixBaseHairColor uint8  `bson:"mix_base_hair_color"` // CharacterStat
	MixAddHairColor  uint8  `bson:"mix_add_hair_color"`  // CharacterStat
	MixHairBaseProb  uint8  `bson:"mix_hair_base_prob"`  // CharacterStat
	MixedHairColor   uint8  `bson:"mixed_hair_color"`    // AvatarLook
	MixHairPercent   uint8  `bson:"mix_hair_percent"`    // AvatarLook
}
