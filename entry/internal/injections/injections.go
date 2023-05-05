package injections

import (
	"log"
	characterservice "mpg/characters/gen/grpc/character_service/client"
	charactergrpc "mpg/entry/clients/characters/grpc"
	inventorygrpc "mpg/entry/clients/inventories/grpc"
	itemgrpc "mpg/entry/clients/items/grpc"
	"mpg/entry/internal/config"
	inventoryservice "mpg/inventories/gen/grpc/inventory_service/client"
	itemservice "mpg/items/gen/grpc/item_service/client"
	"net/url"
)

type Instances struct {
	characterClient *characterservice.Client
	inventoryClient *inventoryservice.Client
	itemClient *itemservice.Client
}

func NewInstances(logger *log.Logger) *Instances {
	characterClient, err := newCharacterClient(logger)
	if err != nil {
		logger.Println(err)
		return nil
	}
	inventoryClient, err := newInventoryClient(logger)
	if err != nil {
		logger.Println(err)
		return nil
	}
	itemClient, err := newItemClient(logger)
	if err != nil {
		logger.Println(err)
		return nil
	}

	return &Instances {
		characterClient: characterClient,
		inventoryClient: inventoryClient,
		itemClient: itemClient,
	}
}

func newCharacterClient(logger *log.Logger) (*characterservice.Client, error) {
	characterSvcCfg := config.NewCharacterServiceConfig()
	characterSvcParsedURL, err := url.Parse(characterSvcCfg.Url)
	if err != nil {
		logger.Println("parsedURL error: ", err)
	}

	characterClient, err := charactergrpc.NewClient(characterSvcParsedURL)
	if err != nil {
		logger.Println("NewClient error: ", err)
		return nil, err
	}
	return characterClient, nil
}

func newInventoryClient(logger *log.Logger) (*inventoryservice.Client, error) {
	inventorySvcCfg := config.NewInventoryServiceConfig()
	inventorySvcParsedURL, err := url.Parse(inventorySvcCfg.Url)
	if err != nil {
		logger.Println("parsedURL error: ", err)
	}

	inventoryClient, err := inventorygrpc.NewClient(inventorySvcParsedURL)
	if err != nil {
		logger.Println("NewClient error: ", err)
		return nil, err
	}
	return inventoryClient, nil
}

func newItemClient(logger *log.Logger) (*itemservice.Client, error) {
	itemSvcCfg := config.NewItemServiceConfig()
	itemSvcParsedURL, err := url.Parse(itemSvcCfg.Url)
	if err != nil {
		logger.Println("parsedURL error: ", err)
	}

	itemClient, err := itemgrpc.NewClient(itemSvcParsedURL)
	if err != nil {
		logger.Println("NewClient error: ", err)
		return nil, err
	}
	return itemClient, nil
}


func (i *Instances) GetCharacterClient() *characterservice.Client {
	return i.characterClient
}

func (i *Instances) GetInventoryClient() *inventoryservice.Client {
	return i.inventoryClient
}

func (i *Instances) GetItemClient() *itemservice.Client {
	return i.itemClient
}