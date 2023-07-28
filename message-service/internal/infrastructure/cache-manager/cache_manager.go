package cachemanager

/*
	This package handles the initialisation of the the redis client (NOT THE ACTUAL IMPLEMENTATION)

	This is where we use redis-go library and make a connection to our redis instance, passing the connection to the cache adapter package
*/

import (
	"context"
	"log"

	"github.com/ZAF07/tiktok-instant-messaging/message-service/config"
	"github.com/redis/go-redis/v9"
)

type CacheManager struct {
	Client *redis.Client
}

func NewRedisClient() *CacheManager {
	client := initRedisClient()

	return &CacheManager{
		Client: client,
	}
}

func initRedisClient() *redis.Client {
	configs, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error getting configs in redis manager: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.GetCacheHost(),
		Password: configs.GetCachePassword(), // no password set
		DB:       configs.GetCachedatabase(), // use default DB
	})

	ctx := context.Background()
	er := rdb.Ping(ctx)
	if err != nil {
		log.Fatalf("error connectingto redis: %v", er)
	}
	rdb.Conn()
	return rdb
}
