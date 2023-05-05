package entryapi

import (
	"context"
	"log"
	entryinventoryservice "mpg/entry/gen/entry_inventory_service"
	"mpg/entry/internal/handler"
	inventoryservice "mpg/inventories/gen/inventory_service"
)

// EntryInventoryService service example implementation.
// The example methods log the requests and return zero values.
type entryInventoryServicesrvc struct {
	logger       *log.Logger
	entryHandler *handler.EntryHandler
}

// NewEntryInventoryService returns the EntryInventoryService service
// implementation.
func NewEntryInventoryService(logger *log.Logger) entryinventoryservice.Service {
	return &entryInventoryServicesrvc{logger, handler.GetEntryHandler(logger)}
}

// GetInventory implements getInventory.
func (s *entryInventoryServicesrvc) GetInventory(ctx context.Context, p *entryinventoryservice.GetInventoryPayload) (res *entryinventoryservice.Inventory, err error) {
	getInventory := s.entryHandler.GetInstance().GetInventoryClient().GetInventory()
	payload := &inventoryservice.GetInventoryPayload{
		ID: p.ID,
	}
	result, err := getInventory(ctx, payload)
	if err != nil || result == nil {
		return nil, err
	}
	inventory := result.(*inventoryservice.Inventory)
	return &entryinventoryservice.Inventory{
		ID:      inventory.ID,
		UserID:  inventory.ID,
		ItemsID: inventory.ItemsID,
	}, nil
}

// UpdateInventory implements updateInventory.
func (s *entryInventoryServicesrvc) UpdateInventory(ctx context.Context, p *entryinventoryservice.UpdateInventoryPayload) (res int, err error) {
	updateInventory := s.entryHandler.GetInstance().GetInventoryClient().UpdateInventory()
	payload := &inventoryservice.UpdateInventoryPayload{
		ID: p.ID,
		ItemsID: p.ItemsID,
	}
	result, err := updateInventory(ctx, payload)
	if err != nil || result == nil {
		return 0, err
	}
	cnt := result.(int)
	return cnt, err
}
