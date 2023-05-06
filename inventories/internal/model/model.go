package model

type Inventory struct {
	ID      string   `bson:"_id,omitempty"`
	UserID  string   `bson:"userId,omitempty"`
	ItemsID []string `bson:"itemsId,omitempty"`
}
