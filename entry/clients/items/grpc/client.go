package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	itemservicec "mpg/items/gen/grpc/item_service/client"
	"net/url"
)


func NewClient(url *url.URL) (*itemservicec.Client, error) {
	conn, err := grpc.Dial(url.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := itemservicec.NewClient(conn)
	return c, nil
}