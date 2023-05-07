package design

import . "goa.design/goa/v3/dsl"

var _ = Service("EntryCharacterService", func() {
	Description("The entry character service communicates with characters microservice.")
	HTTP(func() {
		Path("/character")
	})
	Method("createCharacter", func() {
		Payload(func() {
			Field(1, "name", String, "Name of the Character", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "description", String, "Description of the Character", func() {
				Meta("rpc:tag", "2")
			})
			Required("name")
		})

		Result(Character)
		Error("create_no_criteria", String, "Missing criteria")
		Error("create_invalid_args", String, "Invalid arguments. Required: name  Optional: description ")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("create_no_criteria", StatusBadRequest)
			Response("create_invalid_args", StatusBadRequest)
		})
	})

	Method("getCharacter", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Character", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})

		Result(Character)
		Error("get_invalid_args", String, "Invalid arguments. Required: id ")
		Error("get_no_match", String, "No character matched given criteria")
		Error("get_no_criteria", String, "Missing criteria")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("get_no_criteria", StatusBadRequest)
			Response("get_invalid_args", StatusBadRequest)
			Response("get_no_match", StatusNotFound)
		})
	})

	Method("updateCharacter", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Character", func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "name", String, "Name of the Character", func() {
				Meta("rpc:tag", "2")
			})
			Field(3, "description", String, "Description of the Character", func() {
				Meta("rpc:tag", "3")
			})
			Field(4, "health", Int, "The amount of health value of the Character", func() {
				Meta("rpc:tag", "4")
			})
			Field(5, "experience", Int, "The amount of experience value of the Character", func() {
				Meta("rpc:tag", "5")
			})
			Required("id")
		})

		Result(Int)
		Error("no_criteria", String, "Missing criteria")
		Error("update_invalid_args", String, "Invalid arguments. Required: id  Optional: name, description, health, experience ")
		Error("update_no_match", String, "No character matched given criteria")
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusOK)
			Response("no_criteria", StatusBadRequest)
			Response("update_invalid_args", StatusBadRequest)
			Response("update_no_match", StatusNotFound)
		})
	})

	Method("deleteCharacter", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Character", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})

		Result(Int)
		Error("delete_invalid_args", String, "Invalid arguments. Required: id ")
		Error("delete_no_match", String, "No character matched given criteria")
		Error("no_criteria", String, "Missing criteria")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
			Response("delete_invalid_args", StatusBadRequest)
			Response("no_criteria", StatusBadRequest)
			Response("delete_no_match", StatusNotFound)
		})
	})
})
