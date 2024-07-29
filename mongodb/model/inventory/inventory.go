package inventory

import (
	"goms/maple"
	"goms/packet/inpacket"
	"time"
)

type Inventory struct {
	ID           uint32             `bson:"_id"`
	EquipInv     EquipInventory     `bson:"equip_inv"`
	FamiliarInv  FamiliarInventory  `bson:"familiar_inv"` // SampleItemID 2870001
	ItemInv      ItemInventory      `bson:"item_inv"`
	ItemPotInv   ItemPotInventory   `bson:"item_pot_inv"`   // SampleItemID 2440001
	PotionPotInv PotionPotInventory `bson:"potion_pot_inv"` // SampleItemID 5820000
	BitsInv      BitsInventory      `bson:"bits_inv"`       // SampleItemID 1680000
	UpdateDate   time.Time          `bson:"update_date"`
}

func NewInventory(characterID uint32, in *inpacket.CharPacket) *Inventory {
	inv := &Inventory{
		ID:           characterID,
		EquipInv:     EquipInventory{},
		FamiliarInv:  FamiliarInventory{},
		ItemInv:      ItemInventory{},
		ItemPotInv:   ItemPotInventory{},
		PotionPotInv: PotionPotInventory{},
		BitsInv:      BitsInventory{},
		UpdateDate:   time.Now(),
	}
	// Initial Equip
	// Hat
	inv.EquipInv.Equip.Hat = NewItemSlotEquip(in.Hat, uint16(maple.Hat))
	// Clothes
	if in.Overall > 0 {
		inv.EquipInv.Equip.Overall = NewItemSlotEquip(in.Overall, uint16(maple.Overall))
	} else if in.Top > 0 {
		inv.EquipInv.Equip.Top = NewItemSlotEquip(in.Top, uint16(maple.Top))
	}
	// Bottom
	if in.Bottom > 0 && in.Overall == 0 {
		inv.EquipInv.Equip.Bottom = NewItemSlotEquip(in.Bottom, uint16(maple.Bottom))
	}
	// Cape
	inv.EquipInv.Equip.Cape = NewItemSlotEquip(in.Cape, uint16(maple.Cape))
	// Shoes
	inv.EquipInv.Equip.Shoes = NewItemSlotEquip(in.Shoes, uint16(maple.Shoes))
	// Gloves
	inv.EquipInv.Equip.Gloves = NewItemSlotEquip(in.Gloves, uint16(maple.Gloves))
	// Weapon
	inv.EquipInv.Equip.Weapon = NewItemSlotEquip(in.Weapon, uint16(maple.Weapon))
	// SubWeapon
	inv.EquipInv.Equip.SubWeapon = NewItemSlotEquip(in.SubWeapon, uint16(maple.SubWeapon))
	// Item slot size
	inv.ItemInv.Equip.SlotSize = 72
	inv.ItemInv.Consume.SlotSize = 72
	inv.ItemInv.SetUp.SlotSize = 72
	inv.ItemInv.Etc.SlotSize = 72
	inv.ItemInv.Cash.SlotSize = 48
	return inv
}
