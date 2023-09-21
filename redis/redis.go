package redis

import "github.com/go-redis/redis"

// стоит ли вынести в структуру 3 значения?
// и тут ли должен быть сам редис
func RedisNew(addr string, pas string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pas,
		DB:       db,
	})
}
	