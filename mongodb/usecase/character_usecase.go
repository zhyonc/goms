package usecase

import (
	"context"
	"goms/mongodb/model/character"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type characterUsecase struct {
	baseUsecase
}

func NewCharacterUsecase(db *mongo.Database) CharacterUsecase {
	coll := db.Collection(characterColl)
	u := &characterUsecase{
		baseUsecase: NewBaseUsecase("CharacterUsecase", coll),
	}
	return u
}

// FindCharacterName implements CharacterUsecase.
func (u *characterUsecase) FindCharacterName(ctx context.Context, characterName string) bool {
	filter := bson.M{"name": characterName}
	return u.baseUsecase.FindCount(ctx, filter) > 0
}

// GetCharactersByAccountID implements CharacterUsecase.
func (u *characterUsecase) GetCharactersByAccountID(ctx context.Context, accountID uint32, worldID uint8) []*character.Character {
	filter := bson.D{{Key: "account_id", Value: accountID}, {Key: "world_id", Value: worldID}}
	chars := make([]*character.Character, 0)
	u.baseUsecase.FindMany(ctx, filter, &chars)
	return chars
}

// GetCharacterByID implements CharacterUsecase.
func (u *characterUsecase) GetCharacterByID(ctx context.Context, characterID uint32) *character.Character {
	filter := bson.D{{Key: "_id", Value: characterID}}
	char := &character.Character{}
	if !u.baseUsecase.FindOne(ctx, filter, char) {
		return nil
	}
	return char
}

// CreateNewCharacter implements CharacterUsecase.
func (u *characterUsecase) CreateNewCharacter(ctx context.Context, char *character.Character) bool {
	return u.baseUsecase.InsertOne(ctx, char)
}

// DeleteCharacter implements CharacterUsecase.
func (u *characterUsecase) DeleteCharacter(ctx context.Context, characterID uint32) bool {
	filter := bson.D{{Key: "_id", Value: characterID}}
	update := bson.M{"is_deleted": true}
	return u.baseUsecase.UpdateOne(ctx, filter, update)
}
