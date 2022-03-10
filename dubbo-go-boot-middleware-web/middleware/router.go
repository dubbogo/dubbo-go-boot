package middleware

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	startModel "github.com/dubbogo/dubbo-go-boot-starter/model"
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-web/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-web/model"
	"github.com/gin-gonic/gin"
)

var (
	web = &webComponent{}
)

func init() {
	middleware.RegisterMiddleware(web)
}

type webComponent struct {
}

func (c *webComponent) Setup(config startModel.ApplicationConfig, hooks []extend.DubboGoMiddlewareSetupHook) error {
	serverConfig := &model.ServerConfig{}
	err := util.ParseConfig(config, "server", serverConfig)
	if err != nil {
		logger.Warn(err)
		logger.Warn("please add server config")
		return nil
	}

	if serverConfig == nil {
		logger.Warn("please add server config")
		return nil
	}

	routerEngine := gin.New()

	hasHook := false
	for _, v := range hooks {
		if vv, ok := v.(*WebSetupHook); ok {
			vv.hook(routerEngine)
			hasHook = true
		}
	}

	if !hasHook {
		return fmt.Errorf("please add `WebSetupHook` into starter at first")
	}

	component.WebComponent.Router = routerEngine
	err = routerEngine.Run(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port))
	if err != nil {
		return err
	}

	return nil
}

func (c *webComponent) IsAsync() bool {
	return false
}

func (c *webComponent) Shutdown() {
	component.WebComponent = nil
	web = nil
}
