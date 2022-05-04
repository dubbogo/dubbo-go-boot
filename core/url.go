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

package core

import (
	"net/url"
	"strings"
	"sync"
	"time"
)

type URL struct {
	lock sync.RWMutex

	params url.Values

	Path string

	Username string

	Password string

	// ip+port
	Location string

	Ip string

	Port string

	Scheme string
}

type optionFunc func(*URL)

func (fn optionFunc) apply(vc *URL) {
	fn(vc)
}

type Option interface {
	apply(url *URL)
}

func WithUsername(username string) Option {
	return optionFunc(func(url *URL) {
		url.Username = username
	})
}

func WithPassword(pwd string) Option {
	return optionFunc(func(url *URL) {
		url.Password = pwd
	})
}

func WithPath(path string) Option {
	return optionFunc(func(url *URL) {
		url.Path = path
	})

}

func WithLocation(location string) Option {
	return optionFunc(func(url *URL) {
		url.Location = location
	})
}

func WithIp(ip string) Option {
	return optionFunc(func(url *URL) {
		url.Ip = ip
	})
}

func WithPort(port string) Option {
	return optionFunc(func(url *URL) {
		url.Port = port
	})
}

func WithScheme(scheme string) Option {
	return optionFunc(func(url *URL) {
		url.Scheme = scheme
	})
}

func WithParams(params url.Values) Option {
	return optionFunc(func(url *URL) {
		url.params = params
	})
}

func WithParamsValue(key, val string) Option {
	return optionFunc(func(url *URL) {
		url.SetParam(key, val)
	})
}

func NewURL(urlString string, opts ...Option) (u *URL, err error) {
	var (
		parse *url.URL
	)
	if urlString == "" {
		return nil, err
	}
	if parse, err = url.Parse(urlString); err != nil {
		return nil, err
	}
	u = &URL{}
	u.Scheme = parse.Scheme
	u.Location = parse.Host
	u.Port = parse.Port()
	u.Ip = parse.Hostname()

	for _, opt := range opts {
		opt.apply(u)
	}
	return u, nil
}

// GetParam gets value by key
func (c *URL) getParam(s string, d string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	var r string
	if len(c.params) > 0 {
		r = c.params.Get(s)
	}
	if len(r) == 0 {
		r = d
	}
	return r
}

func (c *URL) GetParam(key, d string) string {
	switch strings.ToLower(key) {
	case "scheme":
		return c.Scheme
	case "username":
		return c.Username
	case "Location", "host":
		return c.Location
	case "password":
		return c.Password
	case "port":
		return c.Port
	case "path":
		return c.Path
	default:
		return c.getParam(key, d)
	}
}

func (c *URL) SetParam(key string, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.params == nil {
		c.params = url.Values{}
	}
	c.params.Set(key, value)
}

func (c *URL) GetParamDuration(s string, d string) time.Duration {
	if t, err := time.ParseDuration(c.GetParam(s, d)); err == nil {
		return t
	}
	return 5 * time.Second
}
