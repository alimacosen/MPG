package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("inventoriesAPI", func() {
	Title("Inventory Service")
	Description("Service for inventories management")
	Server("inventoryServer", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8070")
		})
	})
})

var Inventory = Type("Inventory", func() {
	Field(1, "id", String, "UUId of the inventory", func() {
		Meta("rpc:tag", "1")
	})
	Field(2, "userId", String, "UUId of the character that this inventory belongs to", func() {
		Meta("rpc:tag", "2")
	})
	Field(3, "itemsId", ArrayOf(String), "Array of uuid of items", func() {
		Meta("rpc:tag", "3")
	})

	Required("id", "userId")
})

var _ = Service("InventoryService", func() {
	Description("The inventory service performs CRUD operations for inventories")

	Method("createInventory", func() {
		Payload(func() {
			Field(1, "userId", String, "UUId of the character that this inventory belongs to", func() {
				Meta("rpc:tag", "1")
			})
			Required("userId")
		})
		Result(Inventory)
		Error("create_invalid_args", String, "Invalid arguments. Required: userId ")
		GRPC(func() {
			Response(CodeOK)
			Response("create_invalid_args", CodeInvalidArgument)
		})
	})

	Method("getInventory", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the inventory", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})
		Result(Inventory)
		Error("get_invalid_args", String, "Invalid arguments. Required: id ")
		Error("get_no_match", String, "No inventory matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("get_invalid_args", CodeInvalidArgument)
			Response("get_no_match", CodeNotFound)
		})
	})

	Method("updateInventory", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Inventory", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "itemsId", ArrayOf(String), "Array of uuid of items", func() {
				Meta("rpc:tag", "2")
			})
			Required("id", "itemsId")
		})


		Result(Int)
		Error("update_invalid_args", String, "Invalid arguments. Required: id, itemsId ")
		Error("update_no_match", String, "No inventory matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("update_invalid_args", CodeInvalidArgument)
			Response("update_no_match", CodeNotFound)
		})
	})

	Method("deleteInventory", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Inventory", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})
		Result(Int)
		Error("delete_invalid_args", String, "Invalid arguments. Required: id ")
		Error("delete_no_match", String, "No inventory matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("delete_invalid_args", CodeInvalidArgument)
			Response("delete_no_match", CodeNotFound)
		})
	})
})
