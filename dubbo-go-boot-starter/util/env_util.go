package util

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/common"
	"os"
	"strings"
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
