package service

import (
	"context"
	"mpg/items/internal/model"
	repo "mpg/items/internal/repository"
)

type InventoryService struct {
	repo repo.ItemRepository
}


func NewItemService(repo repo.ItemRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) Create(ctx context.Context, item *model.Item) (*model.Item, error) {
	result, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *InventoryService) GetById(ctx context.Context, id string) (*model.Item, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *InventoryService) GetAll(ctx context.Context) ([]*model.Item, error) {
	result, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *InventoryService) Update(ctx context.Context, id string, updateFields *model.Item) (int, error) {
	result, err := s.repo.Update(ctx, id, *updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *InventoryService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}