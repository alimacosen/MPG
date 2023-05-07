package entryapi

import (
	"context"
	"log"
	characterservice "mpg/characters/gen/character_service"
	entrycharacterservice "mpg/entry/gen/entry_character_service"
	"mpg/entry/internal/handler"
)

// EntryCharacterService service example implementation.
// The example methods log the requests and return zero values.
type entryCharacterServicesrvc struct {
	logger *log.Logger
	entryHandler *handler.EntryHandler
}

// NewEntryCharacterService returns the EntryCharacterService service
// implementation.
func NewEntryCharacterService(logger *log.Logger) entrycharacterservice.Service {
	return &entryCharacterServicesrvc{logger, handler.GetEntryHandler(logger)}
}

// CreateCharacter implements creatCharacter.
func (s *entryCharacterServicesrvc) CreateCharacter(ctx context.Context, p *entrycharacterservice.CreateCharacterPayload) (res *entrycharacterservice.Character, err error) {
	createCharacter := s.entryHandler.GetInstance().GetCharacterClient().CreateCharacter()
	payload := &characterservice.CreateCharacterPayload{
		Name:        p.Name,
		Description: p.Description,
	}
	result, err := createCharacter(ctx, payload)
	if err != nil || result == nil {
		return nil, err
	}
	character := result.(*characterservice.Character)
	return &entrycharacterservice.Character{
		ID: character.ID,
		Name: character.Name,
		Description: character.Description,
		Health: character.Health,
		Experience: character.Experience,
		InventoryID: character.InventoryID,
	}, nil
}

// GetCharacter implements getCharacter.
func (s *entryCharacterServicesrvc) GetCharacter(ctx context.Context, p *entrycharacterservice.GetCharacterPayload) (res *entrycharacterservice.Character, err error) {
	getCharacter := s.entryHandler.GetInstance().GetCharacterClient().GetCharacter()
	payload := &characterservice.GetCharacterPayload{
		ID:        p.ID,
	}
	result, err := getCharacter(ctx, payload)
	if err != nil || result == nil {
		return nil, err
	}
	character := result.(*characterservice.Character)
	return &entrycharacterservice.Character{
		ID: character.ID,
		Name: character.Name,
		Description: character.Description,
		Health: character.Health,
		Experience: character.Experience,
		InventoryID: character.InventoryID,
	}, nil
}

// UpdateCharacter implements updateCharacter.
func (s *entryCharacterServicesrvc) UpdateCharacter(ctx context.Context, p *entrycharacterservice.UpdateCharacterPayload) (res int, err error) {
	updateCharacter := s.entryHandler.GetInstance().GetCharacterClient().UpdateCharacter()
	payload := &characterservice.UpdateCharacterPayload{
		ID:        p.ID,
		Name: p.Name,
		Description: p.Description,
		Health: p.Health,
		Experience: p.Experience,
	}
	result, err := updateCharacter(ctx, payload)
	if err != nil || result == nil {
		return 0, err
	}
	cnt := result.(int)
	return cnt, err
}

// DeleteCharacter implements deleteCharacter.
func (s *entryCharacterServicesrvc) DeleteCharacter(ctx context.Context, p *entrycharacterservice.DeleteCharacterPayload) (res int, err error) {
	deleteCharacter := s.entryHandler.GetInstance().GetCharacterClient().DeleteCharacter()
	payload := &characterservice.DeleteCharacterPayload{
		ID:        p.ID,
	}
	result, err := deleteCharacter(ctx, payload)
	if err != nil || result == nil {
		return 0, err
	}
	cnt := result.(int)
	return cnt, err
}
