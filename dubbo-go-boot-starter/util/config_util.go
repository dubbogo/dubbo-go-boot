package util

import (
	"encoding/json"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-boot-starter/model"
)

func ParseConfig(config model.ApplicationConfig, rootKey string, res interface{}) (err error) {
	var data []byte
	cfg, existed := config[rootKey]
	if !existed || cfg == nil {
		return
	}
	m := make(map[string]interface{})
	for k, v := range cfg.(map[interface{}]interface{}) {
		m[k.(string)] = v
	}
	data, err = json.Marshal(m)
	if err != nil {
		return
	}
	return json.Unmarshal(data, res)
}
