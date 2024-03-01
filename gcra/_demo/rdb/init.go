package rdb

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var rdbOnce = sync.Once{}

func RedisInit() {
	rdbOnce.Do(func() {
		Client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})
}
