package cachemanager

/*
	This package handles the initialisation of the the redis client (NOT THE ACTUAL IMPLEMENTATION)

	This is where we use redis-go library and make a connection to our redis instance, passing the connection to the cache adapter package
*/

import (
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.Conn()
	return rdb
}
