package entryapi

import (
	"context"
	"log"
	entryitemservice "mpg/entry/gen/entry_item_service"
	"mpg/entry/internal/handler"
	itemservice "mpg/items/gen/item_service"
)

// EntryItemService service example implementation.
// The example methods log the requests and return zero values.
type entryItemServicesrvc struct {
	logger *log.Logger
	entryHandler *handler.EntryHandler
}

// NewEntryItemService returns the EntryItemService service implementation.
func NewEntryItemService(logger *log.Logger) entryitemservice.Service {
	return &entryItemServicesrvc{logger, handler.GetEntryHandler(logger)}
}

// CreateItem implements creatItem.
func (s *entryItemServicesrvc) CreateItem(ctx context.Context, p *entryitemservice.CreateItemPayload) (res *entryitemservice.Item, err error) {
	createItem := s.entryHandler.GetInstance().GetItemClient().CreateItem()
	payload := &itemservice.CreateItemPayload{
		Name:        p.Name,
		Description: p.Description,
		Damage: p.Damage,
		Healing: p.Healing,
		Protection: p.Protection,
	}
	result, err := createItem(ctx, payload)
	if err != nil || result == nil {
		return nil, err
	}
	item := result.(*itemservice.Item)
	return &entryitemservice.Item{
		ID: item.ID,
		Name: item.Name,
		Description: item.Description,
		Damage: item.Damage,
		Healing: item.Healing,
		Protection: item.Protection,
	}, nil
}

// GetItem implements getItem.
func (s *entryItemServicesrvc) GetItem(ctx context.Context, p *entryitemservice.GetItemPayload) (res *entryitemservice.Item, err error) {
	getItem := s.entryHandler.GetInstance().GetItemClient().GetItem()
	payload := &itemservice.GetItemPayload{
		ID:        p.ID,
	}
	result, err := getItem(ctx, payload)
	if err != nil || result == nil {
		return nil, err
	}
	item := result.(*itemservice.Item)
	return &entryitemservice.Item{
		ID: item.ID,
		Name: item.Name,
		Description: item.Description,
		Damage: item.Damage,
		Healing: item.Healing,
		Protection: item.Protection,
	}, nil
}

// UpdateItem implements updateItem.
func (s *entryItemServicesrvc) UpdateItem(ctx context.Context, p *entryitemservice.UpdateItemPayload) (res int, err error) {
	updateItem := s.entryHandler.GetInstance().GetItemClient().UpdateItem()
	payload := &itemservice.UpdateItemPayload{
		ID: p.ID,
		Description: p.Description,
		Damage: p.Damage,
		Healing: p.Healing,
		Protection: p.Protection,
	}
	result, err := updateItem(ctx, payload)
	if err != nil || result == nil {
		return 0, err
	}
	cnt := result.(int)
	return cnt, err
}

// DeleteItem implements deleteItem.
func (s *entryItemServicesrvc) DeleteItem(ctx context.Context, p *entryitemservice.DeleteItemPayload) (res int, err error) {
	deleteItem := s.entryHandler.GetInstance().GetItemClient().DeleteItem()
	payload := &itemservice.DeleteItemPayload{
		ID:        p.ID,
	}
	result, err := deleteItem(ctx, payload)
	if err != nil || result == nil {
		return 0, err
	}
	cnt := result.(int)
	return cnt, err
}
