package usecase

import (
	"context"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
)

// db collection name
const (
	counterColl   string = "counter"
	accountColl   string = "accounts"
	characterColl string = "characters"
)

type CounterUsecase interface {
	GetAccountID(ctx context.Context) uint32
	GetCharacterID(ctx context.Context) uint32
}

type AccountUsecase interface {
	FindAccountByUsername(ctx context.Context, username string) *model.Account
	CreateNewAccount(ctx context.Context, account *model.Account) bool
	UpdateGenderAndSecondPassword(ctx context.Context, username string, secondPassWord string, gender bool) bool
	UpdateLoginTime(ctx context.Context, accountID uint32, mac string)
	FindAccountByID(ctx context.Context, accountID uint32) *model.Account
}

type CharacterUsecase interface {
	FindCharacterName(ctx context.Context, characterName string) bool
	GetCharactersByAccountID(ctx context.Context, accountID uint32, worldID uint8) []*character.Character
	GetCharacterByID(ctx context.Context, characterID uint32) *character.Character
	CreateNewCharacter(ctx context.Context, char *character.Character) bool
	DeleteCharacter(ctx context.Context, characterID uint32) bool
}
