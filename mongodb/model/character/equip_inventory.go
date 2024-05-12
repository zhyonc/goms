package character

type EquipInventory struct {
	Equip     EquipInv     `bson:"equip"`
	CashEquip CashEquipInv `bson:"cash_equip"`
	Pet       [3]PetInv    `bson:"pet"`
	Totems    [3]uint32    `bson:"totems"`
}

type EquipInv struct {
	Hat          uint32 `bson:"hat"`
	FaceAcc      uint32 `bson:"face_acc"`
	EyeAcc       uint32 `bson:"eye_acc"`
	EarAcc       uint32 `bson:"ear_acc"`
	Top          uint32 `bson:"top"`
	Bottom       uint32 `bson:"bottom"`
	Overall      uint32 `bson:"overall"`
	Shoes        uint32 `bson:"shoes"`
	Gloves       uint32 `bson:"gloves"`
	Cape         uint32 `bson:"cape"`
	SubWeapon    uint32 `bson:"sub_weapon"`
	Weapon       uint32 `bson:"weapon"`
	Ring1        uint32 `bson:"ring1"`
	Ring2        uint32 `bson:"ring2"`
	Ring3        uint32 `bson:"ring3"`
	Ring4        uint32 `bson:"ring4"`
	Pendant      uint32 `bson:"pendant"`
	TamingMob    uint32 `bson:"taming_mob"`
	Saddle       uint32 `bson:"saddle"`
	MobEquip     uint32 `bson:"mob_equip"`
	Medal        uint32 `bson:"medal"`
	Belt         uint32 `bson:"belt"`
	Shoulder     uint32 `bson:"shoulder"`
	Pocket       uint32 `bson:"pocket"`
	Badge        uint32 `bson:"badge"`
	Emblem       uint32 `bson:"emblem"`
	Android      uint32 `bson:"android"`
	AndroidHeart uint32 `bson:"andorid_heart"`
	MonsterBook  uint32 `bson:"monster_book"`
	PendantExt   uint32 `bson:"pendant_ext"`
}

type CashEquipInv struct {
	Hat       uint32 `bson:"hat"`
	FaceAcc   uint32 `bson:"face_acc"`
	EyeAcc    uint32 `bson:"eye_acc"`
	EarAcc    uint32 `bson:"ear_acc"`
	Top       uint32 `bson:"top"`
	Bottom    uint32 `bson:"bottom"`
	Overall   uint32 `bson:"overall"`
	Shoes     uint32 `bson:"shoes"`
	Gloves    uint32 `bson:"gloves"`
	Cape      uint32 `bson:"cape"`
	SubWeapon uint32 `bson:"sub_weapon"`
	Weapon    uint32 `bson:"weapon"`
	Ring1     uint32 `bson:"ring1"`
	Ring2     uint32 `bson:"ring2"`
	Ring3     uint32 `bson:"ring3"`
	Ring4     uint32 `bson:"ring4"`
	Pendant   uint32 `bson:"pendant"`
	Belt      uint32 `bson:"belt"`
	Shoulder  uint32 `bson:"shoulder"`
}

type PetInv struct {
	ID       uint32 `bson:"id"`
	Name     string `bson:"name"`
	Acc      uint32 `bson:"acc"`
	Skill    uint32 `bson:"skill"`
	SkillExt uint32 `bson:"skill_ext"`
}
