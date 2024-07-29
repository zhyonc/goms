package usecase

import (
	"context"
	"goms/mongodb/model/social"
	"goms/util"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type socialUsecase struct {
	baseUsecase
}

func NewSocialUsecase(db *mongo.Database) SocialUsecase {
	coll := db.Collection(socialCollName)
	u := &socialUsecase{
		baseUsecase: NewBaseUsecase("SocialUsecase", coll),
	}
	return u
}

// CreateNewModel implements SocialUsecase.
func (u *socialUsecase) CreateNewModel(ctx context.Context, model any) bool {
	if model == nil {
		slog.Error("Failed to create new character model")
		return false
	}
	return u.baseUsecase.InsertOne(ctx, model)
}

// FindModelByID implements SocialUsecase.
func (u *socialUsecase) FindModelByID(ctx context.Context, id uint32) any {
	filter := bson.D{{Key: "_id", Value: id}}
	soc := &social.Social{}
	if !u.baseUsecase.FindOne(ctx, filter, soc) {
		return nil
	}
	return soc
}

// UpdateModelByID implements SocialUsecase.
func (u *socialUsecase) UpdateModelByID(ctx context.Context, id uint32, model any) bool {
	filter := bson.D{{Key: "_id", Value: id}}
	soc, ok := model.(*social.Social)
	if !ok {
		slog.Error("Failed to assert type *social.Social", "id", id)
		return false
	}
	if model == nil {
		slog.Error("Social is empty", "id", id)
		return false
	}
	soc.UpdateDate = util.DBTime2Local(time.Now())
	return u.UpdateOne(ctx, filter, soc)
}
