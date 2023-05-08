package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	characterservice "mpg/characters/gen/character_service"
	model "mpg/characters/internal/model"
)

type CharacterRepository interface {
	FindByID(ctx context.Context, id string) (*model.Character, error)
	Create(ctx context.Context, character *model.Character) (*model.Character, error)
	Update(ctx context.Context, id string, updateFields interface{}) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type mongoCharacterRepository struct {
	logger *log.Logger
	db     *mongo.Database
	client *mongo.Client
}

func NewCharacterRepository(logger *log.Logger, client *mongo.Client, dbName string) CharacterRepository {
	return &mongoCharacterRepository{
		logger: logger,
		db:     client.Database(dbName),
		client: client,
	}
}

func (r *mongoCharacterRepository) Create(ctx context.Context, character *model.Character) (*model.Character, error) {
	collection := r.db.Collection("charactersCollection")
	if ok := r.characterNameUniqueCheck(ctx, character.Name, collection); !ok {
		return nil, characterservice.CreateNameNotUnique("Duplicated name found")
	}
	res, err := collection.InsertOne(ctx, character)
	if err != nil {
		return nil, err
	}
	character.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return character, err
}

func (r *mongoCharacterRepository)characterNameUniqueCheck(ctx context.Context, name *string, collection *mongo.Collection) bool {
	if name == nil {
		return true
	}
	filter := bson.M{"name": *name}
	var character model.Character
	_ = collection.FindOne(ctx, filter).Decode(&character)
	if character.Name == nil {
		return true
	}
	return false
}

func (r *mongoCharacterRepository) FindByID(ctx context.Context, id string) (*model.Character, error) {
	collection := r.db.Collection("charactersCollection")
	var character model.Character
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&character)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return nil, characterservice.GetNoMatch("Character not found")
		}
		return nil, err
	}
	character.ID = id
	return &character, nil
}

func (r *mongoCharacterRepository) Update(ctx context.Context, id string, updateFields interface{}) (int, error) {
	collection := r.db.Collection("charactersCollection")

	if ok := r.characterNameUniqueCheck(ctx, updateFields.(model.Character).Name, collection); !ok {
		return 0, characterservice.CreateNameNotUnique("Duplicated name found")
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}

	updateDoc := bson.M{"$set": updateFields}
	res, err := collection.UpdateByID(ctx, objectID, updateDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, characterservice.UpdateNoMatch("Character not found")
		}
		return 0, err
	}

	modifiedCnt := int(res.ModifiedCount)

	return modifiedCnt, nil
}

func (r *mongoCharacterRepository) Delete(ctx context.Context, id string) (int, error) {
	collection := r.db.Collection("charactersCollection")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}
	filter := bson.M{"_id": objectID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, characterservice.DeleteNoMatch("Character not found")
		}
		return 0, err
	}

	return int(res.DeletedCount), nil
}
