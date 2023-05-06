package handler

import (
	"context"
	"log"
	inventoryservice "mpg/inventories/gen/inventory_service"
	"mpg/inventories/internal/injections"
	"mpg/inventories/internal/model"
)

type InventoryHandler struct {
	logger *log.Logger
	instances *injections.Instances
}

func NewInventoryHandler(logger *log.Logger) *InventoryHandler {
	return &InventoryHandler{logger: logger, instances: injections.NewInstances(logger)}
}

func (i *InventoryHandler) CreateInventory(ctx context.Context, p *inventoryservice.CreateInventoryPayload) (res *inventoryservice.Inventory, err error) {
	userId := p.UserID
	if len(userId) == 0 {
		return nil, inventoryservice.CreateInvalidArgs("userId can not be an empty string")
	}

	svc := i.instances.GetSvc()

	inventoryPreliminary := &model.Inventory{
		UserID: userId,
	}

	ipr, err := svc.Create(ctx, inventoryPreliminary)
	if err != nil {
		return nil, err
	}

	res = convert(ipr)
	i.logger.Println("res converted: ", res)
	return
}

func (i *InventoryHandler) GetInventory (ctx context.Context, p *inventoryservice.GetInventoryPayload) (res *inventoryservice.Inventory, err error) {
	id := p.ID
	if len(id) == 0 {
		return nil, inventoryservice.GetInvalidArgs("Id can not be an empty string")
	}

	svc := i.instances.GetSvc()

	cp, err := svc.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res = convert(cp)
	i.logger.Println("res converted: ", res)
	return
}

func (i *InventoryHandler) UpdateInventory(ctx context.Context, p *inventoryservice.UpdateInventoryPayload) (res int, err error) {
	id := p.ID
	updateFields := assembleUpdateFields(p)

	if len(updateFields.ItemsID) == 0 {
		return 0, inventoryservice.UpdateInvalidArgs("ItemsId can not be an empty array")
	}

	svc := i.instances.GetSvc()

	modifiedCnt, err := svc.Update(ctx, id, updateFields)
	if err != nil {
		return 0, err
	}

	res = modifiedCnt
	return
}

func (i *InventoryHandler) DeleteInventory (ctx context.Context, p *inventoryservice.DeleteInventoryPayload) (res int, err error) {
	id := p.ID
	if len(id) == 0 {
		return 0, inventoryservice.DeleteInvalidArgs("Id can not be an empty string")
	}

	svc := i.instances.GetSvc()

	deletedCnt, err := svc.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return deletedCnt, nil
}

func convert(i *model.Inventory) *inventoryservice.Inventory {
	return &inventoryservice.Inventory{
		ID: i.ID,
		UserID: i.UserID,
		ItemsID: i.ItemsID,
	}
}

func assembleUpdateFields (p *inventoryservice.UpdateInventoryPayload) *model.Inventory {
	updateFields := &model.Inventory{}

	if p.ItemsID != nil {
		updateFields.ItemsID = p.ItemsID
	}

	return updateFields
}