package grpc

import (
	"2108a/high-five/home/day12/framework/config"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(ctx context.Context, servername string)(*grpc.ClientConn,error) {
	service, err := config.AgentHealthService(ctx, servername)
	if err != nil {
		return nil,err
	}
	return grpc.Dial(service, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
