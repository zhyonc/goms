package usecase

import (
	"context"
	"goms/mongodb/model/inventory"
	"goms/util"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type eventUsecase struct {
	baseUsecase
}

func NewEventUsecase(db *mongo.Database) EventUsecase {
	coll := db.Collection(inventoryCollName)
	u := &eventUsecase{
		baseUsecase: NewBaseUsecase("EventUsecase", coll),
	}
	return u
}

// CreateNewModel implements EventUsecase.
func (u *eventUsecase) CreateNewModel(ctx context.Context, model any) bool {
	panic("unimplemented")
}

// FindModelByID implements EventUsecase.
func (u *eventUsecase) FindModelByID(ctx context.Context, id uint32) any {
	filter := bson.D{{Key: "_id", Value: id}}
	inv := &inventory.Inventory{}
	if !u.baseUsecase.FindOne(ctx, filter, inv) {
		return nil
	}
	return inv
}

// UpdateModelByID implements EventUsecase.
func (u *eventUsecase) UpdateModelByID(ctx context.Context, id uint32, model any) bool {
	filter := bson.D{{Key: "_id", Value: id}}
	inv, ok := model.(*inventory.Inventory)
	if !ok {
		slog.Error("Failed to assert type *inventory.Inventory", "id", id)
		return false
	}
	if model == nil {
		slog.Error("Inventory is empty", "id", id)
		return false
	}
	inv.UpdateDate = util.DBTime2Local(time.Now())
	return u.UpdateOne(ctx, filter, inv)
}
