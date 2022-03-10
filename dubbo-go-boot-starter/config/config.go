package config

import (
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/common"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/util"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadConfig() (err error) {
	util.PresetEnv()

	var configFile *os.File
	configFile, err = util.CheckFile(common.ConfigPath, common.ApplicationConfigFilePathKey)
	if err != nil {
		return
	}
	err = yaml.NewDecoder(configFile).Decode(&common.Config)
	if err != nil {
		return
	}
	_, err = util.CheckFile(common.DubboConfigPath, common.DubboConfigFilePathKey)
	return
}
