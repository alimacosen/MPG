package service

import (
	"context"
	model "mpg/characters/internal/model"
	repo "mpg/characters/internal/repository"
)

type CharacterService struct {
	repo repo.CharacterRepository
}


func NewCharacterService(repo repo.CharacterRepository) *CharacterService {
	return &CharacterService{repo: repo}
}

func (s *CharacterService) Create(ctx context.Context, character *model.Character) (*model.Character, error) {
	result, err := s.repo.Create(ctx, character)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CharacterService) GetById(ctx context.Context, id string) (*model.Character, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CharacterService) Update(ctx context.Context, id string, updateFields *model.UpdateFields) (int, error) {
	result, err := s.repo.Update(ctx, id, *updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CharacterService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CharacterService) UpdateInventoryId(ctx context.Context, id string, inventoryId string) (int, error) {
	result, err := s.repo.Update(ctx, id, &model.InventoryFields{InventoryId: inventoryId})
	if err != nil {
		return 0, err
	}
	return result, nil
}