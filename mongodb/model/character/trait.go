package character

import "time"

type CharacterTrait struct {
	Fatigue               uint16    `bson:"fatigue"`
	LastFatigueUpdateTime time.Time `bson:"last_fatigue_update_time"`
	CharismaLevel         uint8     `bson:"charisma_level"`
	CharismaExp           uint32    `bson:"charisma_exp"`
	InsightLevel          uint8     `bson:"insight_level"`
	InsightExp            uint32    `bson:"insight_exp"`
	WillLevel             uint8     `bson:"will_level"`
	WillExp               uint32    `bson:"will_exp"`
	CraftLevel            uint8     `bson:"craft_level"`
	CraftExp              uint32    `bson:"craft_exp"`
	SenseLevel            uint32    `bson:"sense_level"`
	SenseExp              uint32    `bson:"sense_exp"`
	CharmLevel            uint8     `bson:"charm_level"`
	CharmExp              uint32    `bson:"charm_exp"`
}
