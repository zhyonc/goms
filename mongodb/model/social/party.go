package social

type Party struct {
	Name      string         `bson:"name"`
	Appliable bool           `bson:"appliable"`
	LeaderID  uint32         `bson:"leader_id"`
	Members   [6]PartyMember `bson:"members"`
}

type PartyMember struct {
	CharacterID uint32 `bson:"character_id"` // member id
	Name        string `bson:"name"`         // member name
	Level       int32  `bson:"level"`
	Job         int32  `bson:"job"`
	SubJob      int32  `bson:"sub_job"`
	MapID       uint32 `bson:"map_id"`
	TownID      uint32 `bson:"town_id"`
	TargetID    uint32 `bson:"target_id"`
	SkillID     uint32 `bson:"skill_id"`
	PosX        int32  `bson:"pos_x"`
	PosY        int32  `bson:"pos_y"`
}
