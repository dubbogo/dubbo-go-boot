package middleware

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/common"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/extend"
	"sync"
)

var (
	lock = &middleware{}
)

type middleware struct {
	sync.Mutex
}

func RegisterMiddleware(middleware extend.DubboGoMiddlewareI) {
	lock.Lock()
	defer lock.Unlock()
	common.DubboGoMiddlewares = append(common.DubboGoMiddlewares, middleware)
}

func Setup(hooks []extend.DubboGoMiddlewareSetupHook) (err error) {
	for _, c := range common.DubboGoMiddlewares {
		m := c
		if m.IsAsync() {
			go func() {
				_ = setup(m, hooks)
			}()
		} else {
			err = setup(m, hooks)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func setup(m extend.DubboGoMiddlewareI, hooks []extend.DubboGoMiddlewareSetupHook) (err error) {
	err = m.Setup(common.Config, hooks)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func Shutdown() {
	for _, c := range common.DubboGoMiddlewares {
		c.Shutdown()
	}
}
