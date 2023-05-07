package itemsapi

import (
	"context"
	"log"
	itemservice "mpg/items/gen/item_service"
	"mpg/items/internal/handler"
)

// ItemService service example implementation.
// The example methods log the requests and return zero values.
type itemServicesrvc struct {
	logger *log.Logger
	itemHandler *handler.ItemHandler
}

// NewItemService returns the ItemService service implementation.
func NewItemService(logger *log.Logger) itemservice.Service {
	return &itemServicesrvc{logger, handler.NewInventoryHandler(logger)}
}

// CreateItem implements createItem.
func (s *itemServicesrvc) CreateItem(ctx context.Context, p *itemservice.CreateItemPayload) (res *itemservice.Item, err error) {
	res, err = s.itemHandler.CreateItem(ctx, p)
	return
}

// GetItems implements getItem.
func (s *itemServicesrvc) GetItems(ctx context.Context, p *itemservice.GetItemsPayload) (res []*itemservice.Item, err error) {
	res, err = s.itemHandler.GetItems(ctx, p)
	return
}

func (s *itemServicesrvc) GetAllItems(ctx context.Context) (res []*itemservice.Item, err error) {
	res, err = s.itemHandler.GetAllItems(ctx)
	return
}

// UpdateItem implements updateItem.
func (s *itemServicesrvc) UpdateItem(ctx context.Context, p *itemservice.UpdateItemPayload) (res int, err error) {
	res, err = s.itemHandler.UpdateItem(ctx, p)
	return
}

// DeleteItem implements deleteItem.
func (s *itemServicesrvc) DeleteItem(ctx context.Context, p *itemservice.DeleteItemPayload) (res int, err error) {
	res, err = s.itemHandler.DeleteItem(ctx, p)
	return
}
