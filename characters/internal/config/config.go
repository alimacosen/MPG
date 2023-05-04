package config

type DBConfig struct {
	DBConnectionString string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DBConnectionString: "mongodb://localhost:27017",
	}
}

