package inventory

import (
	"reflect"
	"time"
)

type EquipInventory struct {
	Totems    [3]ItemSlotEquip `bson:"totems"`
	Equip     Equip            `bson:"equip"`
	CashEquip CashEquip        `bson:"cash_equip"`
	Pets      [3]ItemSlotPet   `bson:"pets"`

	PendantExtExpireDate time.Time `bson:"pendant_ext_expire_date"`
}

type Equip struct {
	Hat          ItemSlotEquip `bson:"hat"`
	FaceAcc      ItemSlotEquip `bson:"face_acc"`
	EyeAcc       ItemSlotEquip `bson:"eye_acc"`
	EarAcc       ItemSlotEquip `bson:"ear_acc"`
	Top          ItemSlotEquip `bson:"top"`
	Bottom       ItemSlotEquip `bson:"bottom"`
	Overall      ItemSlotEquip `bson:"overall"`
	Shoes        ItemSlotEquip `bson:"shoes"`
	Gloves       ItemSlotEquip `bson:"gloves"`
	Cape         ItemSlotEquip `bson:"cape"`
	SubWeapon    ItemSlotEquip `bson:"sub_weapon"`
	Weapon       ItemSlotEquip `bson:"weapon"`
	Ring1        ItemSlotEquip `bson:"ring1"`
	Ring2        ItemSlotEquip `bson:"ring2"`
	Ring3        ItemSlotEquip `bson:"ring3"`
	Ring4        ItemSlotEquip `bson:"ring4"`
	Pendant      ItemSlotEquip `bson:"pendant"`
	TamingMob    ItemSlotEquip `bson:"taming_mob"`
	Saddle       ItemSlotEquip `bson:"saddle"`
	MobEquip     ItemSlotEquip `bson:"mob_equip"`
	Medal        ItemSlotEquip `bson:"medal"`
	Belt         ItemSlotEquip `bson:"belt"`
	Shoulder     ItemSlotEquip `bson:"shoulder"`
	Pocket       ItemSlotEquip `bson:"pocket"`
	Badge        ItemSlotEquip `bson:"badge"`
	Emblem       ItemSlotEquip `bson:"emblem"`
	Android      ItemSlotEquip `bson:"android"`
	AndroidHeart ItemSlotEquip `bson:"andorid_heart"`
	MonsterBook  ItemSlotEquip `bson:"monster_book"`
	PendantExt   ItemSlotEquip `bson:"pendant_ext"`
}

type CashEquip struct {
	Hat       ItemSlotEquip `bson:"hat"`
	FaceAcc   ItemSlotEquip `bson:"face_acc"`
	EyeAcc    ItemSlotEquip `bson:"eye_acc"`
	EarAcc    ItemSlotEquip `bson:"ear_acc"`
	Top       ItemSlotEquip `bson:"top"`
	Bottom    ItemSlotEquip `bson:"bottom"`
	Overall   ItemSlotEquip `bson:"overall"`
	Shoes     ItemSlotEquip `bson:"shoes"`
	Gloves    ItemSlotEquip `bson:"gloves"`
	Cape      ItemSlotEquip `bson:"cape"`
	SubWeapon ItemSlotEquip `bson:"sub_weapon"`
	Weapon    ItemSlotEquip `bson:"weapon"`
	Ring1     ItemSlotEquip `bson:"ring1"`
	Ring2     ItemSlotEquip `bson:"ring2"`
	Ring3     ItemSlotEquip `bson:"ring3"`
	Ring4     ItemSlotEquip `bson:"ring4"`
	Pendant   ItemSlotEquip `bson:"pendant"`
	Belt      ItemSlotEquip `bson:"belt"`
	Shoulder  ItemSlotEquip `bson:"shoulder"`
}

type Pet struct {
	ID       uint32 `bson:"id"`
	Name     string `bson:"name"`
	Acc      uint32 `bson:"acc"`
	Skill    uint32 `bson:"skill"`
	SkillExt uint32 `bson:"skill_ext"`
}

// Call by AvatarLook
func (inv *EquipInventory) GetEquipLook() []ItemSlotEquip {
	equipLook := make([]ItemSlotEquip, 0)
	// Hat
	if inv.CashEquip.Hat.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Hat)
	} else if inv.Equip.Hat.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Hat)
	}
	// FaceAcc
	if inv.CashEquip.FaceAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.FaceAcc)
	} else if inv.Equip.FaceAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.FaceAcc)
	}
	// EyeAcc
	if inv.CashEquip.EyeAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.EyeAcc)
	} else if inv.Equip.EyeAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.EyeAcc)
	}
	// EarAcc
	if inv.CashEquip.EarAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.EarAcc)
	} else if inv.Equip.EarAcc.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.EarAcc)
	}
	// Clothes
	if inv.CashEquip.Overall.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Overall)
	} else if inv.Equip.Overall.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Overall)
	} else if inv.CashEquip.Top.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Top)
	} else if inv.Equip.Top.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Top)
	}
	// Bottom
	if inv.CashEquip.Overall.BagIndex == 0 {
		if inv.CashEquip.Bottom.BagIndex > 0 {
			equipLook = append(equipLook, inv.CashEquip.Bottom)
		} else if inv.Equip.Bottom.BagIndex > 0 {
			equipLook = append(equipLook, inv.Equip.Bottom)
		}
	}
	// Shoes
	if inv.CashEquip.Shoes.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Shoes)
	} else if inv.Equip.Shoes.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Shoes)
	}
	// Gloves
	if inv.CashEquip.Gloves.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Gloves)
	} else if inv.Equip.Gloves.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Gloves)
	}
	// Cape
	if inv.CashEquip.Cape.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Cape)
	} else if inv.Equip.Cape.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Cape)
	}
	// SubWeapon
	if inv.CashEquip.SubWeapon.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.SubWeapon)
	} else if inv.Equip.SubWeapon.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.SubWeapon)
	}
	// Weapon
	if inv.CashEquip.Weapon.BagIndex > 0 {
		equipLook = append(equipLook, inv.CashEquip.Weapon)
	} else if inv.Equip.Weapon.BagIndex > 0 {
		equipLook = append(equipLook, inv.Equip.Weapon)
	}
	return equipLook
}

func getItemSlotEquips(temp any) []ItemSlotEquip {
	v := reflect.ValueOf(temp)
	equips := make([]ItemSlotEquip, 0)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		equip, ok := field.Interface().(ItemSlotEquip)
		if !ok {
			continue
		}
		if equip.BagIndex == 0 || equip.ItemID == 0 {
			continue
		}
		equips = append(equips, equip)
	}
	return equips
}

func (inv *EquipInventory) GetEquips() []ItemSlotEquip {
	return getItemSlotEquips(inv.Equip)
}

func (inv *EquipInventory) GetCashEquips() []ItemSlotEquip {
	return getItemSlotEquips(inv.CashEquip)
}
