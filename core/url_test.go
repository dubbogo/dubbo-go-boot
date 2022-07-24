/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	parse, err := url.Parse("http://127.0.0.1:8500")

	assert.Nil(t, err)

	t.Log(parse.Port())
	t.Log(parse.Hostname())
}

func TestNewUrl(t *testing.T) {
	u, err := NewURL("http://127.0.0.1:8500",
		WithUsername("admin"),
		WithPassword("123456"),
		WithParamsValue("protocol", "consul"),
	)

	assert.Nil(t, err)
	assert.Equal(t, u.Ip, "127.0.0.1")
	assert.Equal(t, u.Scheme, "http")
	assert.Equal(t, u.Port, "8500")
	assert.Equal(t, u.Location, "127.0.0.1:8500")
	assert.Equal(t, u.Location, "127.0.0.1:8500")
	assert.Equal(t, u.GetParam("protocol", "abc"), "consul")
	assert.Equal(t, u.GetParam("scheme", ""), "http")

}
