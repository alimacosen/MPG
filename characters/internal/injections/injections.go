package injections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mpg/characters/internal/config"
	"mpg/characters/internal/database"
	_ "mpg/characters/internal/database"
	"mpg/characters/internal/repository"
	"mpg/characters/internal/service"
)

type Instances struct {
	cfg *config.DBConfig
	mongoClient *mongo.Client
	repo repository.CharacterRepository
	svc *service.CharacterService
}

func NewInstances(logger *log.Logger) *Instances {
	cfg := config.NewDBConfig()
	mongoClient, _ := database.NewConnection(logger, cfg.DBConnectionString)
	repo := repository.NewCharacterRepository(logger, mongoClient, "CharacterDB")
	svc := service.NewCharacterService(repo)

	return &Instances{cfg: cfg, mongoClient: mongoClient, repo: repo, svc: svc}
}

func (i *Instances) GetSvc() *service.CharacterService {
	return i.svc
}

func (i *Instances) GetRepo() repository.CharacterRepository {
	return i.repo
}

