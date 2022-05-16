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

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/extension"
	"github.com/dubbogo/dubbo-go-boot/database"
)

func init() {
	extension.SetDatabase("mysql", newDriver)
}

func newDriver(config *core.URL) (*database.Database, error) {
	//dbConfig := &model.DatabaseConfig{}
	//err := util.ParseConfig(config, "database", dbConfig)
	//if err != nil || dbConfig == nil {
	//	clogger.Warn(err)
	//	clogger.Warn("please add database config")
	//	return nil
	//}
	//
	//host := dbConfig.Host
	//port := dbConfig.Port
	//configDatabase := dbConfig.Database
	//username := dbConfig.Username
	//password := dbConfig.Password
	//

	dialector := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, configDatabase))
	if dialector == nil {
		return nil
	}
	instance, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: databaseLogger(),
	})
	if err != nil {
		return nil
	}

	return nil, nil
}
