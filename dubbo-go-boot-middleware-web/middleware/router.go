/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

import (
	"fmt"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
)

import (
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	startModel "github.com/dubbogo/dubbo-go-boot-starter/model"
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-web/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-web/model"
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
