package service

import (
	"context"
	model "mpg/characters/internal/model"
	repo "mpg/characters/internal/repository"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	inventoryservice "mpg/inventories/gen/inventory_service"
)

type CharacterService struct {
	repo repo.CharacterRepository
}


func NewCharacterService(repo repo.CharacterRepository) *CharacterService {
	return &CharacterService{repo: repo}
}

func (s *CharacterService) Create(ctx context.Context, character *model.Character, inventoryClient *inventoryservicec.Client) (*model.Character, error) {
	result, err := s.repo.Create(ctx, character)
	if err != nil {
		return nil, err
	}

	characterPreliminary, err := s.createInventory(ctx, result, inventoryClient)
	if err != nil {
		return nil, err
	}

	return characterPreliminary, nil
}

func (s *CharacterService) createInventory(ctx context.Context, characterPreliminary *model.Character, inventoryClient *inventoryservicec.Client) (*model.Character, error) {
	createInventoryRpc := inventoryClient.CreateInventory()

	createInventoryRes, err := createInventoryRpc(ctx, &inventoryservice.CreateInventoryPayload{UserID: characterPreliminary.ID})
	if err != nil {
		return nil, err
	}

	InventoryId := createInventoryRes.(*inventoryservice.Inventory).ID

	cnt, err := s.UpdateInventoryId(ctx, characterPreliminary.ID, InventoryId)
	if err != nil {
		return nil, err
	}

	if cnt != 1 {
		// TODO
	}
	characterPreliminary.InventoryID = InventoryId

	return characterPreliminary, nil
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

func (s *CharacterService) Delete(ctx context.Context, id string, inventoryClient *inventoryservicec.Client) (int, error) {
	character, err := s.GetById(ctx, id)
	if err != nil {
		return 0, err
	}
	inventoryId := character.InventoryID
	cnt, err := s.deleteInventory(ctx, inventoryId, inventoryClient)
	if err != nil {
		return 0, err
	}
	if cnt != 1 {
		// TODO
	}

	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CharacterService) deleteInventory(ctx context.Context, inventoryId string, inventoryClient *inventoryservicec.Client) (int, error) {
	deleteInventoryRpc := inventoryClient.DeleteInventory()

	deleteInventoryRes, err := deleteInventoryRpc(ctx, &inventoryservice.DeleteInventoryPayload{ID: inventoryId})
	if err != nil {
		return 0, err
	}
	deleteCnt := deleteInventoryRes.(int)
	return deleteCnt, nil
}

func (s *CharacterService) UpdateInventoryId(ctx context.Context, id string, inventoryId string) (int, error) {
	result, err := s.repo.Update(ctx, id, &model.InventoryFields{InventoryId: inventoryId})
	if err != nil {
		return 0, err
	}
	return result, nil
}