// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client CLI support package
//
// Command:
// $ goa gen mpg/items/design

package client

import (
	"encoding/json"
	"fmt"
	item_servicepb "mpg/items/gen/grpc/item_service/pb"
	itemservice "mpg/items/gen/item_service"
)

// BuildCreateItemPayload builds the payload for the ItemService createItem
// endpoint from CLI flags.
func BuildCreateItemPayload(itemServiceCreateItemMessage string) (*itemservice.CreateItemPayload, error) {
	var err error
	var message item_servicepb.CreateItemRequest
	{
		if itemServiceCreateItemMessage != "" {
			err = json.Unmarshal([]byte(itemServiceCreateItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 6656640518575154227,\n      \"description\": \"Sit aliquid harum iure qui quos.\",\n      \"healing\": 6018419887003017801,\n      \"name\": \"Asperiores cupiditate et non ipsam blanditiis dolores.\",\n      \"protection\": 1946787761586667757\n   }'")
			}
		}
	}
	v := &itemservice.CreateItemPayload{
		Name:        message.Name,
		Description: message.Description,
		Damage:      int(message.Damage),
		Healing:     int(message.Healing),
		Protection:  int(message.Protection),
	}

	return v, nil
}

// BuildGetItemPayload builds the payload for the ItemService getItem endpoint
// from CLI flags.
func BuildGetItemPayload(itemServiceGetItemMessage string) (*itemservice.GetItemPayload, error) {
	var err error
	var message item_servicepb.GetItemRequest
	{
		if itemServiceGetItemMessage != "" {
			err = json.Unmarshal([]byte(itemServiceGetItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Ut eos fuga laborum.\"\n   }'")
			}
		}
	}
	v := &itemservice.GetItemPayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdateItemPayload builds the payload for the ItemService updateItem
// endpoint from CLI flags.
func BuildUpdateItemPayload(itemServiceUpdateItemMessage string) (*itemservice.UpdateItemPayload, error) {
	var err error
	var message item_servicepb.UpdateItemRequest
	{
		if itemServiceUpdateItemMessage != "" {
			err = json.Unmarshal([]byte(itemServiceUpdateItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 3012358727598942804,\n      \"description\": \"Quod dolor.\",\n      \"healing\": 8553807324101554941,\n      \"id\": \"Dolorum earum magnam sequi aliquid optio.\",\n      \"name\": \"Sint mollitia.\",\n      \"protection\": 8197344049870397300\n   }'")
			}
		}
	}
	v := &itemservice.UpdateItemPayload{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		damage := int(*message.Damage)
		v.Damage = &damage
	}
	if message.Healing != nil {
		healing := int(*message.Healing)
		v.Healing = &healing
	}
	if message.Protection != nil {
		protection := int(*message.Protection)
		v.Protection = &protection
	}

	return v, nil
}

// BuildDeleteItemPayload builds the payload for the ItemService deleteItem
// endpoint from CLI flags.
func BuildDeleteItemPayload(itemServiceDeleteItemMessage string) (*itemservice.DeleteItemPayload, error) {
	var err error
	var message item_servicepb.DeleteItemRequest
	{
		if itemServiceDeleteItemMessage != "" {
			err = json.Unmarshal([]byte(itemServiceDeleteItemMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Eaque debitis aut aut autem quia nesciunt.\"\n   }'")
			}
		}
	}
	v := &itemservice.DeleteItemPayload{
		ID: message.Id,
	}

	return v, nil
}
