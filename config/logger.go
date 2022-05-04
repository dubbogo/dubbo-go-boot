/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"

	"github.com/creasty/defaults"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/constant"
	"github.com/dubbogo/dubbo-go-boot/core/extension"
	"github.com/dubbogo/dubbo-go-boot/logger"
)

func init() {
	SetConfig("logger", &Logger{})
}

type Logger struct {
	// log driver
	Driver string `default:"zap" json:"driver"`

	// log level
	Level string `default:"info" json:"level"`
}

func (l *Logger) toURL() *core.URL {
	address := fmt.Sprintf("%s://%s", l.Driver, l.Level)
	u, _ := core.NewURL(address,
		core.WithParamsValue(constant.LoggerLevelKey, l.Level),
		core.WithParamsValue(constant.LoggerDriverKey, l.Driver),
	)
	return u
}

func (l *Logger) Prefix() string {
	return "logger"
}

func (l *Logger) Init() error {
	var (
		log logger.Logger
		err error
	)
	if err = defaults.Set(l); err != nil {
		return err
	}
	if log, err = extension.GetLogger(l.Driver, l.toURL()); err != nil {
		return err
	}
	logger.SetLog(log)
	return nil
}
