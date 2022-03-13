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

package component

import (
	"os"
	"os/signal"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/config"
)

func ObserveSignal(duration time.Duration, beforeShutdown func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, config.ShutdownSignals...)

	for {
		select {
		case sig := <-signals:
			logger.Infof("get signal %s, applicationConfig will shutdown.", sig)
			// gracefulShutdownOnce.Do(func() {
			time.AfterFunc(duration, func() {
				logger.Warn("Shutdown gracefully timeout, applicationConfig will shutdown immediately. ")
				os.Exit(0)
			})

			if beforeShutdown != nil {
				beforeShutdown()
			}
			// those signals' original behavior is exit with dump ths stack, so we try to keep the behavior
			os.Exit(0)
		}
	}
}
