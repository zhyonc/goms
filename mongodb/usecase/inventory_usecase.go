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

type inventoryUsecase struct {
	baseUsecase
}

func NewInventoryUsecase(db *mongo.Database) InventoryUsecase {
	coll := db.Collection(inventoryCollName)
	u := &inventoryUsecase{
		baseUsecase: NewBaseUsecase("InventoryUsecase", coll),
	}
	return u
}

// CreateNewModel implements InventoryUsecase.
func (u *inventoryUsecase) CreateNewModel(ctx context.Context, model any) bool {
	if model == nil {
		slog.Error("Failed to create new character model")
		return false
	}
	return u.baseUsecase.InsertOne(ctx, model)
}

// FindModelByID implements InventoryUsecase.
func (u *inventoryUsecase) FindModelByID(ctx context.Context, id uint32) any {
	filter := bson.D{{Key: "_id", Value: id}}
	inv := &inventory.Inventory{}
	if !u.baseUsecase.FindOne(ctx, filter, inv) {
		return nil
	}
	return inv
}

// UpdateModelByID implements InventoryUsecase.
func (u *inventoryUsecase) UpdateModelByID(ctx context.Context, id uint32, model any) bool {
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
