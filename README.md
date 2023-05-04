# Multiplayer Game

## Installation


mongodb://localhost:27017

go build ./cmd/characters_server  && go build ./cmd/characters_server-cli

./characters_server

./characters_server-cli --url="grpc://localhost:8060" character-service create-character --message '{"name": "caifanglei", "description": "newuser"}'

./characters_server-cli --url="grpc://localhost:8060" character-service get-character --message '{"id": "6453151e3d27604906a6e689"}'

./characters_server-cli --url="grpc://localhost:8060" character-service update-character --message '{"id": "645311b1f04abf6cd260b1cf", "name": "newname"}'

./characters_server-cli --url="grpc://localhost:8060" character-service delete-character --message '{"id": "6453151e3d27604906a6e689"}'