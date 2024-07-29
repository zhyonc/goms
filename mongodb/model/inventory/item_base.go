package inventory

import (
	"goms/maple"
	"goms/nxfile"
	"time"
)

// GW_ItemSlotBase::GW_ItemSlotBase
type ItemBase struct {
	ItemID     uint32    `bson:"item_id"`
	DateExpire time.Time `bson:"date_expire"`
	BagIndex   uint16    `bson:"bag_index"`
	CashItemSN uint64    `bson:"cash_item_sn"`
	IsCash     bool      `bson:"is_cash"`
}

// GW_ItemSlotBundle::GW_ItemSlotBundle
// ItemType Bundle
type ItemSlotBundle struct {
	ItemBase
	Number    uint16 `bson:"number"`
	Attribute uint16 `bson:"attribute"`
	SN        uint64 `bson:"sn"`
	Title     string `bson:"title"`
}

// GW_ItemSlotPet::GW_ItemSlotPet
// ItemType Pet
type ItemSlotPet struct {
	ItemBase      `bson:"base"`
	PetName       string    `bson:"pet_name"`        // sPetName
	Level         uint8     `bson:"level"`           // nLevel
	Tameness      uint16    `bson:"tameness"`        // nTameness
	Repleteness   uint8     `bson:"repleteness"`     // nRepletenes
	PetAttribute  uint16    `bson:"pet_attribute"`   // nPetAttribute
	PetSkill      uint16    `bson:"pet_skill"`       // usPetSkill
	DateDead      time.Time `bson:"date_dead"`       // dateDead
	RemainLife    uint32    `bson:"remain_life"`     // nRemainLife
	Attribute     uint16    `bson:"attribute"`       // nAttribut
	ActiveState   uint8     `bson:"active_state"`    // nActiveState
	AutoBuffSkill uint32    `bson:"auto_buff_skill"` // nAutoBuffSkill
	PetHue        uint32    `bson:"pet_hue"`         // nPetHue
	GiantRate     uint16    `bson:"giant_rate"`      // nGiantRate
}

// GW_ItemSlotEquip::GW_ItemSlotEquip
// ItemType Equip
type ItemSlotEquip struct {
	ItemBase         `bson:"base"`
	EquippedDate     time.Time         `bson:"equipped_date"`       // ftEquipped
	PrevBonusExpRate int32             `bson:"prev_bonus_exp_rate"` // nPrevBonusExpRate
	Stat             ItemSlotEquipBase `bson:"stat"`
	CashOpt          CashItemOption    `bson:"cash_opt"`
	ExtOpt           ItemSlotEquipOpt  `bson:"ext_opt"`
	SN               uint64            `bson:"sn"`    // liSN
	Title            string            `bson:"title"` // sTitle
}

func NewItemSlotEquip(id uint32, bagIndex uint16) ItemSlotEquip {
	equip := ItemSlotEquip{}
	cache := nxfile.GetEquip(id)
	if cache == nil {
		return equip
	}
	equip.ItemID = id
	equip.BagIndex = bagIndex
	equip.IsCash = cache.Cash
	equip.EquippedDate = time.Now()
	equip.PrevBonusExpRate = cache.BonusExp // unsure
	equip.Stat.RUC = cache.Tuc
	equip.Stat.STR = cache.IncSTR
	equip.Stat.DEX = cache.IncDEX
	equip.Stat.LUK = cache.IncLUK
	equip.Stat.MaxHP = cache.IncMHP
	equip.Stat.MaxMP = cache.IncMMP
	equip.Stat.PAD = cache.IncPAD
	equip.Stat.MAD = cache.IncMAD
	equip.Stat.PDD = cache.IncPDD
	equip.Stat.MDD = cache.IncMDD
	equip.Stat.ACC = cache.IncACC
	equip.Stat.EVA = cache.IncEVA
	equip.Stat.Craft = cache.IncCraft
	equip.Stat.Speed = cache.IncSpeed
	equip.Stat.Jump = cache.IncJump
	equip.Stat.Attribute = cache.FAttribute
	equip.Stat.Level = cache.Level
	equip.Stat.Durability = cache.Durability
	equip.Stat.PVPDamage = cache.IncPVPDamage
	equip.Stat.ReduceReq = cache.ReduceReq
	equip.Stat.BDR = cache.BdR
	equip.Stat.Cuttable = cache.CuttableCount
	equip.Stat.ExGradeOption = cache.ExGrade
	return equip
}

