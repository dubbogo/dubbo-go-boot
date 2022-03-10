package extend

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/model"
)

type DubboGoMiddlewareI interface {
	Setup(model.ApplicationConfig, []DubboGoMiddlewareSetupHook) error

	IsAsync() bool

	Shutdown()
}

type DubboGoMiddlewareSetupHook interface {
	Hook()
}
