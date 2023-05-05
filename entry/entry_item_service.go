package entryapi

import (
	"context"
	"log"
	entryitemservice "mpg/entry/gen/entry_item_service"
)

// EntryItemService service example implementation.
// The example methods log the requests and return zero values.
type entryItemServicesrvc struct {
	logger *log.Logger
}

// NewEntryItemService returns the EntryItemService service implementation.
func NewEntryItemService(logger *log.Logger) entryitemservice.Service {
	return &entryItemServicesrvc{logger}
}

// CreatItem implements creatItem.
func (s *entryItemServicesrvc) CreatItem(ctx context.Context, p *entryitemservice.CreatItemPayload) (res *entryitemservice.Item, err error) {
	res = &entryitemservice.Item{}
	s.logger.Print("entryItemService.creatItem")
	return
}

// GetItem implements getItem.
func (s *entryItemServicesrvc) GetItem(ctx context.Context, p *entryitemservice.GetItemPayload) (res *entryitemservice.Item, err error) {
	res = &entryitemservice.Item{}
	s.logger.Print("entryItemService.getItem")
	return
}

// UpdateItem implements updateItem.
func (s *entryItemServicesrvc) UpdateItem(ctx context.Context, p *entryitemservice.UpdateItemPayload) (res int, err error) {
	s.logger.Print("entryItemService.updateItem")
	return
}

// DeleteItem implements deleteItem.
func (s *entryItemServicesrvc) DeleteItem(ctx context.Context, p *entryitemservice.DeleteItemPayload) (res int, err error) {
	s.logger.Print("entryItemService.deleteItem")
	return
}
