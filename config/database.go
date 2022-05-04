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

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/constant"
)

func init() {
	SetConfig("database", &Database{})
}

type Database struct {
	// mongo、mysql
	Driver string `json:"driver"`

	// database url
	Url string `json:"url"`

	// database username
	Username string `yaml:"username"`

	// database password
	Password string `yaml:"password"`

	// database connect timeout
	Timeout string `default:"5s" json:"timeout"`
}

func (Database) Prefix() string {
	return "database"
}

func (d *Database) Init() error {
	fmt.Println(d.Username)
	fmt.Println(d.Driver)
	return nil
}

func (d *Database) toURL() (*core.URL, error) {
	address := fmt.Sprintf("%s://%s", d.Driver, d.Url)
	return core.NewURL(address,
		core.WithParamsValue(constant.DatabaseTimeoutKey, d.Timeout),
		core.WithParamsValue(constant.DatabaseKey, d.Driver),
	)
}
