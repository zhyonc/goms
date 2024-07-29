package usecase

import (
	"context"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
	"time"
)

// db collection name
const (
	counterCollName        string = "counter"
	accountCollName        string = "accounts"
	characterOrderCollName string = "character_orders"
	characterCollName      string = "characters"
	inventoryCollName      string = "inventories"
	socialCollName         string = "socials"
)

type CounterUsecase interface {
	GetAccountID(ctx context.Context) uint32
	GetCharacterID(ctx context.Context) uint32
}

type AccountUsecase interface {
	FindAccountByUsername(ctx context.Context, username string) *model.Account
	CreateNewAccount(ctx context.Context, account *model.Account) bool
	UpdatePassword(ctx context.Context, accountID uint32, isSecondPassword bool, passwrod string) bool
	UpdateGenderAndSecondPassword(ctx context.Context, username string, secondPassWord string, gender bool) bool
	UpdateLoginRecord(ctx context.Context, accountID uint32, ip, mac string)
	FindAccountByID(ctx context.Context, accountID uint32) *model.Account
}

type CharacterOrderUsecase interface {
	CreateNewCharacterOrder(ctx context.Context, charOrder *model.CharacterOrder) bool
	FindCharacterOrder(ctx context.Context, accountID uint32, worldID uint8) *model.CharacterOrder
	UpdateCharacterOrder(ctx context.Context, charOrder *model.CharacterOrder) bool
}

type CommonUsecase interface {
	CreateNewModel(ctx context.Context, model any) bool
	FindModelByID(ctx context.Context, id uint32) any
	UpdateModelByID(ctx context.Context, id uint32, model any) bool
}

type CharacterUsecase interface {
	CommonUsecase
	FindCharacterName(ctx context.Context, characterName string) bool
	FindCharactersByAccountID(ctx context.Context, accountID uint32, worldID uint8) []*character.Character
	FindCharacterCount(ctx context.Context, accountID uint32, worldID uint8) uint8
	CreateNewCharacter(ctx context.Context, char *character.Character) bool
	ReserveCharacter(ctx context.Context, characterID uint32, reservedDate time.Time) bool
	DeleteCharacter(ctx context.Context, characterID uint32) bool
	RestoreCharacter(ctx context.Context, characterID uint32) bool
	UpdateLoginDate(ctx context.Context, characterID uint32)
}

type InventoryUsecase interface {
	CommonUsecase
}

type SocialUsecase interface {
	CommonUsecase
}

type EventUsecase interface {
	CommonUsecase
}
