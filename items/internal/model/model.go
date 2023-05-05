package model

import (
	itemservice "mpg/items/gen/item_service"
)

type UpdateFields struct {
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Damage      int    `bson:"damage,omitempty"`
	Healing     int    `bson:"healing,omitempty"`
	Protection  int    `bson:"protection,omitempty"`
}

type Item itemservice.Item
