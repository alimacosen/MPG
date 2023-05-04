package charactersapi

import (
	characterservice "characters/gen/character_service"
	"characters/internal/handler"
	"context"
	"log"
)

// CharacterService service example implementation.
// The example methods log the requests and return zero values.
type characterServicesrvc struct {
	logger *log.Logger
	characterHandler *handler.CharacterHandler
}

// NewCharacterService returns the CharacterService service implementation.
func NewCharacterService(logger *log.Logger) characterservice.Service {
	return &characterServicesrvc{logger, handler.NewCharacterHandler(logger)}
}

// CreateCharacter implements createCharacter.
func (s *characterServicesrvc) CreateCharacter(ctx context.Context, p *characterservice.CreateCharacterPayload) (res *characterservice.Character, err error) {
	res, err = s.characterHandler.CreateCharacter(ctx, p)
	return
}

// GetCharacter implements getCharacter.
func (s *characterServicesrvc) GetCharacter(ctx context.Context, p *characterservice.GetCharacterPayload) (res *characterservice.Character, err error) {
	res = &characterservice.Character{}
	s.logger.Print("characterService.getCharacter")
	return
}

// UpdateCharacterAttributes implements updateCharacterAttributes.
func (s *characterServicesrvc) UpdateCharacterAttributes(ctx context.Context, p *characterservice.UpdateCharacterAttributesPayload) (res *characterservice.Character, err error) {
	res = &characterservice.Character{}
	s.logger.Print("characterService.updateCharacterAttributes")
	return
}

// DeleteCharacter implements deleteCharacter.
func (s *characterServicesrvc) DeleteCharacter(ctx context.Context, p *characterservice.DeleteCharacterPayload) (res *characterservice.Character, err error) {
	res = &characterservice.Character{}
	s.logger.Print("characterService.deleteCharacter")
	return
}
