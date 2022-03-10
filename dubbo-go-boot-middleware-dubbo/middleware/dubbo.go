package middleware

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	"github.com/dubbogo/dubbo-go-boot-starter/model"
	"time"
)

var (
	dubbo = &dubboComponent{}
)

func init() {
	middleware.RegisterMiddleware(dubbo)
}

type dubboComponent struct {
}

func (c *dubboComponent) Setup(_ model.ApplicationConfig, hooks []extend.DubboGoMiddlewareSetupHook) (err error) {
	for _, v := range hooks {
		if vv, ok := v.(*DubboSetupHook); ok {
			vv.Hook()
		}
	}

	var retry bool
	for {
		retry = false
		setup(func() {
			retry = true
		})
		if !retry {
			break
		}
		logger.Debug("Dubbo Go service load failed [retry after 5s]")
		time.Sleep(5 * time.Second)
	}
	logger.Debug("Dubbo Go service load succeed")

	return nil
}

func setup(hook func()) {
	resolved := false
	defer func() {
		if err := recover(); !(err == nil || resolved) {
			hook()
		}
	}()
	err := config.Load()
	if err != nil {
		resolved = true
		hook()
	}
}

func (c *dubboComponent) IsAsync() bool {
	return true
}

func (c *dubboComponent) Shutdown() {
	dubbo = nil
}
