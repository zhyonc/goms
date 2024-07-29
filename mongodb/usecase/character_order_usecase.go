package usecase

import (
	"context"
	"goms/mongodb/model"
	"goms/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type characterOrderUsecase struct {
	baseUsecase
}

func NewCharacterOrderUsecase(db *mongo.Database) CharacterOrderUsecase {
	coll := db.Collection(characterOrderCollName)
	u := &characterOrderUsecase{
		baseUsecase: NewBaseUsecase("CharacterOrderUsecase", coll),
	}
	return u
}

// CreateNewCharacterOrder implements CharacterOrderUsecase.
func (u *characterOrderUsecase) CreateNewCharacterOrder(ctx context.Context, charOrder *model.CharacterOrder) bool {
	charOrder.ID = primitive.NewObjectID()
	return u.baseUsecase.InsertOne(ctx, charOrder)
}

// FindCharacterOrder implements CharacterOrderUsecase.
func (u *characterOrderUsecase) FindCharacterOrder(ctx context.Context, accountID uint32, worldID uint8) *model.CharacterOrder {
	filter := bson.D{{Key: "account_id", Value: accountID}, {Key: "world_id", Value: worldID}}
	charOrder := &model.CharacterOrder{}
	if !u.baseUsecase.FindOne(ctx, filter, charOrder) {
		return nil
	}
	return charOrder
}

// UpdateCharacterOrder implements CharacterOrderUsecase.
func (u *characterOrderUsecase) UpdateCharacterOrder(ctx context.Context, charOrder *model.CharacterOrder) bool {
	filter := bson.D{{Key: "_id", Value: charOrder.ID}}
	update := bson.M{"character_ids": charOrder.CharacterIDs, "update_date": util.DBTime2Local(time.Now())}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}