func (i *ItemSlotEquip) PutOn(bagIndex uint16) {
	i.BagIndex = bagIndex
}

func (i *ItemSlotEquip) TakeOff() {
	i.BagIndex = 0
}

// GW_ItemSlotEquipBase::GW_ItemSlotEquipBase
type ItemSlotEquipBase struct {
	RUC              int8  `bson:"ruc"`               // nRUC is Remaining Upgrade Counts
	CUC              int8  `bson:"cuc"`               // nCUC is Current Upgrade Count
	STR              int16 `bson:"str"`               // niSTR
	DEX              int16 `bson:"dex"`               // niDEX
	INT              int16 `bson:"int"`               // niINT
	LUK              int16 `bson:"luk"`               // niLUK
	MaxHP            int16 `bson:"max_hp"`            // niMaxHP
	MaxMP            int16 `bson:"max_mp"`            // niMaxMP
	PAD              int16 `bson:"pad"`               // niPAD
	MAD              int16 `bson:"mad"`               // niMAD
	PDD              int16 `bson:"pdd"`               // niPDD
	MDD              int16 `bson:"mdd"`               // niMDD
	ACC              int16 `bson:"acc"`               // niACC
	EVA              int16 `bson:"eva"`               // niEVA
	Craft            int16 `bson:"craft"`             // niCraft
	Speed            int16 `bson:"speed"`             // niSpeed
	Jump             int16 `bson:"jump"`              // niJump
	Attribute        int32 `bson:"attribute"`         // nAttribute
	LevelUpType      int8  `bson:"level_up_type"`     // nLevelUpType
	Level            int8  `bson:"level"`             // nLevel
	EXP              int64 `bson:"exp"`               // nEXP64 10000000=100%?
	Durability       int32 `bson:"durablity"`         // nDurability
	IUC              int32 `bson:"iuc"`               // nIUC is Inc Upgrade count by Golden Hammer
	PVPDamage        int16 `bson:"pvp_damage"`        // niPVPDamage
	ReduceReq        int8  `bson:"reduce_req"`        // niReduceReq
	SpecialAttribute int16 `bson:"special_attribute"` // nSpecialAttribute
	DurabilityMax    int32 `bson:"durability_max"`    // nDurabilityMax
	IncReq           int8  `bson:"inc_req"`           // niIncReq
	GrowthEnchant    int8  `bson:"growth_enchant"`    // nGrowthEnchant
	PSEnchant        int8  `bson:"ps_enchant"`        // nPSEnchant
	BDR              int8  `bson:"bdr"`               // nBDR
	IMDR             int8  `bson:"imdr"`              // nIMDR
	DamR             int8  `bson:"dam_r"`             // nDamR
	StatR            int8  `bson:"stat_r"`            // nStatR
	Cuttable         int8  `bson:"cuttable"`          // nCuttable
	ExGradeOption    int64 `bson:"ex_grage_option"`   // nExGradeOption
	ItemState        int32 `bson:"item_state"`        // nItemState
}

