package service

import (
	"context"
	model "mpg/characters/internal/model"
	repo "mpg/characters/internal/repository"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	inventoryservice "mpg/inventories/gen/inventory_service"
)

type CharacterSvcInterface interface {
	Create(ctx context.Context, character *model.Character, inventoryClient *inventoryservicec.Client) (*model.Character, error)
	GetById(ctx context.Context, id string) (*model.Character, error)
	Update(ctx context.Context, id string, updateFields *model.Character) (int, error)
	Delete(ctx context.Context, id string, inventoryClient *inventoryservicec.Client) (int, error)
}

type characterService struct {
	repo repo.CharacterRepository
}


func NewCharacterService(repo repo.CharacterRepository) CharacterSvcInterface {
	return &characterService{repo: repo}
}

func (s *characterService) Create(ctx context.Context, character *model.Character, inventoryClient *inventoryservicec.Client) (*model.Character, error) {
	defaultHealth := 100
	defaultExp := 0
	character.Health = &defaultHealth
	character.Experience = &defaultExp

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

func (s *characterService) createInventory(ctx context.Context, characterPreliminary *model.Character, inventoryClient *inventoryservicec.Client) (*model.Character, error) {
	createInventoryRpc := inventoryClient.CreateInventory()

	createInventoryRes, err := createInventoryRpc(ctx, &inventoryservice.CreateInventoryPayload{UserID: characterPreliminary.ID})
	if err != nil {
		return nil, err
	}

	InventoryId := createInventoryRes.(*inventoryservice.Inventory).ID

	cnt, err := s.updateInventoryId(ctx, characterPreliminary.ID, InventoryId)
	if err != nil || cnt != 1 {
		return nil, err
	}

	characterPreliminary.InventoryID = &InventoryId

	return characterPreliminary, nil
}

func (s *characterService) GetById(ctx context.Context, id string) (*model.Character, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *characterService) Update(ctx context.Context, id string, updateFields *model.Character) (int, error) {
	result, err := s.repo.Update(ctx, id, *updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *characterService) Delete(ctx context.Context, id string, inventoryClient *inventoryservicec.Client) (int, error) {
	character, err := s.GetById(ctx, id)
	if err != nil {
		return 0, err
	}
	inventoryId := *character.InventoryID
	cnt, err := s.deleteInventory(ctx, inventoryId, inventoryClient)
	if err != nil || cnt != 1 {
		return 0, err
	}

	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *characterService) deleteInventory(ctx context.Context, inventoryId string, inventoryClient *inventoryservicec.Client) (int, error) {
	deleteInventoryRpc := inventoryClient.DeleteInventory()

	deleteInventoryRes, err := deleteInventoryRpc(ctx, &inventoryservice.DeleteInventoryPayload{ID: inventoryId})
	if err != nil {
		return 0, err
	}
	deleteCnt := deleteInventoryRes.(int)
	return deleteCnt, nil
}

func (s *characterService) updateInventoryId(ctx context.Context, id string, inventoryId string) (int, error) {
	result, err := s.repo.Update(ctx, id, &model.Character{InventoryID: &inventoryId})
	if err != nil {
		return 0, err
	}
	return result, nil
}