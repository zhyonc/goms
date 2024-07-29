package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/inventory"
)

// Call by NonBPEquip::Decode
// Call by ReturnEffectInfo::Decode
// Call by MemorialCubeInfo::Decode
// Call by CWvsContext::OnBroadcastMsg
// GW_ItemSlotBase::Decode
func GWItemSlotBaseEncode(p *outPacket, itemType maple.ItemType, item any) {
	p.EncodeByte(byte(itemType))
	// I think GW_ItemSlotBase::CreateItem return the entity class based on itemType
	// And then call RawDecode functon in GW_ItemSlotBase::Decode
	switch itemType {
	case maple.ItemTypeEquip:
		itemEquip, _ := item.(inventory.ItemSlotEquip)
		GWItemSlotEquipRawEncode(p, &itemEquip)
	case maple.ItemTypeBundle:
		itemBundle, _ := item.(inventory.ItemSlotBundle)
		GWItemSlotBundleRawEncode(p, &itemBundle)
	case maple.ItemTypePet:
		itemPet, _ := item.(inventory.ItemSlotPet)
		GWItemSlotPetRawEncode(p, &itemPet)
	}
}

// Call by GW_ItemSlotEquip::RawDecode
// Call by GW_ItemSlotBundle::RawDecode
// Call by GW_ItemSlotPet::RawDecode
// GW_ItemSlotBase::RawDecode
func GWItemSlotBaseRawEncode(p *outPacket, itemBase inventory.ItemBase) {
	p.EncodeUint32(itemBase.ItemID) // nItemID
	p.EncodeBool(itemBase.CashItemSN > 0)
	if itemBase.CashItemSN > 0 {
		p.EncodeUint64(itemBase.CashItemSN)
	}
	p.EncodeFT(itemBase.DateExpire) // dateExpire
	p.EncodeInt32(-1)               // nBagIndex if it's in a bag?
}

// Call by GW_ItemSlotBase::Decode
// GW_ItemSlotEquip::RawDecode
func GWItemSlotEquipRawEncode(p *outPacket, itemEquip *inventory.ItemSlotEquip) {
	// ItemSlotBase
	GWItemSlotBaseRawEncode(p, itemEquip.ItemBase)
	// ItemSlotEquipBase is equip stat
	GWItemSlotEquipBaseEncode(p, itemEquip)
	p.Fill(8) // CMS138 confusing padding
	// ItemSlotEquipOpt
	p.EncodeStr(itemEquip.Title)         // owner name
	p.EncodeByte(itemEquip.ExtOpt.Grade) // nGrade is potential grade
	p.EncodeByte(itemEquip.ExtOpt.CHUC)  // nCHUC is equip star num
	for i := 0; i < 6; i++ {
		p.EncodeUint16(itemEquip.ExtOpt.Options[i]) // nOption1-6 for potential with cube
	}
	p.EncodeUint16(itemEquip.ExtOpt.Options[6]) // The last is fusion anvil
	//  New
	p.EncodeUint16(itemEquip.ExtOpt.SocketState) // socket state, 0 = nothing, 0xFF = see loop?
	for i := 0; i < 3; i++ {
		p.EncodeUint16(itemEquip.ExtOpt.Sockets[i]) // sockets 0 through 2 (-1 = none, 0 = empty, >0 = filled?
	}
	// ItemSlotEquip
	if itemEquip.SN > 0 {
		p.EncodeUint64(itemEquip.SN)
	}
	p.EncodeFT(itemEquip.EquippedDate)
	p.EncodeInt32(itemEquip.PrevBonusExpRate) // nPrevBonusExpRate
	// CashItemOption
	GWCashItemOptionEncode(p, itemEquip)
	// ItemSlotEquipOpt
	p.EncodeUint16(itemEquip.ExtOpt.SoulOptionID)     // nSoulOptionID
	p.EncodeUint16(itemEquip.ExtOpt.SoulSocketID)     // nSoulSocketID
	p.EncodeUint16(itemEquip.ExtOpt.SoulOption)       // nSoulOption
	p.EncodeUint32(itemEquip.ExtOpt.DamageLimitBreak) // CMS Feature
}

