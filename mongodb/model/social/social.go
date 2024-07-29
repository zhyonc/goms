package social

import "time"

type Social struct {
	ID         uint32    `bson:"_id"`
	Friends    []Friend  `bson:"friends"`
	Party      Party     `bson:"party"`
	Guild      Guild     `bson:"guild"`
	UpdateDate time.Time `bson:"update_date"`
}

func NewSocial(characterID uint32) *Social {
	soc := &Social{
		ID:         characterID,
		Friends:    make([]Friend, 0),
		Party:      Party{},
		Guild:      Guild{},
		UpdateDate: time.Now(),
	}
	return soc
}
