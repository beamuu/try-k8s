package repositories

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return RedisRepository{rdb: rdb}
}

func (r RedisRepository) Get(key string) (string, error) {
	return r.rdb.Get(key).Result()
}

func (r RedisRepository) Set(key string, value string, expiration time.Duration) error {
	return r.rdb.Set(key, value, expiration).Err()
}
