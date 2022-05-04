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
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/dubbogo/dubbo-go-boot/config"
)

func Run(conf *Option) error {
	viper.SetConfigName(conf.name)
	viper.SetConfigType(conf.suffix)
	viper.AddConfigPath(conf.path)

	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	return Init()
}

func Init() error {
	for _, conf := range config.GetConfigs() {
		if err := viper.UnmarshalKey(conf.Prefix(), conf); err != nil {
			return err
		}
		if err := conf.Init(); err != nil {
			return err
		}
	}
	return nil
}
