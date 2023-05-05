package model

import (
	inventoryservice "inventories/gen/inventory_service"
)

type UpdateFields struct {
	ItemsId   []string `bson:"itemsId,omitempty"`
}

type Inventory inventoryservice.Inventory
