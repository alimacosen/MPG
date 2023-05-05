package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	inventoryservice "mpg/inventories/gen/inventory_service"
	model "mpg/inventories/internal/model"
)

type InventoryRepository interface {
	FindByID(ctx context.Context, id string) (*model.Inventory, error)
	Create(ctx context.Context, character *model.Inventory) (*model.Inventory, error)
	Update(ctx context.Context, id string, updateFields model.UpdateFields) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type mongoInventoryRepository struct {
	logger *log.Logger
	db     *mongo.Database
	client *mongo.Client
}

func NewInventoryRepository(logger *log.Logger, client *mongo.Client, dbName string) InventoryRepository {
	return &mongoInventoryRepository{
		logger: logger,
		db:     client.Database(dbName),
		client: client,
	}
}

func (r *mongoInventoryRepository) Create(ctx context.Context, inventory *model.Inventory) (*model.Inventory, error) {
	collection := r.db.Collection("inventoryCollection")
	res, err := collection.InsertOne(ctx, inventory)
	if err != nil {
		return nil, err
	}
	inventory.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return inventory, err
}

func (r *mongoInventoryRepository) FindByID(ctx context.Context, id string) (*model.Inventory, error) {
	collection := r.db.Collection("inventoryCollection")
	var inventory model.Inventory
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&inventory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return nil, inventoryservice.GetNoMatch("Inventory not found")
		}
		return nil, err
	}

	return &inventory, nil
}

func (r *mongoInventoryRepository) Update(ctx context.Context, id string, updateFields model.UpdateFields) (int, error) {
	collection := r.db.Collection("inventoryCollection")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}

	updateDoc := bson.M{"$set": updateFields}
	res, err := collection.UpdateByID(ctx, objectID, updateDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, inventoryservice.UpdateNoMatch("Inventory not found")
		}
		return 0, err
	}

	modifiedCnt := int(res.ModifiedCount)

	return modifiedCnt, nil
}

func (r *mongoInventoryRepository) Delete(ctx context.Context, id string) (int, error) {
	collection := r.db.Collection("inventoryCollection")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}
	filter := bson.M{"_id": objectID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, inventoryservice.DeleteNoMatch("Inventory not found")
		}
		return 0, err
	}

	return int(res.DeletedCount), nil
}
