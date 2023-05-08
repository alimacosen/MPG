# Multiplayer Game


## Introduction

### Requirements 
This is a take-home assessment project. The task is to build a microservice application that stores characters and their item inventories for a multiplayer game.  
It is required to use [Goa](https://goa.design/), a design-first microservice framework with [Golang](https://go.dev/) to finish the entire task.
There are 4 microservices in total:
> An HTTP/JSON front entry service which provides an API to manipulate the characters, their inventories, and the items that exist

> A GRPC back service that handles CRUD operations for the characters and their attributes

> A GRPC back service that handles CRUD operations for the characters’ inventories

> A GRPC back service that handles CRUD operations for the items that exist and their attributes

> The front entry service should not have any state of its own and should call the appropriate back services via GRPC to implement its operations. The back services may store their state in memory.

### Database Schemas

The entry service doesn't have a DB schema, and the rest 3 services, which are Character, Inventory, and Item, have their own schema.

![schemaER](readme_files/schema.png)


### Architecture

The overall architecture is as the image below shows.

![architecture](readme_files/mpg.png)

## Installation


## Commands

mongodb://localhost:27017


go build ./cmd/characters_server  && go build ./cmd/characters_server-cli

./characters_server

./character_cli --url="grpc://localhost:8060" character-service create-character --message '{"name": "caifanglei", "description": "newuser"}'

./characters_server-cli --url="grpc://localhost:8060" character-service get-character --message '{"id": "645311b1f04abf6cd260b1cf"}'

./characters_server-cli --url="grpc://localhost:8060" character-service update-character --message '{"id": "645311b1f04abf6cd260b1cf", "name": "newname"}'

./characters_server-cli --url="grpc://localhost:8060" character-service delete-character --message '{"id": "645488667a280b1de71c110a"}'


go build ./cmd/inventory_server  && go build ./cmd/inventory_server-cli

./inventory_server

./inventory_server-cli  --url="grpc://localhost:8070" inventory-service create-inventory --message '{"user_Id": "645312d9575a8f2e8c2fa43e"}'

./inventory_server-cli  --url="grpc://localhost:8070" inventory-service get-inventory --message '{"id": "6454707f78df717203b20bf7"}'

./inventory_server-cli  --url="grpc://localhost:8070" inventory-service update-inventory --message '{"id": "6454707f78df717203b20bf7", "items_Id": ["111", "322"]}'

./inventory_server-cli  --url="grpc://localhost:8070" inventory-service delete-inventory --message '{"id": "64548cfeba9ad5b9bb96a837"}'


8080

go build ./cmd/item_server  && go build ./cmd/item_server-cli

./item_server

./item_server-cli  --url="grpc://localhost:8080" item-service create-item --message '{"name": "sword", "description": "very sharp", "healing": 10, "damage": 50, "protection": 35}'
./item_server-cli  --url="grpc://localhost:8080" item-service create-item --message '{"name": "shield", "description": "very hard", "healing": 15, "damage": 10, "protection": 65}'

./item_server-cli  --url="grpc://localhost:8080" item-service get-item --message '{"id": "6454e1881416cb009192ae31"}'

./item_server-cli  --url="grpc://localhost:8080" item-service update-item --message '{"id": "6454e1881416cb009192ae31", "name": "111shield", "description": "111very hard", "healing": 17, "damage": 17, "protection": 67}'

./item_server-cli  --url="grpc://localhost:8080" item-service delete-item --message '{"id": "6454e1af1416cb009192ae32"}'


# entry

./entry_server-cli --url="http://localhost:8050" entry-character-service creat-character --body '{"description": "user desc 1", "name": "user1"}'

./entry_server-cli --url="http://localhost:8050" entry-character-service get-character --id "6454b38fe5ffb9fa6b99f893"

./entry_server-cli --url="http://localhost:8050" entry-character-service update-character --id "64551f8c34b1c02c61a00c12" --body '{"id": "64551f8c34b1c02c61a00c12", "description": "trtr", "name": "E1E1E1", "health": 42}'

./entry_server-cli --url="http://localhost:8050" entry-character-service delete-character --id "64552546fad4971b156e2534"


./entry_server-cli --url="http://localhost:8050" entry-inventory-service get-inventory --id "6454707f78df7t7203b2ybf7"

./entry_server-cli --url="http://localhost:8050" entry-inventory-service update-inventory --id "6454b38f70ce8b433693251c" --body '{"itemsId": ["1234", "23456"]}'


./entry_server-cli --url="http://localhost:8050" entry-item-service creat-item --body '{"name": "swordd", "description": "very sharp", "healing": 10, "damage": 50, "protection": 35}'

./entry_server-cli --url="http://localhost:8050" entry-item-service get-item --id "6454e1b61416cb009192ae33"

./entry_server-cli --url="http://localhost:8050" entry-item-service update-item --id "6454e31e83890c44515b2341" --body '{"name": "swordd", "description": "very 1 sharp", "healing": 11, "damage": 51, "protection": 31}'

./entry_server-cli --url="http://localhost:8050" entry-item-service delete-item --id "6454ebbd1f80c4a34b0e4799"