// Call by GW_ItemSlotEquip::RawDecode
// GW_ItemSlotEquipBase::Decode
func GWItemSlotEquipBaseEncode(p *outPacket, itemEquip *inventory.ItemSlotEquip) {
	stat := itemEquip.Stat
	mask := stat.GetStatMask()
	p.EncodeUint32(uint32(mask))
	if mask&maple.RUC != 0 {
		p.EncodeInt8(stat.RUC)
	}
	if mask&maple.CUC != 0 {
		p.EncodeInt8(stat.CUC)
	}
	if mask&maple.STR != 0 {
		p.EncodeInt16(stat.STR)
	}
	if mask&maple.DEX != 0 {
		p.EncodeInt16(stat.DEX)
	}
	if mask&maple.INT != 0 {
		p.EncodeInt16(stat.INT)
	}
	if mask&maple.LUK != 0 {
		p.EncodeInt16(stat.LUK)
	}
	if mask&maple.MaxHP != 0 {
		p.EncodeInt16(stat.MaxHP)
	}
	if mask&maple.MaxMP != 0 {
		p.EncodeInt16(stat.MaxMP)
	}
	if mask&maple.PAD != 0 {
		p.EncodeInt16(stat.PAD)
	}
	if mask&maple.MAD != 0 {
		p.EncodeInt16(stat.MAD)
	}
	if mask&maple.PDD != 0 {
		p.EncodeInt16(stat.PDD)
	}
	if mask&maple.MDD != 0 {
		p.EncodeInt16(stat.MDD)
	}
	if mask&maple.ACC != 0 {
		p.EncodeInt16(stat.ACC)
	}
	if mask&maple.EVA != 0 {
		p.EncodeInt16(stat.EVA)
	}
	if mask&maple.Craft != 0 {
		p.EncodeInt16(stat.Craft)
	}
	if mask&maple.Speed != 0 {
		p.EncodeInt16(stat.Speed)
	}
	if mask&maple.Jump != 0 {
		p.EncodeInt16(stat.Jump)
	}
	if mask&maple.Attribute != 0 {
		p.EncodeInt32(stat.Attribute) // new uint16->uint32
	}
	if mask&maple.LevelUpType != 0 {
		p.EncodeInt8(stat.LevelUpType)
	}
	if mask&maple.Level != 0 {
		p.EncodeInt8(stat.Level)
	}
	if mask&maple.EXP != 0 {
		p.EncodeInt64(stat.EXP)
	}
	if mask&maple.Durability != 0 {
		p.EncodeInt32(stat.Durability)
	}
	if mask&maple.IUC != 0 {
		p.EncodeInt32(stat.IUC)
	}
	if mask&maple.PVPDamage != 0 {
		p.EncodeInt16(stat.PVPDamage)
	}
	if mask&maple.ReduceReq != 0 {
		p.EncodeInt8(stat.ReduceReq)
	}
	if mask&maple.SpecialAttribute != 0 {
		p.EncodeInt16(stat.SpecialAttribute)
	}
	if mask&maple.DurabilityMax != 0 {
		p.EncodeInt32(stat.DurabilityMax)
	}
	if mask&maple.IncReq != 0 {
		p.EncodeInt8(stat.IncReq)
	}
	if mask&maple.GrowthEnchant != 0 {
		p.EncodeInt8(stat.GrowthEnchant)
	}
	if mask&maple.PSEnchant != 0 {
		p.EncodeInt8(stat.PSEnchant)
	}
	if mask&maple.BDR != 0 {
		p.EncodeInt8(stat.BDR)
	}
	if mask&maple.MDR != 0 {
		p.EncodeInt8(stat.IMDR)
	}
	// New
	subMask := stat.GetSubStatMask()
	p.EncodeUint32(uint32(subMask))
	if subMask&maple.DamR != 0 {
		p.EncodeInt8(stat.DamR)
	}
	if subMask&maple.StatR != 0 {
		p.EncodeInt8(stat.StatR)
	}
	if subMask&maple.Cuttable != 0 {
		p.EncodeInt8(stat.Cuttable)
	}
	if subMask&maple.ExGradeOption != 0 {
		p.EncodeInt64(stat.ExGradeOption)
	}
	if subMask&maple.ItemState != 0 {
		p.EncodeInt32(stat.ItemState)
	}
}

