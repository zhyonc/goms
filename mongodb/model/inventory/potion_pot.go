package inventory

import "time"

type PotionPotInventory struct {
	Pots []PotionPot `bson:"pots"`
}

type PotionPot struct {
	ItemBase
	MaxCapcity uint32    `bson:"max_capcity"`
	HP         uint32    `bson:"hp"`
	MP         uint32    `bson:"mp"`
	StartTime  time.Time `bson:"start_time"`
	EndTime    time.Time `bson:"end_time"`
}
