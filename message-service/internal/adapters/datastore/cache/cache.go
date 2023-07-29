package cache

import (
	"context"
	"log"
	"time"

	messageProto "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/core/domain/message-service/proto"
	"github.com/golang/protobuf/proto"
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

type ChatMessages struct {
	Messages []Message
}

type Message struct {
	ID        string
	ChatID    string // ChatID is the ID between two users joined
	SenderID  string
	Content   string
	TimeStamp int
}

// Keys in cache-> (chatid:json string)

// When saving messages to cache, i want to store it as a JSON string
func (c *RedisCache) Save(msg *messageProto.Message) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("error marshalling struct into protobuf: %v", err)
	}
	ctx := context.Background()

	ttl := 300 * time.Second
	if err := c.client.Set(ctx, "test-key", b, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) Get(key string) ([]byte, error) {
	ctx := context.Background()

	res, err := c.client.Get(ctx, "test-key").Result()
	if err != nil {
		log.Printf("error getting string from cache: %v", err)
		return nil, err
	}

	return []byte(res), nil
}
