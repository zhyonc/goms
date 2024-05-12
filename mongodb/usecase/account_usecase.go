package usecase

import (
	"context"
	"goms/mongodb/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountUsecase struct {
	baseUsecase
}

func NewAccountUsecase(db *mongo.Database) AccountUsecase {
	coll := db.Collection(accountColl)
	u := &accountUsecase{
		baseUsecase: NewBaseUsecase("AccountUsecase", coll),
	}
	return u
}

// FindAccountByUsername implements AccountUsecase.
func (u *accountUsecase) FindAccountByUsername(ctx context.Context, username string) *model.Account {
	filter := bson.D{{Key: "username", Value: username}}
	account := &model.Account{}
	ok := u.baseUsecase.FindOne(ctx, filter, account)
	if !ok {
		return nil
	}
	return account
}

// CreateNewAccount implements AccountUsecase.
func (u *accountUsecase) CreateNewAccount(ctx context.Context, account *model.Account) bool {
	return u.baseUsecase.InsertOne(ctx, account)
}

// UpdateGenderAndSecondPassword implements AccountUsecase.
func (u *accountUsecase) UpdateGenderAndSecondPassword(ctx context.Context, username, secondPassWord string, gender bool) bool {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.M{
		"second_password": secondPassWord,
		"gender":          gender,
		"update_time":     time.Now(),
	}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}

// UpdateLoginTime implements AccountUsecase.
func (u *accountUsecase) UpdateLoginTime(ctx context.Context, accountID uint32, mac string) {
	filter := bson.D{{Key: "_id", Value: accountID}}
	update := bson.M{
		"login_time": time.Now(),
		"login_mac":  mac,
	}
	_ = u.baseUsecase.UpdateOne(ctx, filter, update)
}

// FindAccountByID implements AccountUsecase.
func (u *accountUsecase) FindAccountByID(ctx context.Context, accountID uint32) *model.Account {
	filter := bson.D{{Key: "_id", Value: accountID}}
	account := &model.Account{}
	ok := u.baseUsecase.FindOne(ctx, filter, account)
	if !ok {
		return nil
	}
	return account
}
