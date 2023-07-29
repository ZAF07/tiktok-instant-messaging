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

// When saving messages to cache, i want to store it as a JSON string
func (c *RedisCache) Save(msg *messageProto.Message) error {
	ctx := context.Background()
	ttl := 25 * time.Second

	textHistory := &messageProto.Messages{}

	// Get the value in cache with key == key (IF EXISTS)
	oldTextsBin, getErr := c.client.Get(ctx, "test-key").Bytes()
	if getErr == nil {
		log.Println("Found old keys in cache")
		// 💡 If there are old text content in the history
		// Unmarshall into proto
		unmarshalErr := proto.Unmarshal(oldTextsBin, textHistory)
		if unmarshalErr != nil {
			log.Printf("error unmarshalling into textHistory in cache.Save(): %v\n", unmarshalErr)
			return nil
		}
	}

	// Keep only up to 10 text history in the cache
	if len(textHistory.Messages) >= 2 {
		textHistory.Messages = append(textHistory.Messages[1:], msg)
	} else {
		textHistory.Messages = append(textHistory.Messages, msg)
	}

	b, err := proto.Marshal(textHistory)
	if err != nil {
		log.Fatalf("error marshalling struct into protobuf: %v", err)
		return err
	}

	if err := c.client.Set(ctx, "test-key", b, ttl).Err(); err != nil {
		log.Fatalf("error setting value to cache: %v", err)
		return err
	}
	return nil
}

func (c *RedisCache) Get(key string) ([]byte, error) {
	ctx := context.Background()

	res, err := c.client.Get(ctx, "test-key").Bytes()
	if err != nil {
		log.Printf("error getting string from cache: %v\n", err)
		return nil, err
	}

	return res, nil
}
