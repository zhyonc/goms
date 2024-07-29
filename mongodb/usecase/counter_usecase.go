package usecase

import (
	"context"
	"goms/mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type counterUsecase struct {
	baseUsecase
}

func NewCounterUsecase(db *mongo.Database) CounterUsecase {
	coll := db.Collection(counterCollName)
	u := &counterUsecase{
		baseUsecase: NewBaseUsecase("counterUsecase", coll),
	}
	return u
}
func (c *counterUsecase) getIncID(ctx context.Context, keyName string) *model.Counter {
	filter := bson.M{keyName: bson.M{"$exists": true}}
	update := bson.M{"$inc": bson.M{keyName: 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	opts.SetReturnDocument(options.After) // avoid first find return err before upsert
	model := &model.Counter{}
	c.baseUsecase.FindOneAndUpdate(ctx, filter, update, opts, model)
	return model
}

// GetAccountID implements CounterUsecase.
func (c *counterUsecase) GetAccountID(ctx context.Context) uint32 {
	return c.getIncID(ctx, "account_id").AccountID
}

// GetCharacterID implements CounterUsecase.
func (c *counterUsecase) GetCharacterID(ctx context.Context) uint32 {
	return c.getIncID(ctx, "character_id").CharacterID
}
