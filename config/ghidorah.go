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

type Config interface {
	// Prefix config prefix
	Prefix() string

	// Init init config
	Init() error

	// Order load order
	Order() int
}

var configs = make(map[string]Config)

func SetConfig(name string, config Config) {
	configs[name] = config
}

func GetConfigs() map[string]Config {
	//keys := make([]string, len(configs))
	//i := 0
	//for key := range configs {
	//	keys[i] = key
	//	i++
	//}
	//sort.Slice(keys, func(i, j int) bool {
	//	return configs[keys[i]].Order() > configs[keys[j]].Order()
	//})
	//
	//l := list.New()
	//for _, key := range keys {
	//	l.PushFront(configs[key])
	//}

	return configs
}

type Ghidorah struct {
	// 数据库配置
	Database map[string]*Database `json:"database"`

	// 日志配置
	Logger *Logger `json:"logger"`
}
