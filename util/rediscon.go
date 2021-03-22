package util

import (
	"fmt"
	"kumparan/config"

	"github.com/go-redis/redis"
)

type CacheConnection struct {
	Redis *redis.Client
}

func NewCacheConnection(config *config.AppConfig) *CacheConnection {
	var db CacheConnection
	RedisConnect(&db, config)

	return &db
}

func RedisConnect(db *CacheConnection, config *config.AppConfig) {
	address := fmt.Sprintf("%v:%v", config.Cache.Address, config.Cache.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Cache.Password,
		DB:       config.Cache.DBNumber,
	})

	db.Redis = client
}
