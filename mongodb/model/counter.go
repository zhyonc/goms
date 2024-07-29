package model

type Counter struct {
	AccountID   uint32 `bson:"account_id"`
	CharacterID uint32 `bson:"character_id"`
}
