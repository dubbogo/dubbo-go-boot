package component

import (
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	RedisComponent = &redisComponent{}
)

type redisComponent struct {
	sync.Mutex

	Addr       string
	Password   string
	MaxRetries int

	Redis *redis.Client

	DbMap map[int]*redis.Client
}

func GetRedis() *redis.Client {
	return RedisComponent.Redis
}

func GetRedisByIndex(dbIndex int) *redis.Client {
	RedisComponent.Lock()
	if client, notNil := RedisComponent.DbMap[dbIndex]; !notNil || client == nil {
		client = util.NewRedisDb(RedisComponent.Addr, RedisComponent.Password, dbIndex, RedisComponent.MaxRetries)
		RedisComponent.DbMap[dbIndex] = client
	}
	RedisComponent.Unlock()
	return RedisComponent.DbMap[dbIndex]
}
