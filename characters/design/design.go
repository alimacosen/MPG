package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("charactersAPI", func() {
	Title("Characters Service")
	Description("Service for characters management")
	Server("charactersServer", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8060")
		})
	})
})

var Character = Type("Character", func() {
	Field(1, "id", String, "UUId of the character", func() {
		Meta("rpc:tag", "1")
	})
	Field(2, "name", String, "Name of the character", func() {
		Meta("rpc:tag", "2")
	})
	Field(3, "description", String, "Description of the character", func() {
		Meta("rpc:tag", "3")
	})
	Field(4, "health", Int, "The amount of health value of the character", func() {
		Meta("rpc:tag", "4")
	})
	Field(5, "experience", Int, "The amount of experience value of the character", func() {
		Meta("rpc:tag", "5")
	})
	Field(6, "inventoryId", String, "UUId of the inventory record", func() {
		Meta("rpc:tag", "6")
	})
	Required("id", "name", "description", "health", "experience", "inventoryId")
})

var _ = Service("CharacterService", func() {
	Description("The character service performs CRUD operations for characters and their attributes.")

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
		Error("create_invalid_args", String, "Invalid arguments. Required: name  Optional: description ")
		GRPC(func() {
			Response(CodeOK)
			Response("create_invalid_args", CodeInvalidArgument)
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
		GRPC(func() {
			Response(CodeOK)
			Response("get_invalid_args", CodeInvalidArgument)
			Response("get_no_match", CodeNotFound)
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
		Error("update_invalid_args", String, "Invalid arguments. Required: id  Optional: name, description, health, experience ")
		Error("update_no_match", String, "No character matched given criteria")
		GRPC(func() {
			Response(CodeOK)
			Response("update_invalid_args", CodeInvalidArgument)
			Response("update_no_match", CodeNotFound)
		})
	})

	Method("deleteCharacter", func() {
		Payload(func() {
			Field(1, "id", String, "UUId of the Character", func() {
				Meta("rpc:tag", "1")
			})
			Required("id")
		})
		Error("delete_invalid_args", String, "Invalid arguments. Required: id ")
		Error("delete_no_match", String, "No character matched given criteria")
		Result(Int)
		GRPC(func() {
			Response(CodeOK)
			Response("delete_invalid_args", CodeInvalidArgument)
			Response("delete_no_match", CodeNotFound)
		})
	})
})
