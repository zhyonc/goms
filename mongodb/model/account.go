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
	GMLevel              uint8     `bson:"gm_level"`
	GradeCode            uint8     `bson:"grade_code"`
	PurchaseExp          uint8     `bson:"purchase_exp"`
	ShiningStarCount     uint32    `bson:"shining_star_count"`
	IsForeverBanned      bool      `bson:"is_forever_banned"`
	TempBannedExpireDate time.Time `bson:"temp_banned_expire_date"`
	ChatUnblockDate      time.Time `bson:"chat_unblock_date"`
	CashPoint            uint32    `bson:"cash_point"`
	MaplePoint           uint32    `bson:"maple_point"`
	RegisterIP           string    `bson:"register_ip"`
	RegisterMAC          string    `bson:"register_mac"`
	RegisterDate         time.Time `bson:"register_date"`
	LoginIP              string    `bson:"login_ip"`
	LoginMAC             string    `bson:"login_mac"`
	LoginDate            time.Time `bson:"login_date"`
	UpdateDate           time.Time `bson:"update_date"`
}

func NewAccount(id uint32, username, password, ip, mac string) *Account {
	return &Account{
		ID:          id,
		Username:    username,
		Password:    password,
		RegisterIP:  ip,
		RegisterMAC: mac,
	}
}
