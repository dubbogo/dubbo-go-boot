package common

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/model"
)

var (
	Config             model.ApplicationConfig
	DubboGoMiddlewares []extend.DubboGoMiddlewareI

	ConfigPath      string
	DubboConfigPath string
)
