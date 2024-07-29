package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterOrder struct {
	ID           primitive.ObjectID `bson:"_id"`
	AccountID    uint32             `bson:"account_id"`
	WorldID      uint8              `bson:"world_id"`
	CharacterIDs []uint32           `bson:"character_ids"`
	UpdateDate   time.Time          `bson:"update_date"`
}

func NewCharacterOrder(accountID uint32, worldID uint8) *CharacterOrder {
	charOrder := &CharacterOrder{
		AccountID:    accountID,
		WorldID:      worldID,
		CharacterIDs: make([]uint32, 0),
	}
	return charOrder
}
