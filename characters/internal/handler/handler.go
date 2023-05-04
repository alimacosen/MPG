package handler

import (
	characterservice "characters/gen/character_service"
	"characters/internal/injections"
	"characters/internal/model"
	"context"
	"log"
)

type CharacterHandler struct {
	logger *log.Logger
	instances *injections.Instances
}

func NewCharacterHandler(logger *log.Logger) *CharacterHandler {
	return &CharacterHandler{logger: logger, instances: injections.NewInstances()}
}

func (c *CharacterHandler) CreateCharacter(ctx context.Context, p *characterservice.CreateCharacterPayload) (res *characterservice.Character, err error) {
	name := p.Name
	if len(name) == 0 {
		return nil, characterservice.NoMatch("Name can not be an empty string")
	}
	description := *p.Description

	svc := c.instances.GetSvc()

	characterPreliminary := &model.Character{
		ID:          "",
		Name:        name,
		Description: description,
		Health:      0,
		Experience:  0,
		InventoryID: "",
	}

	//character.InventoryID = GetInventory().Id
	cp, err := svc.Create(ctx, characterPreliminary)
	if err != nil {
		return nil, err
	}

	res = convert(cp)
	return
}

func convert(c *model.Character) *characterservice.Character {
	return &characterservice.Character{
		ID: c.ID,
		Name: c.Name,
		Description: c.Description,
		Health: c.Health,
		Experience: c.Experience,
		InventoryID: c.InventoryID,
	}
}