func (i *ItemSlotEquipBase) GetStatMask() maple.EquipStat {
	var mask maple.EquipStat = 0
	if i.RUC > 0 {
		mask |= maple.RUC
	}
	if i.CUC > 0 {
		mask |= maple.CUC
	}
	if i.STR > 0 {
		mask |= maple.STR
	}
	if i.DEX > 0 {
		mask |= maple.DEX
	}
	if i.INT > 0 {
		mask |= maple.INT
	}
	if i.LUK > 0 {
		mask |= maple.LUK
	}
	if i.MaxHP > 0 {
		mask |= maple.MaxHP
	}
	if i.MaxMP > 0 {
		mask |= maple.MaxMP
	}
	if i.PAD > 0 {
		mask |= maple.PAD
	}
	if i.MAD > 0 {
		mask |= maple.MAD
	}
	if i.PDD > 0 {
		mask |= maple.PDD
	}
	if i.MDD > 0 {
		mask |= maple.MDD
	}
	if i.ACC > 0 {
		mask |= maple.ACC
	}
	if i.EVA > 0 {
		mask |= maple.EVA
	}
	if i.Craft > 0 {
		mask |= maple.Craft
	}
	if i.Speed > 0 {
		mask |= maple.Speed
	}
	if i.Jump > 0 {
		mask |= maple.Jump
	}
	if i.Attribute > 0 {
		mask |= maple.Attribute
	}
	if i.LevelUpType > 0 {
		mask |= maple.LevelUpType
	}
	if i.Level > 0 {
		mask |= maple.Level
	}
	if i.EXP > 0 {
		mask |= maple.EXP
	}
	if i.Durability > 0 {
		mask |= maple.Durability
	}
	if i.IUC > 0 {
		mask |= maple.IUC
	}
	if i.PVPDamage > 0 {
		mask |= maple.PVPDamage
	}
	if i.ReduceReq > 0 {
		mask |= maple.ReduceReq
	}
	if i.SpecialAttribute > 0 {
		mask |= maple.SpecialAttribute
	}
	if i.DurabilityMax > 0 {
		mask |= maple.DurabilityMax
	}
	if i.IncReq > 0 {
		mask |= maple.IncReq
	}
	if i.GrowthEnchant > 0 {
		mask |= maple.GrowthEnchant
	}
	if i.PSEnchant > 0 {
		mask |= maple.PSEnchant
	}
	if i.BDR > 0 {
		mask |= maple.BDR
	}
	if i.IMDR > 0 {
		mask |= maple.MDR
	}
	return mask
}

func (i *ItemSlotEquipBase) GetSubStatMask() maple.EquipStat {
	var mask maple.EquipStat = 0
	if i.DamR > 0 {
		mask |= maple.DamR
	}
	if i.StatR > 0 {
		mask |= maple.StatR
	}
	if i.Cuttable > 0 {
		mask |= maple.Cuttable
	}
	if i.ExGradeOption > 0 {
		mask |= maple.ExGradeOption
	}
	if i.ItemState > 0 {
		mask |= maple.ItemState
	}
	return mask
}

// GW_CashItemOption::GW_CashItemOption
type CashItemOption struct {
	CashItemSN uint64    `bson:"cash_item_sn"` // liCashItemSN.QuadPart
	Options    [3]uint32 `bson:"options"`      // anOption
	ExpireDate time.Time `bson:"expire_date"`  // ftExpireDate
	Grade      uint32    `bson:"grade"`        // nGrade
}

// GW_ItemSlotEquipOpt::GW_ItemSlotEquipOpt
type ItemSlotEquipOpt struct {
	Grade            uint8     `bson:"grade"`              // nGrade is potential grade
	CHUC             uint8     `bson:"chuc"`               // nCHUC is equip star num
	Options          [7]uint16 `bson:"options"`            // nOption1-6 for potential with cube and last is fusion anvil
	SocketState      uint16    `bson:"socket_state"`       // Socket mask
	Sockets          [3]uint16 `bson:"sockets"`            // Socket item id
	SoulOptionID     uint16    `bson:"soul_option_id"`     // nSoulOptionID
	SoulSocketID     uint16    `bson:"soul_socket_id"`     // nSoulSocketID
	SoulOption       uint16    `bson:"soul_option"`        // nSoulOption
	DamageLimitBreak uint32    `bson:"damage_limit_break"` // CMS Feature
}
