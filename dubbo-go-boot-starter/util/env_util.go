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

package util

import (
	"fmt"
	"os"
	"strings"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/common"
)

func GetEnvAndTrim(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func PresetEnv() {
	common.ConfigPath = GetEnvAndTrim(common.ApplicationConfigFilePathKey) // 通过环境变量获取应用配置文件
	common.DubboConfigPath = GetEnvAndTrim(common.DubboConfigFilePathKey)  // Dubbo服务配置文件

	if common.ConfigPath == "" {
		common.ConfigPath = common.DefaultApplicationConfigFilePath // 默认配置文件
	}
	if common.DubboConfigPath == "" { // 默认配置文件
		err := os.Setenv(common.DubboConfigFilePathKey, common.DefaultApplicationConfigFilePath)
		if err != nil {
			logger.Error(err)
		} else {
			common.DubboConfigPath = common.DefaultApplicationConfigFilePath
		}
	}
}

func CheckFile(path string, pathKey string) (file *os.File, err error) {
	file, err = os.Open(path)
	if err != nil {
		logger.Error(err)
		err = fmt.Errorf("环境变量[%s]所配置的文件地址[%s]不存在", pathKey, path)
	}
	return
}
