package design

import . "goa.design/goa/v3/dsl"

var _ = Service("EntryItemService", func() {
	Description("The entry item service communicates with items microservice.")
	HTTP(func() {
		Path("/item")
	})
	Method("creatItem", func() {
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
		Error("create_no_criteria", String, "Missing criteria")
		Error("create_invalid_args", String, "Invalid arguments. Required: name, description, damage, healing, protection ")
		Error("create_duplicated_name", String, "Duplicated name. This item name already exists ")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("create_no_criteria", StatusBadRequest)
			Response("create_invalid_args", StatusBadRequest)
			Response("create_duplicated_name", StatusBadRequest)

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
		Error("get_no_criteria", String, "Missing criteria ")
		Error("get_invalid_args", String, "Invalid arguments. Required: id ")
		Error("get_no_match", String, "No item matched given criteria")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("get_no_criteria", StatusBadRequest)
			Response("get_invalid_args", StatusBadRequest)
			Response("get_no_match", StatusNotFound)
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
		Error("update_no_criteria", String, "Missing criteria")
		Error("update_invalid_args", String, "Invalid arguments. Required: id  Optional: name, description, health, experience ")
		Error("update_no_match", String, "No item matched given criteria")
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusOK)
			Response("update_no_criteria", StatusBadRequest)
			Response("update_invalid_args", StatusBadRequest)
			Response("update_no_match", StatusNotFound)
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
		Error("delete_no_criteria", String, "Missing criteria ")
		Error("delete_invalid_args", String, "Invalid arguments. Required: id ")
		Error("delete_no_match", String, "No item matched given criteria")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
			Response("delete_invalid_args", StatusBadRequest)
			Response("delete_no_criteria", StatusBadRequest)
			Response("delete_no_match", StatusNotFound)
		})
	})
})
