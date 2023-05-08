package service

import (
	"context"
	"mpg/items/internal/model"
	repo "mpg/items/internal/repository"
)

type ItemSvcInterface interface {
	GetById(ctx context.Context, idList []string) ([]*model.Item, error)
	GetAll(ctx context.Context) ([]*model.Item, error)
	Create(ctx context.Context, item *model.Item) (*model.Item, error)
	Update(ctx context.Context, id string, updateFields *model.Item) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type ItemService struct {
	repo repo.ItemRepository
}

func NewItemService(repo repo.ItemRepository) ItemSvcInterface {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(ctx context.Context, item *model.Item) (*model.Item, error) {
	result, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ItemService) GetById(ctx context.Context, idList []string) ([]*model.Item, error) {
	result, err := s.repo.FindByID(ctx, idList)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ItemService) GetAll(ctx context.Context) ([]*model.Item, error) {
	result, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ItemService) Update(ctx context.Context, id string, updateFields *model.Item) (int, error) {
	result, err := s.repo.Update(ctx, id, updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *ItemService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}
