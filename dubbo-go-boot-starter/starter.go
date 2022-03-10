package starter

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/config"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/middleware"
	"time"
)

var (
	DefaultSurvivalTimeout = time.Duration(int(3e9))
)

type Starter struct {
	survivalTimeout time.Duration

	middlewareSetupHooks []extend.DubboGoMiddlewareSetupHook
}

func NewStarter() *Starter {
	return &Starter{
		survivalTimeout:      DefaultSurvivalTimeout,
		middlewareSetupHooks: make([]extend.DubboGoMiddlewareSetupHook, 0),
	}
}

func (s *Starter) SetSurvivalTimeout(survivalTimeout time.Duration) *Starter {
	s.survivalTimeout = survivalTimeout
	return s
}

func (s *Starter) SetMiddlewareSetupHooks(hooks ...extend.DubboGoMiddlewareSetupHook) *Starter {
	s.middlewareSetupHooks = hooks
	return s
}

func (s *Starter) AddMiddlewareSetupHooks(hooks ...extend.DubboGoMiddlewareSetupHook) *Starter {
	h := s.middlewareSetupHooks
	for _, v := range hooks {
		h = append(h, v)
	}
	s.middlewareSetupHooks = h
	return s
}

func (s *Starter) GetMiddlewareSetupHooks() []extend.DubboGoMiddlewareSetupHook {
	return s.middlewareSetupHooks
}

func (s *Starter) Start() (err error) {
	err = config.LoadConfig()
	if err != nil {
		return err
	}
	err = middleware.Setup(s.middlewareSetupHooks)
	if err != nil {
		return err
	}
	component.ObserveSignal(DefaultSurvivalTimeout, func() {
		middleware.Shutdown()
	})
	return nil
}
