package usecase

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type baseUsecase struct {
	name string
	coll *mongo.Collection
}

func NewBaseUsecase(name string, coll *mongo.Collection) baseUsecase {
	u := baseUsecase{name: name, coll: coll}
	return u
}

func (u *baseUsecase) InsertOne(ctx context.Context, model any) bool {
	res, err := u.coll.InsertOne(ctx, model)
	if err != nil || res == nil {
		slog.Error("Failed to insert new doc", "err", err, "model", model, "name", u.name)
		return false
	}
	return true
}

func (u *baseUsecase) FindCount(ctx context.Context, filter bson.M) int64 {
	count, err := u.coll.CountDocuments(ctx, filter)
	if err != nil {
		slog.Error("Failed to find count", "err", err, "filter", filter, "name", u.name)
		return 0
	}
	return count
}

func (u *baseUsecase) FindOne(ctx context.Context, filter bson.D, model any) bool {
	err := u.coll.FindOne(ctx, filter).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.Error("Find doc is empty", "err", err, "filter", filter, "name", u.name)
		} else {
			slog.Error("Failed to find doc", "err", err, "filter", filter, "name", u.name)
		}
		return false
	}
	return true
}

func (u *baseUsecase) FindMany(ctx context.Context, filter bson.D, models any) bool {
	cur, err := u.coll.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.Error("Find docs are empty", "err", err, "filter", filter, "name", u.name)
		} else {
			slog.Error("Failed to find many docs ", "err", err, "filter", filter, "name", u.name)
		}
		return false
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, models)
	if err != nil {
		slog.Error("Failed to decode when find many docs", "err", err, "name", u.name)
		return false
	}
	return true
}

func (u *baseUsecase) FindOneAndUpdate(ctx context.Context, filter, update bson.M, opts *options.FindOneAndUpdateOptions, model any) {
	res := u.coll.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		slog.Error("Find doc is empty", "err", res.Err(), "filter", filter, "update", update, "name", u.name)
		return
	}
	err := res.Decode(model)
	if err != nil {
		slog.Error("Failed to decode when find one and update doc", "err", err, "filter", filter, "update", update, "name", u.name)
		return
	}
	slog.Debug("Find one and update doc ok")
}

func (u *baseUsecase) UpdateOne(ctx context.Context, filter bson.D, update any) bool {
	res, err := u.coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		slog.Error("Failed to update one doc", "err", err, "filter", filter, "name", u.name)
		return false
	}
	slog.Debug("Update one doc result", "filter", filter, "matchedCount", res.MatchedCount, "modifiedCount", res.ModifiedCount)
	return true
}

func (u *baseUsecase) DeleteOne(ctx context.Context, filter bson.M) bool {
	res, err := u.coll.DeleteOne(ctx, filter)
	if err != nil {
		slog.Error("Failed to delete one doc", "err", err, "filter", filter, "name", u.name)
		return false
	}
	slog.Debug("Delete one doc ok", "deletedCount", res.DeletedCount, "filter", filter, "name", u.name)
	return true
}
