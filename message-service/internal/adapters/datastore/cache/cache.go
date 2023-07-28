package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	httpdomain "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/http_domain"
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
	return &RedisCache{
		client: c,
	}
}

type usersMsg struct {
	messages []msg
}

type msg struct {
	chatID    string
	sender    string
	text      string
	timeStamp int
}

// Keys in cache-> (chatid:json string)

// When saving messages to cache, i want to store it as a JSON string
func (c *RedisCache) Save(msg httpdomain.Message) error {
	ctx := context.Background()

	ttl := 300 * time.Second
	if err := c.client.Set(ctx, "test-key", msg.Text, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()

	// 🚨 Testing to see how much data is used when a sting of such length is stored
	strLen := c.client.StrLen(ctx, "test-key")
	if se := strLen.Err(); se != nil {
		log.Fatalf("error getting strlen: %v", se)
	}
	fmt.Println("SIXE +++++> y ", strLen.Val())

	res := c.client.Get(ctx, "test-key")
	r, err := res.Result()
	if err != nil {
		log.Printf("error getting string from cache: %v", err)
		return "", err
	}
	return r, nil
}
