package design

import . "goa.design/goa/v3/dsl"

var _ = Service("EntryInventoryService", func() {
	Description("The entry inventory service communicates with inventory microservice.")
	HTTP(func() {
		Path("/inventory")
	})

	Method("getInventory", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the inventory", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})

		Result(Inventory)
		Error("get_no_match", String, "No item matched given criteria")
		Error("get_no_criteria", String, "Missing criteria")
		Error("get_invalid_args", String, "Invalid arguments. Required: id ")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("get_no_criteria", StatusBadRequest)
			Response("get_invalid_args", StatusBadRequest)
			Response("get_no_match", StatusNotFound)
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
		Error("update_no_criteria", String, "Missing criteria")
		Error("update_invalid_args", String, "Invalid arguments. Required: id, itemsId ")
		Error("update_no_match", String, "No character matched given criteria")
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusOK)
			Response("update_no_criteria", StatusBadRequest)
			Response("update_invalid_args", StatusBadRequest)
			Response("update_no_match", StatusNotFound)
		})
	})

})
