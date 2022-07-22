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

package sqlserver

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"testing"
)

type Student struct {
	//*gorm.Model
	Name string
	ID   string
}
type v2Suite struct {
	db      *gorm.DB
	mock    sqlmock.Sqlmock
	student Student
}

func TestGORMV2(t *testing.T) {
	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}
	dialector := sqlserver.New(sqlserver.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "sqlserver",
		Conn:                 db,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
}
