package model

import (
	itemservice "mpg/items/gen/item_service"
)

type UpdateFields struct {
	Name        string `bson:"name,omitempty"`
	Description string `bson:"name,omitempty"`
	Damage      int    `bson:"description,omitempty"`
	Healing     int    `bson:"healing,omitempty"`
	Protection  int    `bson:"protection,omitempty"`
}

type Item itemservice.Item
