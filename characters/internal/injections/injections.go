package injections

import (
	"characters/internal/config"
	"characters/internal/database"
	_ "characters/internal/database"
	"characters/internal/repository"
	"characters/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

