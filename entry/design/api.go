package design

import . "goa.design/goa/v3/dsl"

var _ = API("entryAPI", func() {
	Title("Entry Service")
	Description("HTTP service for a multiple players game")

	Server("EntryServer", func() {
		Description("EntryServer hosts the Character, Inventory, Item and swagger services.")
		Services("EntryCharacterService", "EntryInventoryService", "EntryItemService")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8050/mpg")
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
