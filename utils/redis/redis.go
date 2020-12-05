package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var redisClint Client

type Client struct {
	redis.UniversalClient
}

func newRedisClint(options *redis.UniversalOptions) {
	redisClint = Client{redis.NewUniversalClient(options)}
}

func InstanceRedisClint(options ...*redis.UniversalOptions) Client {
	once.Do(func() {
		newRedisClint(options[0])
	})
	return redisClint
}
