package grpc

import (
	"encoding/json"
	"fmt"
	"github.com/kenny1S/framework/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

type Conf struct {
	Grpc struct {
		Ip   string `json:"Ip"`
		Port string `json:"Port"`
	} `json:"grpc"`
}

func getConfig(servername string) (*Conf, error) {
	s, err := config.GetConfig(servername, "DEFAULT_GROUP")
	if err != nil {
		return nil, err
	}
	c := new(Conf)
	err = json.Unmarshal([]byte(s), &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func RegisterGrpc(servername string, f func(server *grpc.Server)) error {
	conf, err := getConfig(servername)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(conf.Grpc.Port)
	if err != nil {
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Grpc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	err = config.RegisterConsul(servername, port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	f(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
