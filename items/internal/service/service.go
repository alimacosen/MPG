package service

import (
	"context"
	"mpg/items/internal/model"
	repo "mpg/items/internal/repository"
)

type ItemSvcInterface interface {
	GetById(ctx context.Context, id string) (*model.Item, error)
	GetAll(ctx context.Context) ([]*model.Item, error)
	Create(ctx context.Context, character *model.Item) (*model.Item, error)
	Update(ctx context.Context, id string, updateFields *model.Item) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type itemService struct {
	repo repo.ItemRepository
}

func NewItemService(repo repo.ItemRepository) ItemSvcInterface {
	return &itemService{repo: repo}
}

func (s *itemService) Create(ctx context.Context, item *model.Item) (*model.Item, error) {
	result, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *itemService) GetById(ctx context.Context, id string) (*model.Item, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *itemService) GetAll(ctx context.Context) ([]*model.Item, error) {
	result, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *itemService) Update(ctx context.Context, id string, updateFields *model.Item) (int, error) {
	result, err := s.repo.Update(ctx, id, updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *itemService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}
