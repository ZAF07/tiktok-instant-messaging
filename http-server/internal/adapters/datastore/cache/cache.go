package cache

import (
	"log"

	httpdomain "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/domain/http_domain"
	"github.com/redis/go-redis/v9"
)

/*
	Driven Adapter

	This package implements the cache port/interface

	It's purpose is to store recent messages in-memory for fast retrieval
*/

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{}
}

func (c *RedisCache) Save(msg httpdomain.Message) {
	log.Printf("Saving to cache!!: %+v\n", msg)
}
