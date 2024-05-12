package character

import (
	"goms/maple/class"
	"goms/maple/world"
	"goms/packet/inpacket"
	"time"
)

type Character struct {
	ID               uint32          `bson:"_id"`
	AccountID        uint32          `bson:"account_id"`
	WorldID          uint8           `bson:"world_id"`
	Name             string          `bson:"name"`
	Gender           bool            `bson:"gender"`
	KeySettingType   uint8           `bson:"key_setting_type"`
	Class            uint32          `bson:"class"`
	Job              uint16          `bson:"job"`
	SubJob           uint16          `bson:"sub_job"`
	WeaponPoint      uint32          `bson:"weapon_point"`
	GachExp          uint64          `bson:"gach_exp"`
	PlayTimeUnix     uint64          `bson:"play_time_unix"`
	MapID            uint32          `bson:"map_id"`
	Portal           uint8           `bson:"portal"`
	IsBurning        bool            `bson:"is_burining"`
	BurningStartTime time.Time       `bson:"burning_start_time"`
	BurningEndTime   time.Time       `bson:"burning_end_time"`
	BurningType      uint8           `bson:"burning_type"`
	IsDeleted        bool            `bson:"is_deleted"`
	FutureDeleteTime time.Time       `bson:"future_delete_time"`
	IsRenamed        bool            `bson:"is_renamed"`
	Look             *CharacterLook  `bson:"look"`
	EquipInventory   *EquipInventory `bson:"equip_inventory"`
	ItemInventory    *ItemInventory  `bson:"item_inventory"`
	Stat             *CharacterStat  `bson:"stat"`
	Skill            *CharacterSkill `bson:"skill"`
	Trait            *CharacterTrait `bson:"trait"`
	PVP              *CharacterPVP   `bson:"pvp"`
}

func NewCharacter(characterID, accountID uint32, worldID world.WorldID, in *inpacket.CharPacket) *Character {
	char := &Character{
		ID:             characterID,
		AccountID:      accountID,
		WorldID:        uint8(worldID),
		Name:           in.CharacterName,
		Gender:         in.Gender,
		KeySettingType: uint8(in.KeySettingType),
		Class:          in.Class,
		Job:            in.Job,
		Look:           &CharacterLook{},
		EquipInventory: &EquipInventory{},
		ItemInventory:  &ItemInventory{},
		Stat:           &CharacterStat{},
		Skill:          &CharacterSkill{},
		Trait:          &CharacterTrait{},
		PVP:            &CharacterPVP{},
	}
	// Inital Look
	char.Look.SkinColor = in.SkinColor
	char.Look.Face = in.Face
	char.Look.Hair = in.Hair
	char.Look.SpecialFace = in.SpecialFace
	// Inital equip
	char.EquipInventory.Equip.Hat = in.Hat
	char.EquipInventory.Equip.Top = in.Top
	char.EquipInventory.Equip.Bottom = in.Bottom
	char.EquipInventory.Equip.Overall = in.Overall
	char.EquipInventory.Equip.Cape = in.Cape
	char.EquipInventory.Equip.Shoes = in.Shoes
	char.EquipInventory.Equip.Gloves = in.Gloves
	char.EquipInventory.Equip.Weapon = in.Weapon
	char.EquipInventory.Equip.SubWeapon = in.SubWeapon
	// Inital stat
	if char.Class == uint32(class.Adventures) {
		char.MapID = 10000
		char.Stat.Level = 1
		char.Stat.HP = 50
		char.Stat.MaxHP = 50
		char.Stat.MP = 50
		char.Stat.MaxMP = 50
		char.Stat.Str = 4
		char.Stat.Dex = 4
		char.Stat.Int = 4
		char.Stat.Luk = 4
	}
	return char
}
