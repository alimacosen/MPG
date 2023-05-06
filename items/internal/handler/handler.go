package handler

import (
	"context"
	"log"
	itemservice "mpg/items/gen/item_service"
	"mpg/items/internal/injections"
	"mpg/items/internal/model"
)

type ItemHandler struct {
	logger *log.Logger
	instances *injections.Instances
}

func NewInventoryHandler(logger *log.Logger) *ItemHandler {
	return &ItemHandler{logger: logger, instances: injections.NewInstances(logger)}
}

func (i *ItemHandler) CreateItem(ctx context.Context, p *itemservice.CreateItemPayload) (res *itemservice.Item, err error) {
	itemPreliminary, err := i.payloadCheck(p)
	if err != nil {
		return nil, err
	}

	svc := i.instances.GetSvc()

	itemRes, err := svc.Create(ctx, itemPreliminary)
	if err != nil {
		return nil, err
	}

	res = convertOne(itemRes)
	return
}

func (i *ItemHandler) payloadCheck(p *itemservice.CreateItemPayload) (*model.Item, error) {
	if len(p.Name) == 0 || len(p.Description) == 0 || p.Damage == 0 || p.Healing == 0 || p.Protection == 0 {
		return nil, itemservice.CreateInvalidArgs("all payload arguments can not be an zero value")
	}

	item := &model.Item{}
	item.Name = p.Name
	item.Description = p.Description
	item.Damage = p.Damage
	item.Healing = p.Healing
	item.Protection = p.Protection

	return item, nil
}

func (i *ItemHandler) GetItem (ctx context.Context, p *itemservice.GetItemPayload) (res *itemservice.Item, err error) {
	id := p.ID
	if len(id) == 0 {
		return nil, itemservice.GetInvalidArgs("Id can not be an empty string")
	}

	svc := i.instances.GetSvc()

	item, err := svc.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res = convertOne(item)
	return
}

func (i *ItemHandler) GetAllItems (ctx context.Context) (res []*itemservice.Item, err error) {
	svc := i.instances.GetSvc()
	items, err := svc.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = convertAll(items)
	return
}

func (i *ItemHandler) UpdateItem(ctx context.Context, p *itemservice.UpdateItemPayload) (res int, err error) {
	id := p.ID
	updateFields := assembleUpdateFields(p)

	svc := i.instances.GetSvc()

	modifiedCnt, err := svc.Update(ctx, id, updateFields)
	if err != nil {
		return 0, err
	}

	res = modifiedCnt
	return
}

func (i *ItemHandler) DeleteItem (ctx context.Context, p *itemservice.DeleteItemPayload) (res int, err error) {
	id := p.ID
	if len(id) == 0 {
		return 0, itemservice.DeleteInvalidArgs("Id can not be an empty string")
	}

	svc := i.instances.GetSvc()

	deletedCnt, err := svc.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return deletedCnt, nil
}

func convertAll(items []*model.Item) []*itemservice.Item {
	convertedItems := make([]*itemservice.Item, 0)
	for i := 0; i < len(items); i++ {
		convertedItems = append(convertedItems, convertOne(items[i]))
	}
	return convertedItems
}

func convertOne(i *model.Item) *itemservice.Item {
	return &itemservice.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Damage:      i.Damage,
		Healing:     i.Healing,
		Protection:  i.Protection,
	}
}

func assembleUpdateFields (p *itemservice.UpdateItemPayload) *model.Item {
	updateFields := &model.Item{}

	if p.Description != nil && len(*p.Description) != 0{
		updateFields.Description = *p.Description
	}
	if p.Damage != nil && *p.Damage != 0{
		updateFields.Damage = *p.Damage
	}
	if p.Healing != nil && *p.Healing != 0{
		updateFields.Healing = *p.Healing
	}
	if p.Protection != nil && *p.Protection != 0{
		updateFields.Protection = *p.Protection
	}

	return updateFields
}