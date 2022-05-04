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
	"github.com/dubbogo/dubbo-go-boot/database"
)

var databases = make(map[string]func(config *core.URL) (database.Database, error))

func GetDatabase(name string, config *core.URL) (database.Database, error) {
	if databases[name] == nil {
		return nil, errors.Errorf("database for %s driver does not exist. "+
			"please make sure that you have imported the package "+
			"github.com/dubbogo/dubbo-go-boot/database/%s",
			name, name)
	}
	return databases[name](config)
}

func SetDatabase(name string, driver func(config *core.URL) (database.Database, error)) {
	databases[name] = driver
}
