package cache

import "log"

/*
	Driven Adapter

	This package implements the cache port/interface

	It's purpose is to store recent messages in-memory for fast retrieval
*/

type Cache struct{}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Save() {
	log.Println("Saving to cache!!")
}
