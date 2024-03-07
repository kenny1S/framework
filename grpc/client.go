package grpc

import (
	"github.com/kenny1S/framework/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(servername string) (*grpc.ClientConn, error) {
	service, err := config.AgentHealthService(servername)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(service, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
