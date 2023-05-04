package model

import (
	characterservice "characters/gen/character_service"
)


type Character characterservice.Character

func (c *Character) GetId() string {
	return c.ID
}

func (c *Character) SetId(id string)  {
	c.ID = id
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) SetName(name string)  {
	c.Name = name
}

func (c *Character) GetDescription() string {
	return c.Description
}

func (c *Character) SetDescription(description string)  {
	c.Description = description
}

func (c *Character) GetHealth() int {
	return c.Health
}

func (c *Character) SetHealth(health int)  {
	c.Health = health
}