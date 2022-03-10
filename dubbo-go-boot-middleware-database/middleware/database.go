/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

import (
	"fmt"
	"log"
	"os"
	"time"
)

import (
	clogger "dubbo.apache.org/dubbo-go/v3/common/logger"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

import (
	"github.com/dubbogo/dubbo-go-boot-starter/extend"
	"github.com/dubbogo/dubbo-go-boot-starter/middleware"
	startModel "github.com/dubbogo/dubbo-go-boot-starter/model"
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/component"
	"github.com/dubbogo/dubbo-go-boot/dubbo-go-middleware-database/model"
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
