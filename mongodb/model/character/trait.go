package character

import "time"

type CharacterTrait struct {
	Fatigue                 uint8     `bson:"fatigue"`
	LastFatigueUpdateTime   time.Time `bson:"last_fatigue_update_time"`
	CharismaLevel           uint8     `bson:"charisma_level"`
	CharismaExp             uint16    `bson:"charisma_exp"`
	InsightLevel            uint8     `bson:"insight_level"`
	InsightExp              uint16    `bson:"insight_exp"`
	WillLevel               uint8     `bson:"will_level"`
	WillExp                 uint16    `bson:"will_exp"`
	CraftLevel              uint8     `bson:"craft_level"`
	CraftExp                uint16    `bson:"craft_exp"`
	SenseLevel              uint16    `bson:"sense_level"`
	SenseExp                uint16    `bson:"sense_exp"`
	CharmLevel              uint8     `bson:"charm_level"`
	CharmExp                uint16    `bson:"charm_exp"`
	CharmByCashPR           uint8     `bson:"charm_by_cash_pr"`
	LastUpdateCharmByCashPR time.Time `bson:"last_update_charm_by_cash_pr"`
}
