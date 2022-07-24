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

package extension

import (
	"github.com/pkg/errors"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/logger"
)

var logs = make(map[string]func(conf *core.URL) (logger.Logger, error))

func SetLogger(name string, v func(conf *core.URL) (logger.Logger, error)) {
	logs[name] = v
}

func GetLogger(name string, conf *core.URL) (logger.Logger, error) {
	if logs[name] == nil {
		return nil, errors.Errorf("logger for %s does not exist. "+
			"please make sure that you have imported the package "+
			"github.com/dubbogo/dubbo-go-boot/logger/%s", name, name)
	}
	return logs[name](conf)
}
