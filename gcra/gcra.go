package gcra

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimit struct {
	rdb     *redis.Client
	prefix  string
	gcra    *redis.Script
	timeout time.Duration
}

func NewRateLimiter(rdb *redis.Client, prefix string, timeout time.Duration) *RateLimit {
	gcra := redis.NewScript(gcraScript)
	return &RateLimit{
		rdb:     rdb,
		gcra:    gcra,
		prefix:  prefix,
		timeout: timeout,
	}
}

func (r *RateLimit) Take(ctx context.Context, key string, burst int, interval float64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	values, err := r.gcra.Run(ctx, r.rdb, []string{r.key(key)}, burst, interval).Result()
	if err != nil {
		log.Println("error while executing GCRA script")
		return false, err
	}

	return values.([]interface{})[0].(int64) == 1, nil
}

func (r *RateLimit) key(key string) string {
	return fmt.Sprintf("%s:%s", r.prefix, key)
}
