package injections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mpg/items/internal/config"
	"mpg/items/internal/database"
	_ "mpg/items/internal/database"
	"mpg/items/internal/repository"
	"mpg/items/internal/service"
)

type Instances struct {
	cfg *config.DBConfig
	mongoClient *mongo.Client
	repo repository.ItemRepository
	svc *service.InventoryService
}

func NewInstances(logger *log.Logger) *Instances {
	cfg := config.NewDBConfig()
	mongoClient, _ := database.NewConnection(logger, cfg.DBConnectionString)
	repo := repository.NewItemRepository(logger, mongoClient, "ItemDB")
	svc := service.NewItemService(repo)

	return &Instances{cfg: cfg, mongoClient: mongoClient, repo: repo, svc: svc}
}

func (i *Instances) GetSvc() *service.InventoryService {
	return i.svc
}

func (i *Instances) GetRepo() repository.ItemRepository {
	return i.repo
}

