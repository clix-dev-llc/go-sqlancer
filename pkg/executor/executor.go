// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package executor

import (
	"regexp"

	"github.com/juju/errors"

	"github.com/chaos-mesh/go-sqlancer/pkg/connection"
	"github.com/chaos-mesh/go-sqlancer/pkg/types"
)

var (
	dbnameRegex = regexp.MustCompile(`([a-z0-9A-Z_]+)$`)
)

// Executor define test executor
type Executor struct {
	conn   *connection.Connection
	db     string
	tables map[string]*types.Table
}

// New create Executor
func New(dsn string, db string) (*Executor, error) {
	conn, err := connection.New(dsn, &connection.Option{
		Log:        "",
		Mute:       false,
		GeneralLog: true,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	e := Executor{
		conn:   conn,
		db:     db,
		tables: make(map[string]*types.Table),
	}
	return &e, nil
}

// GetTables get table map
func (e *Executor) GetTables() map[string]*types.Table {
	return e.tables
}
