package character

import (
	"goms/packet/inpacket"
	"time"
)

type Character struct {
	ID               uint32           `bson:"_id"`
	AccountID        uint32           `bson:"account_id"`
	WorldID          uint8            `bson:"world_id"`
	Name             string           `bson:"name"`
	Gender           bool             `bson:"gender"`
	JobClass         uint32           `bson:"job_class"`
	Job              uint16           `bson:"job"`
	SubJob           uint16           `bson:"sub_job"`
	TotalCHUC        uint32           `bson:"total_chuc"`
	WeaponPoint      uint32           `bson:"weapon_point"`
	GachaponExp      uint32           `bson:"gachapon_exp"`
	PosMap           uint32           `bson:"pos_map"`
	Portal           uint8            `bson:"portal"`
	IsBurning        bool             `bson:"is_burining"`
	BurningStartDate time.Time        `bson:"burning_start_date"`
	BurningEndDate   time.Time        `bson:"burning_end_date"`
	BurningType      uint8            `bson:"burning_type"`
	IsReserved       bool             `bson:"is_reserved"`
	ReservedDate     time.Time        `bson:"reserved_date"`
	IsDeleted        bool             `bson:"is_deleted"`
	DeletedDate      time.Time        `bson:"deleted_date"`
	IsRenamed        bool             `bson:"is_renamed"`
	CreateDate       time.Time        `bson:"create_date"`
	LoginDate        time.Time        `bson:"login_date"`
	UpdateDate       time.Time        `bson:"update_date"`
	Keymap           CharacterKeymap  `bson:"keymap"`
	Stat             CharacterStat    `bson:"stat"`
	Look             CharacterLook    `bson:"look"`
	Skill            CharacterSkill   `bson:"skill"`
	Trait            CharacterTrait   `bson:"trait"`
	PVP              CharacterPVP     `bson:"pvp"`
	Cards            [9]CharacterCard `bson:"cards"`
	TamingMob        TamingMob        `bson:"taming_mob"`
}

func NewCharacter(characterID, accountID uint32, worldID uint8, in *inpacket.CharPacket) *Character {
	char := &Character{
		ID:        characterID,
		AccountID: accountID,
		WorldID:   worldID,
		Name:      in.CharacterName,
		Gender:    in.Gender,
		JobClass:  in.JobClass,
		Job:       in.Job,
		Keymap:    CharacterKeymap{},
		Stat:      CharacterStat{},
		Look:      CharacterLook{},
		Skill:     CharacterSkill{},
		Trait:     CharacterTrait{},
		PVP:       CharacterPVP{},
		TamingMob: TamingMob{},
	}
	// Keymap
	char.Keymap.KeySettingType = in.KeySettingType
	// Inital Look
	char.Look.SkinColor = in.SkinColor
	char.Look.Face = in.Face
	char.Look.Hair = in.Hair
	char.Look.DefFaceAcc = in.DefFaceAcc
	// Inital stat
	char.PosMap = 4000020
	char.Stat.Level = 1
	char.Stat.HP = 50
	char.Stat.MaxHP = 50
	char.Stat.MP = 50
	char.Stat.MaxMP = 50
	char.Stat.Str = 4
	char.Stat.Dex = 4
	char.Stat.Int = 4
	char.Stat.Luk = 4
	return char
}
