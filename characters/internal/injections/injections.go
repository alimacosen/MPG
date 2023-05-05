package injections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mpg/characters/clients/inventories/grpc"
	"mpg/characters/internal/config"
	"mpg/characters/internal/database"
	_ "mpg/characters/internal/database"
	"mpg/characters/internal/repository"
	"mpg/characters/internal/service"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	"net/url"
)

type Instances struct {
	cfg *config.DBConfig
	mongoClient *mongo.Client
	repo repository.CharacterRepository
	svc *service.CharacterService
	inventoryClient *inventoryservicec.Client
}

func NewInstances(logger *log.Logger) (*Instances, error) {
	dbCfg := config.NewDBConfig()
	mongoClient, _ := database.NewConnection(logger, dbCfg.DBConnectionString)
	repo := repository.NewCharacterRepository(logger, mongoClient, "CharacterDB")
	svc := service.NewCharacterService(repo)

	inventorySvcCfg := config.NewInventoryServiceConfig()

	parsedURL, err := url.Parse(inventorySvcCfg.Url)
	if err != nil {
		logger.Println("parsedURL error: ", err)
		return &Instances{cfg: dbCfg, mongoClient: mongoClient, repo: repo, svc: svc, inventoryClient: nil}, err
	}

	inventoryClient, err := grpc.NewClient(parsedURL)
	if err != nil {
		logger.Println("NewClient error: ", err)
		return &Instances{cfg: dbCfg, mongoClient: mongoClient, repo: repo, svc: svc, inventoryClient: nil}, err
	}

	return &Instances{cfg: dbCfg, mongoClient: mongoClient, repo: repo, svc: svc, inventoryClient: inventoryClient}, nil
}

func (i *Instances) GetSvc() *service.CharacterService {
	return i.svc
}

func (i *Instances) GetRepo() repository.CharacterRepository {
	return i.repo
}

func (i *Instances) GetInventoryClient() *inventoryservicec.Client {
	return i.inventoryClient
}

