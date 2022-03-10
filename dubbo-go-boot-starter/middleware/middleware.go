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
