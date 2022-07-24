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

package mysql

import (
	"fmt"
	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/extension"
	"github.com/dubbogo/dubbo-go-boot/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	extension.SetDatabase("mysql", newMysqlDriver)
}

func newMysqlDriver(config *core.URL) (*database.Database, error) {
	host := config.Ip
	port := config.Port
	path := config.Path
	username := config.Username
	password := config.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, path)

	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	d := database.Database{}
	d.SetDriver(instance)
	database.SetDatabase("mysql", &d)

	return &d, nil
}
