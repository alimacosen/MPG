package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("itemsAPI", func() {
	Title("Items Service")
	Description("Service for items management")
	Server("itemServer", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8080")
		})
	})
})

var Item = Type("Item", func() {
	Field(1, "id", String, "UUId of the item", func() {
		Meta("rpc:tag", "1")
	})
	Field(2, "name", String, "name of the item", func() {
		Meta("rpc:tag", "2")
	})
	Field(3, "description", String, "Description of the item", func() {
		Meta("rpc:tag", "3")
	})
	Field(4, "damage", Int, "The amount of damage the item can do", func() {
		Meta("rpc:tag", "4")
	})
	Field(5, "healing", Int, "The amount of healing the item can do", func() {
		Meta("rpc:tag", "5")
	})
	Field(6, "protection", Int, "The amount of protection the item can do", func() {
		Meta("rpc:tag", "6")
	})
	Required("id", "name", "description", "damage", "healing", "protection")
})

var _ = Service("ItemService", func() {
	Description("The item service performs CRUD operations for items")

	Method("createItem", func() {
		Payload(func() {
			Field(1, "name", String, "name of the item", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "description", String, "Description of the item", func() {
				Meta("rpc:tag", "3")
			})
			Field(3, "damage", Int, "The amount of damage the item can do", func() {
				Meta("rpc:tag", "4")
			})
			Field(4, "healing", Int, "The amount of healing the item can do", func() {
				Meta("rpc:tag", "5")
			})
			Field(5, "protection", Int, "The amount of protection the item can do", func() {
				Meta("rpc:tag", "6")
			})
			Required("name", "description", "damage", "healing", "protection")
		})
		Result(Item)
		Error("create_invalid_args", String, "Invalid arguments. Required: name, description, damage, healing, protection ")
		Error("create_duplicated_name", String, "Duplicated name. This item name already exists ")
		GRPC(func() {
			Response(CodeOK)
			Response("create_invalid_args", CodeInvalidArgument)
		})
	})

	Method("getItem", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the item", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})
		Result(Item)
		Error("get_invalid_args", String, "Invalid arguments. Required: id ")
		Error("get_no_match", String, "No item matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("get_invalid_args", CodeInvalidArgument)
			Response("get_no_match", CodeNotFound)
		})
	})

	Method("updateItem", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the item", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "description", String, "Description of the item", func() {
				Meta("rpc:tag", "3")
			})
			Field(3, "damage", Int, "The amount of damage the item can do", func() {
				Meta("rpc:tag", "4")
			})
			Field(4, "healing", Int, "The amount of healing the item can do", func() {
				Meta("rpc:tag", "5")
			})
			Field(5, "protection", Int, "The amount of protection the item can do", func() {
				Meta("rpc:tag", "6")
			})
			Required("id")
		})
		Result(Int)
		Error("update_invalid_args", String, "Invalid arguments. Required: id, itemsId ")
		Error("update_no_match", String, "No item matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("update_invalid_args", CodeInvalidArgument)
			Response("update_no_match", CodeNotFound)
		})
	})

	Method("deleteItem", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the item", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})
		Result(Int)
		Error("delete_invalid_args", String, "Invalid arguments. Required: id ")
		Error("delete_no_match", String, "No item matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("delete_invalid_args", CodeInvalidArgument)
			Response("delete_no_match", CodeNotFound)
		})
	})
})
