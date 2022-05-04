package extension

import (
	"github.com/pkg/errors"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/logger"
)

var logs = make(map[string]func(conf *core.URL) (logger.Logger, error))

func SetLogger(name string, v func(conf *core.URL) (logger.Logger, error)) {
	logs[name] = v
}

func GetLogger(name string, conf *core.URL) (logger.Logger, error) {
	if logs[name] == nil {
		return nil, errors.Errorf("logger for %s does not exist. "+
			"please make sure that you have imported the package "+
			"github.com/dubbogo/dubbo-go-boot/logger/%s", name, name)
	}
	return logs[name](conf)
}
