package inventory

import "time"

type ItemPotInventory struct {
	Pots [3]ItemPotSlot `bson:"pots"`
}

// GW_ItemPotSlot::RawDecode
type ItemPotSlot struct {
	LifeID             uint32    `bson:"lift_id"`
	Level              uint8     `bson:"level"`
	LastState          uint8     `bson:"last_state"`
	Satiety            uint32    `bson:"satiety"`
	Friendly           uint32    `bson:"friendly"`
	RemainAbleFriendly uint32    `bson:"remain_Able_friendly"`
	RemainFriendlyTime uint32    `bson:"remain_friendly_time"`
	MaximumIncLevel    uint8     `bson:"maximum_inc_level"`
	MaximumIncSatiety  uint32    `bson:"maximum_inc_satiety"`
	LastEatTime        time.Time `bson:"last_eat_time"`
	LastSleepStartTime time.Time `bson:"last_sleep_start_time"`
	LastDecSatietyTime time.Time `bson:"last_dec_satiety_time"`
	ConsumedLastTime   time.Time `bson:"consumed_last_time"`
}
