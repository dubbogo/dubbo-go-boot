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

package database

import (
	"gorm.io/gorm"
)

type Database struct {
	driver interface{}
}

func (db *Database) SetDriver(driver *gorm.DB) {
	db.driver = driver
}

func (db *Database) GetDriver() *gorm.DB {
	return db.driver.(*gorm.DB)
}

var databases = make(map[string]*Database)

func SetDatabase(name string, db *Database) {
	databases[name] = db
}

func GetDatabase(name string) *Database {
	return databases[name]
}

func Ignore(name string) bool {
	if _, ok := databases[name]; ok {
		return true
	}
	return false
}
