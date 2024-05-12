package character

type CharacterSkill struct {
	SP        uint16     `bson:"sp"`
	ExtendSPs []ExtendSP `bson:"extend_sps"`
}

type ExtendSP struct {
	JobLevel uint8  `bson:"job_level"`
	JobSP    uint32 `bson:"job_sp"`
}
