package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	"net/url"
)


func NewClient(url *url.URL) (*inventoryservicec.Client, error) {
	conn, err := grpc.Dial(url.Host, grpc.WithTransportCredentials(insecure.NewCredentials())) // TODO Add interceptors and options
	if err != nil {
		return nil, err
	}
	c := inventoryservicec.NewClient(conn)
	return c, nil
}