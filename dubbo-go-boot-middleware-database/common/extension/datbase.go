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

package extension

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/database"
	"github.com/pkg/errors"
)

var databases = make(map[string]database.Database)

func GetDatabase(name string) (database.Database, error) {
	if databases[name] == nil {
		return nil, errors.Errorf("Database for %s does not exist. "+
			"please make sure that you have imported the package "+
			"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/database/%s",
			name, name)
	}
	return databases[name], nil
}

func SetDatabase(name string, diver database.Database) {
	databases[name] = diver
}
