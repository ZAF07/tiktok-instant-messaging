package cache

import (
	"log"

	httpdomain "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/domain/http_domain"
)

/*
	Driven Adapter

	This package implements the cache port/interface

	It's purpose is to store recent messages in-memory for fast retrieval
*/

type Cache struct{}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Save(msg httpdomain.Message) {
	log.Printf("Saving to cache!!: %+v\n", msg)
}
