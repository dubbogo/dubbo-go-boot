package extension

import (
	"github.com/pkg/errors"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/database"
)

var databases = make(map[string]func(config *core.URL) (database.Database, error))

func GetDatabase(name string, config *core.URL) (database.Database, error) {
	if databases[name] == nil {
		return nil, errors.Errorf("database for %s driver does not exist. "+
			"please make sure that you have imported the package "+
			"github.com/dubbogo/dubbo-go-boot/database/%s",
			name, name)
	}
	return databases[name](config)
}

func SetDatabase(name string, driver func(config *core.URL) (database.Database, error)) {
	databases[name] = driver
}
