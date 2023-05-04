package service

import (
	model "characters/internal/model"
	repo "characters/internal/repository"
	"context"
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

//func (s *CharacterService) GetAll() ([]*model.Character, error) {
//	return s.repo.GetAll()
//}
//
//func (s *CharacterService) Update(book *model.Character) error {
//	return s.repo.Update(book)
//}
//
//func (s *CharacterService) Delete(id int64) error {
//	return s.repo.Delete(id)
//}