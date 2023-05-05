// Code generated by goa v3.11.3, DO NOT EDIT.
//
// InventoryService gRPC server
//
// Command:
// $ goa gen mpg/inventories/design

package server

import (
	"context"
	"errors"
	inventory_servicepb "mpg/inventories/gen/grpc/inventory_service/pb"
	inventoryservice "mpg/inventories/gen/inventory_service"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the inventory_servicepb.InventoryServiceServer interface.
type Server struct {
	CreateInventoryH goagrpc.UnaryHandler
	GetInventoryH    goagrpc.UnaryHandler
	UpdateInventoryH goagrpc.UnaryHandler
	DeleteInventoryH goagrpc.UnaryHandler
	inventory_servicepb.UnimplementedInventoryServiceServer
}

// New instantiates the server struct with the InventoryService service
// endpoints.
func New(e *inventoryservice.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateInventoryH: NewCreateInventoryHandler(e.CreateInventory, uh),
		GetInventoryH:    NewGetInventoryHandler(e.GetInventory, uh),
		UpdateInventoryH: NewUpdateInventoryHandler(e.UpdateInventory, uh),
		DeleteInventoryH: NewDeleteInventoryHandler(e.DeleteInventory, uh),
	}
}

// NewCreateInventoryHandler creates a gRPC handler which serves the
// "InventoryService" service "createInventory" endpoint.
func NewCreateInventoryHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateInventoryRequest, EncodeCreateInventoryResponse)
	}
	return h
}

// CreateInventory implements the "CreateInventory" method in
// inventory_servicepb.InventoryServiceServer interface.
func (s *Server) CreateInventory(ctx context.Context, message *inventory_servicepb.CreateInventoryRequest) (*inventory_servicepb.CreateInventoryResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "createInventory")
	ctx = context.WithValue(ctx, goa.ServiceKey, "InventoryService")
	resp, err := s.CreateInventoryH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "create_invalid_args":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventory_servicepb.CreateInventoryResponse), nil
}

// NewGetInventoryHandler creates a gRPC handler which serves the
// "InventoryService" service "getInventory" endpoint.
func NewGetInventoryHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetInventoryRequest, EncodeGetInventoryResponse)
	}
	return h
}

// GetInventory implements the "GetInventory" method in
// inventory_servicepb.InventoryServiceServer interface.
func (s *Server) GetInventory(ctx context.Context, message *inventory_servicepb.GetInventoryRequest) (*inventory_servicepb.GetInventoryResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getInventory")
	ctx = context.WithValue(ctx, goa.ServiceKey, "InventoryService")
	resp, err := s.GetInventoryH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "get_invalid_args":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "get_no_match":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventory_servicepb.GetInventoryResponse), nil
}

// NewUpdateInventoryHandler creates a gRPC handler which serves the
// "InventoryService" service "updateInventory" endpoint.
func NewUpdateInventoryHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateInventoryRequest, EncodeUpdateInventoryResponse)
	}
	return h
}

// UpdateInventory implements the "UpdateInventory" method in
// inventory_servicepb.InventoryServiceServer interface.
func (s *Server) UpdateInventory(ctx context.Context, message *inventory_servicepb.UpdateInventoryRequest) (*inventory_servicepb.UpdateInventoryResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "updateInventory")
	ctx = context.WithValue(ctx, goa.ServiceKey, "InventoryService")
	resp, err := s.UpdateInventoryH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "update_invalid_args":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "update_no_match":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventory_servicepb.UpdateInventoryResponse), nil
}

// NewDeleteInventoryHandler creates a gRPC handler which serves the
// "InventoryService" service "deleteInventory" endpoint.
func NewDeleteInventoryHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteInventoryRequest, EncodeDeleteInventoryResponse)
	}
	return h
}

// DeleteInventory implements the "DeleteInventory" method in
// inventory_servicepb.InventoryServiceServer interface.
func (s *Server) DeleteInventory(ctx context.Context, message *inventory_servicepb.DeleteInventoryRequest) (*inventory_servicepb.DeleteInventoryResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "deleteInventory")
	ctx = context.WithValue(ctx, goa.ServiceKey, "InventoryService")
	resp, err := s.DeleteInventoryH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "delete_invalid_args":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "delete_no_match":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventory_servicepb.DeleteInventoryResponse), nil
}
