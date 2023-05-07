// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC server types
//
// Command:
// $ goa gen mpg/items/design

package server

import (
	item_servicepb "mpg/items/gen/grpc/item_service/pb"
	itemservice "mpg/items/gen/item_service"
)

// NewCreateItemPayload builds the payload of the "createItem" endpoint of the
// "ItemService" service from the gRPC request type.
func NewCreateItemPayload(message *item_servicepb.CreateItemRequest) *itemservice.CreateItemPayload {
	v := &itemservice.CreateItemPayload{
		Name:        message.Name,
		Description: message.Description,
		Damage:      int(message.Damage),
		Healing:     int(message.Healing),
		Protection:  int(message.Protection),
	}
	return v
}

// NewProtoCreateItemResponse builds the gRPC response type from the result of
// the "createItem" endpoint of the "ItemService" service.
func NewProtoCreateItemResponse(result *itemservice.Item) *item_servicepb.CreateItemResponse {
	message := &item_servicepb.CreateItemResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Damage:      int32(result.Damage),
		Healing:     int32(result.Healing),
		Protection:  int32(result.Protection),
	}
	return message
}

// NewGetItemPayload builds the payload of the "getItem" endpoint of the
// "ItemService" service from the gRPC request type.
func NewGetItemPayload(message *item_servicepb.GetItemRequest) *itemservice.GetItemPayload {
	v := &itemservice.GetItemPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoGetItemResponse builds the gRPC response type from the result of the
// "getItem" endpoint of the "ItemService" service.
func NewProtoGetItemResponse(result *itemservice.Item) *item_servicepb.GetItemResponse {
	message := &item_servicepb.GetItemResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Damage:      int32(result.Damage),
		Healing:     int32(result.Healing),
		Protection:  int32(result.Protection),
	}
	return message
}

// NewProtoGetAllItemsResponse builds the gRPC response type from the result of
// the "getAllItems" endpoint of the "ItemService" service.
func NewProtoGetAllItemsResponse(result []*itemservice.Item) *item_servicepb.GetAllItemsResponse {
	message := &item_servicepb.GetAllItemsResponse{}
	message.Field = make([]*item_servicepb.Item, len(result))
	for i, val := range result {
		message.Field[i] = &item_servicepb.Item{
			Id:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Damage:      int32(val.Damage),
			Healing:     int32(val.Healing),
			Protection:  int32(val.Protection),
		}
	}
	return message
}

// NewUpdateItemPayload builds the payload of the "updateItem" endpoint of the
// "ItemService" service from the gRPC request type.
func NewUpdateItemPayload(message *item_servicepb.UpdateItemRequest) *itemservice.UpdateItemPayload {
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
	return v
}

// NewProtoUpdateItemResponse builds the gRPC response type from the result of
// the "updateItem" endpoint of the "ItemService" service.
func NewProtoUpdateItemResponse(result int) *item_servicepb.UpdateItemResponse {
	message := &item_servicepb.UpdateItemResponse{}
	message.Field = int32(result)
	return message
}

// NewDeleteItemPayload builds the payload of the "deleteItem" endpoint of the
// "ItemService" service from the gRPC request type.
func NewDeleteItemPayload(message *item_servicepb.DeleteItemRequest) *itemservice.DeleteItemPayload {
	v := &itemservice.DeleteItemPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoDeleteItemResponse builds the gRPC response type from the result of
// the "deleteItem" endpoint of the "ItemService" service.
func NewProtoDeleteItemResponse(result int) *item_servicepb.DeleteItemResponse {
	message := &item_servicepb.DeleteItemResponse{}
	message.Field = int32(result)
	return message
}
