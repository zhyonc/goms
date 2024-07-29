package character

type CharacterKeymap struct {
	KeySettingType uint32     `bson:"key_setting_type"`
	QuickSlotKeys  [28]uint32 `bson:"quick_slot_keys"`
}
