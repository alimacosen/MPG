package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	itemservice "mpg/items/gen/item_service"
	"mpg/items/internal/model"
)

type ItemRepository interface {
	FindByID(ctx context.Context, idList []string) ([]*model.Item, error)
	FindAll(ctx context.Context) ([]*model.Item, error)
	Create(ctx context.Context, character *model.Item) (*model.Item, error)
	Update(ctx context.Context, id string, updateFields *model.Item) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type mongoItemRepository struct {
	logger *log.Logger
	db     *mongo.Database
	client *mongo.Client
}

func NewItemRepository(logger *log.Logger, client *mongo.Client, dbName string) ItemRepository {
	return &mongoItemRepository{
		logger: logger,
		db:     client.Database(dbName),
		client: client,
	}
}

func (r *mongoItemRepository) Create(ctx context.Context, item *model.Item) (*model.Item, error) {
	collection := r.db.Collection("itemCollection")

	if ok := r.itemNameUniqueCheck(ctx, collection, item.Name); !ok {
		return nil, itemservice.CreateDuplicatedName("item name" + *item.Name + "already exists")
	}

	res, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	item.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return item, err
}

func (r *mongoItemRepository) itemNameUniqueCheck(ctx context.Context, collection *mongo.Collection, name *string) bool {
	if name == nil {
		return true
	}
	filter := bson.M{"name": name}
	var item model.Item
	_ = collection.FindOne(ctx, filter).Decode(&item)
	if item.Name == nil {
		return true
	}
	return false
}

func (r *mongoItemRepository) FindByID(ctx context.Context, idList []string) ([]*model.Item, error) {
	collection := r.db.Collection("itemCollection")

	objectIDs := make([]primitive.ObjectID, len(idList))
	for i, id := range idList {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			r.logger.Println(err)
		}
		objectIDs[i] = objectID
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Item
	if err = cursor.All(ctx, &items); err != nil {
		r.logger.Println(err)
	}

	return items, nil
}

func (r *mongoItemRepository) FindAll(ctx context.Context) ([]*model.Item, error) {
	collection := r.db.Collection("itemCollection")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*model.Item
	if err = cursor.All(ctx, &items); err != nil {
		r.logger.Println(err)
	}

	return items, nil
}

func (r *mongoItemRepository) Update(ctx context.Context, id string, updateFields *model.Item) (int, error) {
	collection := r.db.Collection("itemCollection")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}

	ok := r.itemNameUniqueCheck(ctx, collection, updateFields.Name)
	if !ok {
		return 0, itemservice.CreateDuplicatedName("item name" + *updateFields.Name + "already exists")
	}

	updateDoc := bson.M{"$set": *updateFields}
	res, err := collection.UpdateByID(ctx, objectID, updateDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, itemservice.UpdateNoMatch("Item not found")
		}
		return 0, err
	}

	modifiedCnt := int(res.ModifiedCount)

	return modifiedCnt, nil
}

func (r *mongoItemRepository) Delete(ctx context.Context, id string) (int, error) {
	collection := r.db.Collection("itemCollection")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Println(err)
	}
	filter := bson.M{"_id": objectID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Println("No document found with the specified filter.")
			return 0, itemservice.DeleteNoMatch("Item not found")
		}
		return 0, err
	}

	return int(res.DeletedCount), nil
}
