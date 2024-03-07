package redis

import (
	"2108a/high-five/home/day12/framework/config"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

var Res *redis.Client

type redisConfig struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
}
type Val struct {
	Redis redisConfig `json:"Redis"`
}

func withClient(servername string, hand func(cli *redis.Client) error) error {
	newRedis := new(Val)
	getConfig, err := config.GetConfig(servername, "DEFAULT_GROUP")
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(getConfig), &newRedis)
	if err != nil {
		return err
	}
	Res = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", newRedis.Redis.Host, newRedis.Redis.Port),
	})
	defer Res.Close()
	err = hand(Res)
	if err != nil {
		return err
	}
	return nil

}
func Exists(ctx context.Context, servername string, key string) (bool, error) {
	var data int64
	var err error
	err = withClient(servername, func(cli *redis.Client) error {
		data, err = cli.Exists(ctx, key).Result()
		return err
	})
	if err != nil {
		return false, err
	}
	if data > 0 {
		return true, nil
	}
	return false, nil
}
func GetRedis(ctx context.Context, servername string, key string) (string, error) {
	var data string
	var err error
	err = withClient(servername, func(cli *redis.Client) error {
		data, err = cli.Get(ctx, key).Result()
		return err
	})
	if err != nil {
		return "", err
	}
	return data, nil
}
func SetKey(ctx context.Context, servername string, key string, val interface{}, duration time.Duration) error {
	return withClient(servername, func(cli *redis.Client) error {
		return cli.Set(ctx, key, val, duration).Err()

	})
}
