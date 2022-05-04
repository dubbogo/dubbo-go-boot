package config

import (
	"fmt"

	"github.com/dubbogo/dubbo-go-boot/core"
	"github.com/dubbogo/dubbo-go-boot/core/constant"
)

type Database struct {
	// mongo、mysql
	Driver string `json:"driver"`

	// database url
	Url string `json:"url"`

	// database username
	Username string `yaml:"username"`

	// database password
	Password string `yaml:"password"`

	// database connect timeout
	Timeout string `default:"5s" json:"timeout"`
}

func (Database) Prefix() string {
	return "database"
}

func (d *Database) toURL() (*core.URL, error) {
	address := fmt.Sprintf("%s://%s", d.Driver, d.Url)
	return core.NewURL(address,
		core.WithParamsValue(constant.DatabaseTimeoutKey, d.Timeout),
		core.WithParamsValue(constant.DatabaseKey, d.Driver),
	)
}
