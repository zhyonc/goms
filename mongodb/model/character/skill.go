package character

import "time"

type CharacterSkill struct {
	SP        uint16               `bson:"sp"`
	ExtendSPs []ExtendSP           `bson:"extend_sps"`
	Lists     []Skill              `bson:"lists"`
	Links     []LinkSkill          `bson:"links"`
	CoolTimes map[uint32]time.Time `bson:"cool_times"`
}

func (c *CharacterSkill) GetSkill(skillID uint32) Skill {
	for _, skill := range c.Lists {
		if skill.ID == skillID {
			return skill
		}
	}
	return Skill{}
}

type ExtendSP struct {
	JobLevel uint8  `bson:"job_level"`
	JobSP    uint32 `bson:"job_sp"`
}

type Skill struct {
	ID             uint32    `bson:"id"`
	CurrentLevel   uint32    `bson:"current_level"`
	MaxLevel       uint16    `bson:"max_level"`
	ExpirationTime time.Time `bson:"expiration_time"`
	MasterLevel    uint32    `bson:"master_level"`
}

type LinkSkill struct {
	ID          uint32 `bson:"id"`
	OwnerID     uint32 `bson:"owner_id"`
	LinkSkillID uint32 `bson:"link_skill_id"`
	Level       uint16 `bson:"level"`
}
