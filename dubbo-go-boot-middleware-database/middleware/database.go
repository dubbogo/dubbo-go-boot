package middleware

import (
	clogger "dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	startModel "github.com/dubbogo/dubbo-go-boot-starter/model"
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/model"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	db = &databaseComponent{}
)

func init() {
	middleware.RegisterMiddleware(db)
}

type databaseComponent struct {
}

func databaseLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
}

func (c *databaseComponent) Setup(config startModel.ApplicationConfig, hooks []extend.DubboGoMiddlewareSetupHook) error {
	dbConfig := &model.DatabaseConfig{}
	err := util.ParseConfig(config, "database", dbConfig)
	if err != nil {
		clogger.Warn(err)
		clogger.Warn("please add database config")
		return nil
	}

	if dbConfig == nil {
		clogger.Warn("please add database config")
		return nil
	}

	for _, v := range hooks {
		if vv, ok := v.(*DatabaseSetupHook); ok {
			vv.Hook()
		}
	}

	dialect := dbConfig.Dialect
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.Database
	username := dbConfig.Username
	password := dbConfig.Password
	var dialector gorm.Dialector
	switch dialect {
	case "mysql":
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, host, port, database))
		break
	case "postgres":
		dialector = postgres.Open(fmt.Sprintf("user=%s password=%s host=%s port=%d DB.name=%s sslmode=disable TimeZone=Asia/Shanghai",
			username, password, host, port, database))
		break
	case "sqlserver":
		dialector = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			username, password, host, port, database))
		break
	case "clickhouse":
		dialector = clickhouse.Open(fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s",
			host, port, database, username, password))
		break
	}
	if dialector == nil {
		return fmt.Errorf("无效的数据库配置")
	}
	instance, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: databaseLogger(),
	})
	if err != nil {
		return err
	}
	component.DatabaseComponent.Db = instance
	return nil
}

func (c *databaseComponent) IsAsync() bool {
	return false
}

func (c *databaseComponent) Shutdown() {
	component.DatabaseComponent = nil
	db = nil
}
