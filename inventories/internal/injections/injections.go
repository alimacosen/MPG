package injections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"inventories/internal/config"
	"inventories/internal/database"
	_ "inventories/internal/database"
	"inventories/internal/repository"
	"inventories/internal/service"
	"log"
)

type Instances struct {
	cfg *config.DBConfig
	mongoClient *mongo.Client
	repo repository.InventoryRepository
	svc *service.InventoryService
}

func NewInstances(logger *log.Logger) *Instances {
	cfg := config.NewDBConfig()
	mongoClient, _ := database.NewConnection(logger, cfg.DBConnectionString)
	repo := repository.NewInventoryRepository(logger, mongoClient, "InventoryDB")
	svc := service.NewInventoryService(repo)

	return &Instances{cfg: cfg, mongoClient: mongoClient, repo: repo, svc: svc}
}

func (i *Instances) GetSvc() *service.InventoryService {
	return i.svc
}

func (i *Instances) GetRepo() repository.InventoryRepository {
	return i.repo
}

