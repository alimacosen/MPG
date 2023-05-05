package config


type CharacterServiceConfig struct {
	Url string
}

func NewCharacterServiceConfig() *CharacterServiceConfig {
	return &CharacterServiceConfig{
		Url: "grpc://localhost:8060",
	}
}

type InventoryServiceConfig struct {
	Url string
}

func NewInventoryServiceConfig() *InventoryServiceConfig {
	return &InventoryServiceConfig{
		Url: "grpc://localhost:8070",
	}
}

type ItemServiceConfig struct {
	Url string
}

func NewItemServiceConfig() *ItemServiceConfig {
	return &ItemServiceConfig{
		Url: "grpc://localhost:8080",
	}
}