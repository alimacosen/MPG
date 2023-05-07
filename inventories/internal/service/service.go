package service

import (
	"context"
	model "mpg/inventories/internal/model"
	repo "mpg/inventories/internal/repository"
)

type InventorySvcInterface interface {
	Create(ctx context.Context, inventoryPreliminary *model.Inventory) (*model.Inventory, error)
	GetById(ctx context.Context, id string) (*model.Inventory, error)
	Update(ctx context.Context, id string, updateFields *model.Inventory) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type inventoryService struct {
	repo repo.InventoryRepository
}


func NewInventoryService(repo repo.InventoryRepository) InventorySvcInterface {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) Create(ctx context.Context, inventoryPreliminary *model.Inventory) (*model.Inventory, error) {
	result, err := s.repo.Create(ctx, inventoryPreliminary)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *inventoryService) GetById(ctx context.Context, id string) (*model.Inventory, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *inventoryService) Update(ctx context.Context, id string, updateFields *model.Inventory) (int, error) {
	result, err := s.repo.Update(ctx, id, *updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *inventoryService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}