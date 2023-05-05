package config

type DBConfig struct {
	DBConnectionString string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DBConnectionString: "mongodb://localhost:27017",
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
