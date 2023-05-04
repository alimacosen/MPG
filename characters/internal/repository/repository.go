package repository

import (
	model "characters/internal/model"
	"context"
	"errors"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CharacterRepository interface {
	FindByID(ctx context.Context, id string) (*model.Character, error)
	Create(ctx context.Context, character *model.Character) (*model.Character, error)
	Update(ctx context.Context, character *model.Character) (*model.Character, error)
	Delete(ctx context.Context, id string) (*model.Character, error)
}

type mongoCharacterRepository struct {
	db     *mongo.Database
	client *mongo.Client
}

func NewCharacterRepository(client *mongo.Client, dbName string) CharacterRepository {
	return &mongoCharacterRepository{
		db:     client.Database(dbName),
		client: client,
	}
}

func (r *mongoCharacterRepository) Create(ctx context.Context, character *model.Character) (*model.Character, error) {
	collection := r.db.Collection("charactersCollection")
	res, err := collection.InsertOne(ctx, character)
	if err != nil {
		return nil, err
	}
	id, ok := res.InsertedID.(string)
	if !ok {
		return nil, errors.New("cannot convert InsertedID to string")
	}
	character.ID = id
	return character, err
}

func (r *mongoCharacterRepository) FindByID(ctx context.Context, id string) (*model.Character, error) {
	return &model.Character{}, nil
}

func (r *mongoCharacterRepository) Update(ctx context.Context, user *model.Character) (*model.Character, error) {
	return &model.Character{}, nil
}

func (r *mongoCharacterRepository) Delete(ctx context.Context, id string) (*model.Character, error) {
	return &model.Character{}, nil
}
