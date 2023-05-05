package entryapi

import (
	"context"
	"log"
	entryinventoryservice "mpg/entry/gen/entry_inventory_service"
)

// EntryInventoryService service example implementation.
// The example methods log the requests and return zero values.
type entryInventoryServicesrvc struct {
	logger *log.Logger
}

// NewEntryInventoryService returns the EntryInventoryService service
// implementation.
func NewEntryInventoryService(logger *log.Logger) entryinventoryservice.Service {
	return &entryInventoryServicesrvc{logger}
}

// GetInventory implements getInventory.
func (s *entryInventoryServicesrvc) GetInventory(ctx context.Context, p *entryinventoryservice.GetInventoryPayload) (res *entryinventoryservice.Inventory, err error) {
	res = &entryinventoryservice.Inventory{}
	s.logger.Print("entryInventoryService.getInventory")
	return
}

// UpdateInventory implements updateInventory.
func (s *entryInventoryServicesrvc) UpdateInventory(ctx context.Context, p *entryinventoryservice.UpdateInventoryPayload) (res int, err error) {
	s.logger.Print("entryInventoryService.updateInventory")
	return
}
