// Code generated by goa v3.11.3, DO NOT EDIT.
//
// InventoryService gRPC server types
//
// Command:
// $ goa gen mpg/inventories/design

package server

import (
	inventory_servicepb "mpg/inventories/gen/grpc/inventory_service/pb"
	inventoryservice "mpg/inventories/gen/inventory_service"

	goa "goa.design/goa/v3/pkg"
)

// NewCreateInventoryPayload builds the payload of the "createInventory"
// endpoint of the "InventoryService" service from the gRPC request type.
func NewCreateInventoryPayload(message *inventory_servicepb.CreateInventoryRequest) *inventoryservice.CreateInventoryPayload {
	v := &inventoryservice.CreateInventoryPayload{
		UserID: message.UserId,
	}
	return v
}

// NewProtoCreateInventoryResponse builds the gRPC response type from the
// result of the "createInventory" endpoint of the "InventoryService" service.
func NewProtoCreateInventoryResponse(result *inventoryservice.Inventory) *inventory_servicepb.CreateInventoryResponse {
	message := &inventory_servicepb.CreateInventoryResponse{
		Id:     result.ID,
		UserId: result.UserID,
	}
	if result.ItemsID != nil {
		message.ItemsId = make([]string, len(result.ItemsID))
		for i, val := range result.ItemsID {
			message.ItemsId[i] = val
		}
	}
	return message
}

// NewGetInventoryPayload builds the payload of the "getInventory" endpoint of
// the "InventoryService" service from the gRPC request type.
func NewGetInventoryPayload(message *inventory_servicepb.GetInventoryRequest) *inventoryservice.GetInventoryPayload {
	v := &inventoryservice.GetInventoryPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoGetInventoryResponse builds the gRPC response type from the result
// of the "getInventory" endpoint of the "InventoryService" service.
func NewProtoGetInventoryResponse(result *inventoryservice.Inventory) *inventory_servicepb.GetInventoryResponse {
	message := &inventory_servicepb.GetInventoryResponse{
		Id:     result.ID,
		UserId: result.UserID,
	}
	if result.ItemsID != nil {
		message.ItemsId = make([]string, len(result.ItemsID))
		for i, val := range result.ItemsID {
			message.ItemsId[i] = val
		}
	}
	return message
}

// NewUpdateInventoryPayload builds the payload of the "updateInventory"
// endpoint of the "InventoryService" service from the gRPC request type.
func NewUpdateInventoryPayload(message *inventory_servicepb.UpdateInventoryRequest) *inventoryservice.UpdateInventoryPayload {
	v := &inventoryservice.UpdateInventoryPayload{
		ID: message.Id,
	}
	if message.ItemsId != nil {
		v.ItemsID = make([]string, len(message.ItemsId))
		for i, val := range message.ItemsId {
			v.ItemsID[i] = val
		}
	}
	return v
}

// NewProtoUpdateInventoryResponse builds the gRPC response type from the
// result of the "updateInventory" endpoint of the "InventoryService" service.
func NewProtoUpdateInventoryResponse(result int) *inventory_servicepb.UpdateInventoryResponse {
	message := &inventory_servicepb.UpdateInventoryResponse{}
	message.Field = int32(result)
	return message
}

// NewDeleteInventoryPayload builds the payload of the "deleteInventory"
// endpoint of the "InventoryService" service from the gRPC request type.
func NewDeleteInventoryPayload(message *inventory_servicepb.DeleteInventoryRequest) *inventoryservice.DeleteInventoryPayload {
	v := &inventoryservice.DeleteInventoryPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoDeleteInventoryResponse builds the gRPC response type from the
// result of the "deleteInventory" endpoint of the "InventoryService" service.
func NewProtoDeleteInventoryResponse(result int) *inventory_servicepb.DeleteInventoryResponse {
	message := &inventory_servicepb.DeleteInventoryResponse{}
	message.Field = int32(result)
	return message
}

// ValidateUpdateInventoryRequest runs the validations defined on
// UpdateInventoryRequest.
func ValidateUpdateInventoryRequest(message *inventory_servicepb.UpdateInventoryRequest) (err error) {
	if message.ItemsId == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("itemsId", "message"))
	}
	return
}
