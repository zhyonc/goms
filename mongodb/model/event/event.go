package event

import "time"

type Event struct {
	ID         uint32    `bson:"_id"`
	UpdateDate time.Time `bson:"update_date"`
}
