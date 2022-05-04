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

package bootstrap

import (
	"encoding/json"

	"github.com/spf13/viper"

	"github.com/dubbogo/dubbo-go-boot/config"
	"github.com/dubbogo/dubbo-go-boot/core/constant"
	"github.com/dubbogo/dubbo-go-boot/logger"
	_ "github.com/dubbogo/dubbo-go-boot/logger/zap"
)

func Run(conf *Option) error {
	logger.Infof("dubbo-go boot version %s", constant.Version)

	viper.SetConfigName(conf.name)
	viper.SetConfigType(conf.suffix)
	viper.AddConfigPath(conf.path)

	logger.Infof("start load config %s", conf.GetConfig())
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("read config err=%v", err)
		return err
	}
	return Init()
}

func Init() error {
	var (
		err  error
		data []byte
	)

	for _, conf := range config.GetConfigs() {
		// init database
		if database, ok := conf.(*config.Database); ok {
			for k, v := range viper.GetStringMap(conf.Prefix()) {

				if data, err = json.Marshal(v); err != nil {
					return err
				}
				if err = json.Unmarshal(data, database); err != nil {
					return err
				}
				if err = database.InitDatabase(k); err != nil {
					return err
				}
			}
			continue
		}
		if err = viper.UnmarshalKey(conf.Prefix(), conf); err != nil {
			return err
		}
		if err = conf.Init(); err != nil {
			return err
		}
	}
	return nil
}
