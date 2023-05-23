package storage

import (
	"horus-api/configs"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func GetServicesNames() []string {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	v, err := Redis.HGetAll("defense-services").Result()
	if err != nil {
		panic(err)
	}
	var servicesNames []string
	for k := range v {
		servicesNames = append(servicesNames, k)
	}
	return servicesNames
}

func SetServiceName(service string) error {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	err := Redis.HSet("defense-services", service, service).Err()
	if err != nil {
		return err
	}
	return nil
}

func RemoveServiceName(service string) error {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	err := Redis.HDel("defense-services", service).Err()
	if err != nil {
		return err
	}
	return nil
}

func SubscribeClient(url string) error {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	err := Redis.HSet("subscribe", url, url).Err()
	if err != nil {
		return err
	}
	return nil
}