// Call by GW_ItemSlotEquip::RawDecode
// GW_CashItemOption::Decode
func GWCashItemOptionEncode(p *outPacket, itemEquip *inventory.ItemSlotEquip) {
	p.EncodeUint64(itemEquip.CashOpt.CashItemSN) // liCashItemSN
	p.EncodeFT(itemEquip.CashOpt.ExpireDate)     // ftExpireDate
	p.EncodeUint32(itemEquip.CashOpt.Grade)      // nGrade
	for i := 0; i < 3; i++ {
		p.EncodeUint32(itemEquip.CashOpt.Options[i]) // anOption
	}
}

// Call by GW_ItemSlotBase::Decode
// GW_ItemSlotBundle::RawDecode
func GWItemSlotBundleRawEncode(p *outPacket, itemBundle *inventory.ItemSlotBundle) {
	// ItemSlotBase
	GWItemSlotBaseRawEncode(p, itemBundle.ItemBase)
	p.EncodeUint16(itemBundle.Number)
	p.EncodeStr(itemBundle.Title)
	p.EncodeUint16(itemBundle.Attribute)
	id := itemBundle.ItemID / 10000 // TSecType<long>::GetData(&this->baseclass_0.nItemID) / 10000;
	if id == 207 || id == 233 {
		p.EncodeUint64(itemBundle.SN)
	}
}

// Call by GW_ItemSlotBase::Decode
// GW_ItemSlotPet::RawDecode
func GWItemSlotPetRawEncode(p *outPacket, itemPet *inventory.ItemSlotPet) {
	// ItemSlotBase
	GWItemSlotBaseRawEncode(p, itemPet.ItemBase)
	p.EncodeLocalName(itemPet.PetName, maple.CharacterNameLength) // sPetName
	p.EncodeByte(itemPet.Level)                                   // nLevel
	p.EncodeUint16(itemPet.Tameness)                              // nTameness
	p.EncodeByte(itemPet.Repleteness)                             // nRepletenes
	p.EncodeFT(itemPet.DateDead)                                  // dateDead
	p.EncodeUint16(itemPet.PetAttribute)                          // nPetAttribute
	p.EncodeUint16(itemPet.PetSkill)                              // usPetSkill
	p.EncodeUint32(itemPet.RemainLife)                            // nRemainLife
	p.EncodeUint16(itemPet.Attribute)                             // nAttribute
	p.EncodeByte(itemPet.ActiveState)                             // nActiveState
	p.EncodeUint32(itemPet.AutoBuffSkill)                         // nAutoBuffSkill
	p.EncodeUint32(itemPet.PetHue)                                // nPetHue
	p.EncodeUint16(itemPet.GiantRate)                             // nGiantRate
}

// Call by CharacterData::Decode
// GW_ItemPotSlot::RawDecode
func GWItemPotSlotEncode(p *outPacket, itemPot *inventory.ItemPotSlot) {
	p.EncodeUint32(itemPot.LifeID)
	p.EncodeByte(itemPot.Level)
	p.EncodeByte(itemPot.LastState)
	p.EncodeUint32(itemPot.Satiety)
	p.EncodeUint32(itemPot.Friendly)
	p.EncodeUint32(itemPot.RemainAbleFriendly)
	p.EncodeUint32(itemPot.RemainFriendlyTime)
	p.EncodeByte(itemPot.MaximumIncLevel)
	p.EncodeUint32(itemPot.MaximumIncSatiety)
	p.EncodeFT(itemPot.LastEatTime)
	p.EncodeFT(itemPot.LastSleepStartTime)
	p.EncodeFT(itemPot.LastDecSatietyTime)
	p.EncodeFT(itemPot.ConsumedLastTime)
}
