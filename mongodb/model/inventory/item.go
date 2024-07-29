package inventory

type ItemInventory struct {
	Money   uint64      `bson:"money"`
	Equip   EquipSlot   `bson:"equip"`
	Consume ConsumeSlot `bson:"consume"`
	SetUp   SetUpSlot   `bson:"setup"`
	Etc     EtcSlot     `bson:"etc"`
	Cash    CashSlot    `bson:"cash"`
}

type BaseSlot struct {
	SlotSize uint8 `bson:"slot_size"`
	Items    []ItemSlotBundle
}

type EquipSlot struct {
	BaseSlot
	Items []ItemSlotEquip
}

type ConsumeSlot struct {
	BaseSlot
	ExpConsumeItems []ExpConsumeItem
}

type ExpConsumeItem struct {
	ItemID       uint32 `bson:"item_id"`
	MinLev       uint32 `bson:"min_lev"`
	MaxLev       uint32 `bson:"max_lev"`
	RemainingExp uint64 `bson:"remaining_exp"`
}

type ItemSlotPot struct {
}

type SetUpSlot struct {
	BaseSlot
}

type EtcSlot struct {
	BaseSlot
}

type CashSlot struct {
	BaseSlot
}
