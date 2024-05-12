package model

import (
	"time"
)

type Account struct {
	ID                   uint32    `bson:"_id"`
	Username             string    `bson:"username"`
	Password             string    `bson:"password"`
	SecondPassword       string    `bson:"second_password"`
	Gender               bool      `bson:"gender"`
	GradeCode            uint8     `bson:"grade_code"`
	IsGM                 bool      `bson:"is_gm"`
	GMLevel              uint8     `bson:"gm_level"`
	IsForeverBanned      bool      `bson:"is_forever_banned"`
	BannedReason         string    `bson:"banned_reson"`
	TempBannedExpireTime time.Time `bson:"temp_banned_expire_time"`
	ChatUnblockTime      time.Time `bson:"chat_unblock_time"`
	RegisterMAC          string    `bson:"register_mac"`
	RegisterTime         time.Time `bson:"register_time"`
	LoginMAC             string    `bson:"login_mac"`
	LoginTime            time.Time `bson:"login_time"`
	UpdateTime           time.Time `bson:"update_time"`
}

func NewAccount(id uint32, username, password, mac string) *Account {
	return &Account{
		ID:           id,
		Username:     username,
		Password:     password,
		RegisterMAC:  mac,
		RegisterTime: time.Now(),
	}
}
