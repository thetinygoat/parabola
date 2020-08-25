// Copyright 2020 Sachin Saini

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package parser provides query parsing functionality
package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/thetinygoat/DictX/protocol"
)

const (
	setQuery = "SET"
	getQuery = "GET"
	delQuery = "DEL"
)

var (
	errInvalid = errors.New("invalid query")
)

// Query holds all the query data
type Query struct {
	Cmd, Key, Value string
	TTL             int64
}

// Parse parses the incoming query
func Parse(query []*protocol.Message) (*Query, error) {
	ok := validate(query)
	if !ok {
		return nil, errInvalid
	}
	tokens := make([]string, len(query))
	for i, token := range query {
		tokens[i], _ = token.String()
	}
	cmd := strings.ToUpper(tokens[0])

	switch cmd {
	case setQuery:
		return parseSetQuery(tokens)
	case getQuery:
		return parseGetQuery(tokens)
	case delQuery:
		return parseDelQuery(tokens)
	default:
		return nil, errInvalid
	}
}

func parseSetQuery(tokens []string) (*Query, error) {
	if len(tokens) < 4 {
		return nil, errInvalid
	}
	ttl, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, err
	}
	q := &Query{Cmd: strings.ToUpper(tokens[0]), Key: tokens[1], Value: tokens[2], TTL: ttl}
	return q, nil
}
func parseGetQuery(tokens []string) (*Query, error) {
	if len(tokens) < 2 {
		return nil, errInvalid
	}
	q := &Query{Cmd: strings.ToUpper(tokens[0]), Key: tokens[1]}
	return q, nil
}
func parseDelQuery(tokens []string) (*Query, error) {
	if len(tokens) < 2 {
		return nil, errInvalid
	}
	q := &Query{Cmd: strings.ToUpper(tokens[0]), Key: tokens[1]}
	return q, nil
}

func validate(query []*protocol.Message) bool {
	if len(query) < 2 {
		return false
	}
	return true
}
