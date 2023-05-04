package injections

import (
	"characters/internal/config"
	"characters/internal/database"
	_ "characters/internal/database"
	"characters/internal/repository"
	"characters/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Instances struct {
	cfg *config.Config
	mongoClient *mongo.Client
	repo repository.CharacterRepository
	svc *service.CharacterService
}

func NewInstances() *Instances {
	cfg := config.New()
	mongoClient, _ := database.NewConnection(cfg.DBConnectionString)
	repo := repository.NewCharacterRepository(mongoClient, "CharacterDB")
	svc := service.NewCharacterService(repo)

	return &Instances{cfg: cfg, mongoClient: mongoClient, repo: repo, svc: svc}
}

func (i *Instances) GetSvc() *service.CharacterService {
	return i.svc
}

func (i *Instances) GetRepo() repository.CharacterRepository {
	return i.repo
}

