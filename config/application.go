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

package config

import (
	"github.com/creasty/defaults"
)

func init() {
	SetConfig("application", &Application{})
}

type Application struct {
	Organization string `default:"dubbo-go-boot" yaml:"organization" json:"organization"`
	Name         string `default:"dubbo.io" yaml:"name" json:"name" property:"name"`
	Module       string `default:"sample" yaml:"module" json:"module"`
	Version      string `default:"1.0.0" yaml:"version" json:"version"`
	Owner        string `default:"dubbo-go" yaml:"owner" json:"owner"`
	Environment  string `default:"dev" yaml:"environment" json:"environment"`
}

// Prefix dubbo.application
func (Application) Prefix() string {
	return "application"
}

// Init  application config and set default value
func (a *Application) Init() error {
	var (
		err error
	)

	if err = defaults.Set(a); err != nil {
		return err
	}
	return nil
}

func (a *Application) Order() int {
	return 2
}
