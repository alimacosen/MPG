// Code generated by goa v3.11.3, DO NOT EDIT.
//
// CharacterService gRPC client types
//
// Command:
// $ goa gen characters/design

package client

import (
	characterservice "characters/gen/character_service"
	character_servicepb "characters/gen/grpc/character_service/pb"
)

// NewProtoCreateCharacterRequest builds the gRPC request type from the payload
// of the "createCharacter" endpoint of the "CharacterService" service.
func NewProtoCreateCharacterRequest(payload *characterservice.CreateCharacterPayload) *character_servicepb.CreateCharacterRequest {
	message := &character_servicepb.CreateCharacterRequest{
		Name:        payload.Name,
		Description: payload.Description,
	}
	return message
}

// NewCreateCharacterResult builds the result type of the "createCharacter"
// endpoint of the "CharacterService" service from the gRPC response type.
func NewCreateCharacterResult(message *character_servicepb.CreateCharacterResponse) *characterservice.Character {
	result := &characterservice.Character{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
		Health:      int(message.Health),
		Experience:  int(message.Experience),
		InventoryID: message.InventoryId,
	}
	return result
}

// NewProtoGetCharacterRequest builds the gRPC request type from the payload of
// the "getCharacter" endpoint of the "CharacterService" service.
func NewProtoGetCharacterRequest(payload *characterservice.GetCharacterPayload) *character_servicepb.GetCharacterRequest {
	message := &character_servicepb.GetCharacterRequest{
		Id: payload.ID,
	}
	return message
}

// NewGetCharacterResult builds the result type of the "getCharacter" endpoint
// of the "CharacterService" service from the gRPC response type.
func NewGetCharacterResult(message *character_servicepb.GetCharacterResponse) *characterservice.Character {
	result := &characterservice.Character{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
		Health:      int(message.Health),
		Experience:  int(message.Experience),
		InventoryID: message.InventoryId,
	}
	return result
}

// NewProtoUpdateCharacterRequest builds the gRPC request type from the payload
// of the "updateCharacter" endpoint of the "CharacterService" service.
func NewProtoUpdateCharacterRequest(payload *characterservice.UpdateCharacterPayload) *character_servicepb.UpdateCharacterRequest {
	message := &character_servicepb.UpdateCharacterRequest{
		Id:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
	}
	if payload.Health != nil {
		health := int32(*payload.Health)
		message.Health = &health
	}
	if payload.Experience != nil {
		experience := int32(*payload.Experience)
		message.Experience = &experience
	}
	return message
}

// NewUpdateCharacterResult builds the result type of the "updateCharacter"
// endpoint of the "CharacterService" service from the gRPC response type.
func NewUpdateCharacterResult(message *character_servicepb.UpdateCharacterResponse) int {
	result := int(message.Field)
	return result
}

// NewProtoDeleteCharacterRequest builds the gRPC request type from the payload
// of the "deleteCharacter" endpoint of the "CharacterService" service.
func NewProtoDeleteCharacterRequest(payload *characterservice.DeleteCharacterPayload) *character_servicepb.DeleteCharacterRequest {
	message := &character_servicepb.DeleteCharacterRequest{
		Id: payload.ID,
	}
	return message
}

// NewDeleteCharacterResult builds the result type of the "deleteCharacter"
// endpoint of the "CharacterService" service from the gRPC response type.
func NewDeleteCharacterResult(message *character_servicepb.DeleteCharacterResponse) int {
	result := int(message.Field)
	return result
}
