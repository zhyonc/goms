package usecase

import (
	"context"
	"goms/mongodb/model/character"
	"goms/util"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type characterUsecase struct {
	baseUsecase
}

func NewCharacterUsecase(db *mongo.Database) CharacterUsecase {
	coll := db.Collection(characterCollName)
	u := &characterUsecase{
		baseUsecase: NewBaseUsecase("CharacterUsecase", coll),
	}
	return u
}

// CreateNewModel implements CharacterUsecase.
func (u *characterUsecase) CreateNewModel(ctx context.Context, model any) bool {
	if model == nil {
		slog.Error("Failed to create new character model")
		return false
	}
	return u.baseUsecase.InsertOne(ctx, model)
}

// FindModelByID implements CharacterUsecase.
func (u *characterUsecase) FindModelByID(ctx context.Context, id uint32) any {
	filter := bson.D{{Key: "_id", Value: id}}
	char := &character.Character{}
	if !u.baseUsecase.FindOne(ctx, filter, char) {
		return nil
	}
	return char
}

// UpdateModelByID implements CharacterUsecase.
func (u *characterUsecase) UpdateModelByID(ctx context.Context, id uint32, model any) bool {
	filter := bson.D{{Key: "_id", Value: id}}
	char, ok := model.(*character.Character)
	if !ok {
		slog.Error("Failed to assert type *character.Character", "id", id)
		return false
	}
	if model == nil {
		slog.Error("Character is empty", "id", id)
		return false
	}
	char.UpdateDate = util.DBTime2Local(time.Now())
	return u.UpdateOne(ctx, filter, char)
}

// FindCharacterName implements CharacterUsecase.
func (u *characterUsecase) FindCharacterName(ctx context.Context, characterName string) bool {
	filter := bson.M{"name": characterName}
	return u.baseUsecase.FindCount(ctx, filter) > 0
}

// FindCharactersByAccountID implements CharacterUsecase.
func (u *characterUsecase) FindCharactersByAccountID(ctx context.Context, accountID uint32, worldID uint8) []*character.Character {
	filter := bson.D{{Key: "account_id", Value: accountID}, {Key: "world_id", Value: worldID}}
	chars := make([]*character.Character, 0)
	ok := u.baseUsecase.FindMany(ctx, filter, &chars)
	if !ok {
		return nil
	}
	return chars
}

// FindCharacterCount implements CharacterUsecase.
func (u *characterUsecase) FindCharacterCount(ctx context.Context, accountID uint32, worldID uint8) uint8 {
	_ = bson.D{{Key: "account_id", Value: accountID}, {Key: "world_id", Value: worldID}}
	return 0
}

// CreateNewCharacter implements CharacterUsecase.
func (u *characterUsecase) CreateNewCharacter(ctx context.Context, char *character.Character) bool {
	char.CreateDate = util.DBTime2Local(time.Now())
	return u.baseUsecase.InsertOne(ctx, char)
}

// ReserveCharacter implements CharacterUsecase.
func (u *characterUsecase) ReserveCharacter(ctx context.Context, characterID uint32, reservedDate time.Time) bool {
	filter := bson.D{{Key: "_id", Value: characterID}}
	update := bson.M{"is_reserved": true, "is_deleted": false, "reserved_date": util.DBTime2Local(time.Now())}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}

// DeleteCharacter implements CharacterUsecase.
func (u *characterUsecase) DeleteCharacter(ctx context.Context, characterID uint32) bool {
	filter := bson.D{{Key: "_id", Value: characterID}}
	update := bson.M{"is_reserved": false, "is_deleted": true, "deleted_date": util.DBTime2Local(time.Now())}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}

// RestoreCharacter implements CharacterUsecase.
func (u *characterUsecase) RestoreCharacter(ctx context.Context, characterID uint32) bool {
	filter := bson.D{{Key: "_id", Value: characterID}}
	update := bson.M{"is_reserved": false, "is_deleted": false}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}

// UpdateLoginDate implements CharacterUsecase.
func (u *characterUsecase) UpdateLoginDate(ctx context.Context, characterID uint32) {
	filter := bson.D{{Key: "_id", Value: characterID}}
	update := bson.M{"login_date": util.DBTime2Local(time.Now())}
	_ = u.baseUsecase.UpdateOne(ctx, filter, update)
}
