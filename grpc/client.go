package grpc

import (
	"context"
	"github.com/kenny1S/framework/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(ctx context.Context, servername string) (*grpc.ClientConn, error) {
	service, err := config.AgentHealthService(ctx, servername)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(service, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
