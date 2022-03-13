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
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	"github.com/dubbogo/dubbo-go-boot-starter/model"
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
