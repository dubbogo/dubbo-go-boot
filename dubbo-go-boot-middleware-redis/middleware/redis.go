package middleware

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	startModel "github.com/dubbogo/dubbo-go-boot-starter/model"
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-redis/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-redis/model"
	"github.com/go-redis/redis/v8"
)

var (
	rds = &redisComponent{}
)

func init() {
	middleware.RegisterMiddleware(rds)
}

type redisComponent struct {
}

func (c *redisComponent) Setup(config startModel.ApplicationConfig, hooks []extend.DubboGoMiddlewareSetupHook) error {
	rdConfig := &model.RedisConfig{}
	err := util.ParseConfig(config, "redis", rdConfig)
	if err != nil {
		logger.Warn(err)
		logger.Warn("please add redis config")
		return nil
	}

	if rdConfig == nil {
		logger.Warn("please add redis config")
		return nil
	}

	for _, v := range hooks {
		if vv, ok := v.(*RedisSetupHook); ok {
			vv.Hook()
		}
	}

	component.RedisComponent.Addr = fmt.Sprintf("%s:%d", rdConfig.Host, rdConfig.Port)
	component.RedisComponent.Password = rdConfig.Password
	component.RedisComponent.MaxRetries = rdConfig.MaxRetries

	component.RedisComponent.DbMap = make(map[int]*redis.Client)

	dbIndex := rdConfig.DefaultDB
	defaultClient := util.NewRedisDb(component.RedisComponent.Addr, component.RedisComponent.Password, dbIndex, component.RedisComponent.MaxRetries)
	component.RedisComponent.Redis = defaultClient
	component.RedisComponent.DbMap[dbIndex] = defaultClient
	return nil
}

func (c *redisComponent) IsAsync() bool {
	return false
}

func (c *redisComponent) Shutdown() {
	for k, r := range component.RedisComponent.DbMap {
		go shutdown(r)
		delete(component.RedisComponent.DbMap, k)
	}
	rds = nil
}

func shutdown(r *redis.Client) {
	if r != nil {
		_ = r.Close()
	}
}
