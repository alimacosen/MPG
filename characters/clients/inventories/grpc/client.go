package grpc

import (
	"google.golang.org/grpc"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	"net/url"
)


type inventoryGrpcClient struct {
	Client *inventoryservicec.Client
}

func NewClient(url *url.URL) (*inventoryservicec.Client, error) {
	conn, err := grpc.Dial(url.Host) // TODO Add interceptors and options
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := inventoryservicec.NewClient(conn)
	return c, nil
}