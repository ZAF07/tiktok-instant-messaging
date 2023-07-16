package cachemanager

/*
	This package handles the initialisation of the the redis client (NOT THE ACTUAL IMPLEMENTATION)

	This is where we use redis-go library and make a connection to our redis instance, passing the connection to the cache adapter package
*/

type CacheManager struct{}
