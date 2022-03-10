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

package component

import (
	"github.com/dubbogo/dubbo-go-boot-starter/util"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	RedisComponent = &redisComponent{}
)

type redisComponent struct {
	sync.Mutex

	Addr       string
	Password   string
	MaxRetries int

	Redis *redis.Client

	DbMap map[int]*redis.Client
}

func GetRedis() *redis.Client {
	return RedisComponent.Redis
}

func GetRedisByIndex(dbIndex int) *redis.Client {
	RedisComponent.Lock()
	if client, notNil := RedisComponent.DbMap[dbIndex]; !notNil || client == nil {
		client = util.NewRedisDb(RedisComponent.Addr, RedisComponent.Password, dbIndex, RedisComponent.MaxRetries)
		RedisComponent.DbMap[dbIndex] = client
	}
	RedisComponent.Unlock()
	return RedisComponent.DbMap[dbIndex]
}
