package config

import (
	"2108a/high-five/home/day12/framework/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"net"
	"strconv"
	"time"
)

const Consul_Key = "consul_index"

type Consul struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port"`
}

func getConfig(servername string) (*Consul, error) {
	s := new(Consul)
	config, err := GetConfig(servername, "DEFAULT_GROUP")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(config), &s)
	return s, err
}
func GetIp() (ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range addrs {
		ipNet, isVailIpNet := addr.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = append(ip, ipNet.IP.String())
			}
		}

	}
	return ip
}

func RegisterConsul(servername string, port int) error {
	config, err := getConfig(servername)
	if err != nil {
		return err
	}
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", config.Ip, config.Port),
	})
	if err != nil {
		return err
	}
	ip := GetIp()
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    servername,
		Tags:    []string{"Grpc"},
		Port:    port,
		Address: ip[0],
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			GRPC:                           fmt.Sprintf("%v:%v", ip[0], port),
			DeregisterCriticalServiceAfter: "10s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func getIndex(ctx context.Context, servername string, indexLen int) (int, error) {
	exists, err := redis.Exists(ctx, servername, Consul_Key)
	if err != nil {
		return 0, err
	}
	if exists {
		getRedis, err := redis.GetRedis(ctx, servername, Consul_Key)
		if err != nil {
			return 0, err
		}
		index, err := strconv.Atoi(getRedis)
		if err != nil {
			return 0, err
		}
		index += 1
		if index >= indexLen {
			index = 0
		}
		err = redis.SetKey(ctx, servername, Consul_Key, index, time.Minute*1)
		if err != nil {
			return 0, err
		}
		return index, nil
	}
	err = redis.SetKey(ctx, servername, Consul_Key, 0, time.Minute*1)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func AgentHealthService(ctx context.Context, servername string) (string, error) {
	config, err := getConfig(servername)
	if err != nil {
		return "", err
	}
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", config.Ip, config.Port),
	})
	if err != nil {
		return "", err
	}
	name, i, err := client.Agent().AgentHealthServiceByName(servername)
	if err != nil {
		return "", err
	}
	if name != "passing" {
		return "", fmt.Errorf("is not health service")
	}
	index, err := getIndex(ctx, servername, len(i))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:%v", i[index].Service.Address, i[index].Service.Port), nil
}
