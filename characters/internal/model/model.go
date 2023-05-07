package model

type Character struct {
	ID          string `bson:"_id,omitempty"`
	Name        *string `bson:"name,omitempty"`
	Description *string `bson:"description,omitempty"`
	Health      *int    `bson:"health,omitempty"`
	Experience  *int    `bson:"experience,omitempty"`
	InventoryID *string `bson:"inventoryId,omitempty"`
}
