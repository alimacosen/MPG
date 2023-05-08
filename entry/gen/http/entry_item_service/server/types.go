// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryItemService HTTP server types
//
// Command:
// $ goa gen mpg/entry/design

package server

import (
	entryitemservice "mpg/entry/gen/entry_item_service"

	goa "goa.design/goa/v3/pkg"
)

// CreateItemRequestBody is the type of the "EntryItemService" service
// "createItem" endpoint HTTP request body.
type CreateItemRequestBody struct {
	// name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The amount of damage the item can do
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// The amount of healing the item can do
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// The amount of protection the item can do
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateItemRequestBody is the type of the "EntryItemService" service
// "updateItem" endpoint HTTP request body.
type UpdateItemRequestBody struct {
	// name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The amount of damage the item can do
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// The amount of healing the item can do
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// The amount of protection the item can do
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// CreateItemResponseBody is the type of the "EntryItemService" service
// "createItem" endpoint HTTP response body.
type CreateItemResponseBody struct {
	// UUId of the item
	ID string `form:"id" json:"id" xml:"id"`
	// name of the item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the item
	Description string `form:"description" json:"description" xml:"description"`
	// The amount of damage the item can do
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// The amount of healing the item can do
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// The amount of protection the item can do
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// GetItemsResponseBody is the type of the "EntryItemService" service
// "getItems" endpoint HTTP response body.
type GetItemsResponseBody []*ItemResponse

// ItemResponse is used to define fields on response body types.
type ItemResponse struct {
	// UUId of the item
	ID string `form:"id" json:"id" xml:"id"`
	// name of the item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the item
	Description string `form:"description" json:"description" xml:"description"`
	// The amount of damage the item can do
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// The amount of healing the item can do
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// The amount of protection the item can do
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// NewCreateItemResponseBody builds the HTTP response body from the result of
// the "createItem" endpoint of the "EntryItemService" service.
func NewCreateItemResponseBody(res *entryitemservice.Item) *CreateItemResponseBody {
	body := &CreateItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewGetItemsResponseBody builds the HTTP response body from the result of the
// "getItems" endpoint of the "EntryItemService" service.
func NewGetItemsResponseBody(res []*entryitemservice.Item) GetItemsResponseBody {
	body := make([]*ItemResponse, len(res))
	for i, val := range res {
		body[i] = marshalEntryitemserviceItemToItemResponse(val)
	}
	return body
}

// NewCreateItemPayload builds a EntryItemService service createItem endpoint
// payload.
func NewCreateItemPayload(body *CreateItemRequestBody) *entryitemservice.CreateItemPayload {
	v := &entryitemservice.CreateItemPayload{
		Name:        *body.Name,
		Description: *body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}

	return v
}

// NewGetItemsPayload builds a EntryItemService service getItems endpoint
// payload.
func NewGetItemsPayload(ids []string) *entryitemservice.GetItemsPayload {
	v := &entryitemservice.GetItemsPayload{}
	v.Ids = ids

	return v
}

// NewUpdateItemPayload builds a EntryItemService service updateItem endpoint
// payload.
func NewUpdateItemPayload(body *UpdateItemRequestBody, id string) *entryitemservice.UpdateItemPayload {
	v := &entryitemservice.UpdateItemPayload{
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}
	v.ID = id

	return v
}

// NewDeleteItemPayload builds a EntryItemService service deleteItem endpoint
// payload.
func NewDeleteItemPayload(id string) *entryitemservice.DeleteItemPayload {
	v := &entryitemservice.DeleteItemPayload{}
	v.ID = id

	return v
}

// ValidateCreateItemRequestBody runs the validations defined on
// CreateItemRequestBody
func ValidateCreateItemRequestBody(body *CreateItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	return
}
