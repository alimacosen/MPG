# Multiplayer Game

## Installation


mongodb://localhost:27017


go build ./cmd/characters_server  && go build ./cmd/characters_server-cli

./characters_server

./characters_server-cli --url="grpc://localhost:8060" character-service create-character --message '{"name": "caifanglei", "description": "newuser"}'

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