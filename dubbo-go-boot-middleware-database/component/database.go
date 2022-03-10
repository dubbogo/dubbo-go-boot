package component

import (
	"gorm.io/gorm"
)

var (
	DatabaseComponent = &databaseComponent{}
)

type databaseComponent struct {
	Db *gorm.DB
}

func GetDatabase() *gorm.DB {
	return DatabaseComponent.Db
}
