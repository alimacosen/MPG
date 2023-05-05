package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	characterservicec "mpg/characters/gen/grpc/character_service/client"
	"net/url"
)


func NewClient(url *url.URL) (*characterservicec.Client, error) {
	conn, err := grpc.Dial(url.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := characterservicec.NewClient(conn)
	return c, nil
}