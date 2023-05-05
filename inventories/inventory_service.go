package inventoriesapi

import (
	"context"
	"log"
	inventoryservice "mpg/inventories/gen/inventory_service"
	"mpg/inventories/internal/handler"
)

// InventoryService service example implementation.
// The example methods log the requests and return zero values.
type inventoryServicesrvc struct {
	logger *log.Logger
	InventoryHandler *handler.InventoryHandler
}

// NewInventoryService returns the InventoryService service implementation.
func NewInventoryService(logger *log.Logger) inventoryservice.Service {
	return &inventoryServicesrvc{logger, handler.NewInventoryHandler(logger)}
}

// CreateInventory implements createInventory.
func (s *inventoryServicesrvc) CreateInventory(ctx context.Context, p *inventoryservice.CreateInventoryPayload) (res *inventoryservice.Inventory, err error) {
	s.logger.Println("payload: ", p)
	res, err = s.InventoryHandler.CreateInventory(ctx, p)
	return
}

// GetInventory implements getInventory.
func (s *inventoryServicesrvc) GetInventory(ctx context.Context, p *inventoryservice.GetInventoryPayload) (res *inventoryservice.Inventory, err error) {
	s.logger.Println("payload: ", p)
	res, err = s.InventoryHandler.GetInventory(ctx, p)
	return
}

// UpdateInventory implements updateInventory.
func (s *inventoryServicesrvc) UpdateInventory(ctx context.Context, p *inventoryservice.UpdateInventoryPayload) (res int, err error) {
	s.logger.Println("payload: ", p)
	res, err = s.InventoryHandler.UpdateInventory(ctx, p)
	return
}

// DeleteInventory implements deleteInventory.
func (s *inventoryServicesrvc) DeleteInventory(ctx context.Context, p *inventoryservice.DeleteInventoryPayload) (res int, err error) {
	s.logger.Println("payload: ", p)
	res, err = s.InventoryHandler.DeleteInventory(ctx, p)
	return
}
