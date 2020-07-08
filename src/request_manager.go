// Copyright (C) 2020 Sachin Saini

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"regexp"
	"strings"
)

// RequestManager holds handlers for different data structures
type RequestManager struct {
	lruHandler  *LruHandler
	listHandler *ListHandler
}

// NewRequestManager instantiates a new request manager
func NewRequestManager(memoryManager *MemoryManager) *RequestManager {
	m := RequestManager{}
	m.lruHandler = NewLruHandler(memoryManager)
	m.listHandler = NewListHandler(memoryManager)

	return &m
}

// ParseRequestString parses request string and passes the request to
// suitable handler
func (m *RequestManager) ParseRequestString(req string) string {
	req = req[:len(req)-1]
	listRegex := regexp.MustCompile("^L")
	lruRegex := regexp.MustCompile("^LRU")
	if lruRegex.MatchString(req) {
		return m.parseLruQuery(req)
	} else if listRegex.MatchString(req) {
		return "List handler reached"
	}
	return InvalidOperation
}

func (m *RequestManager) parseLruQuery(query string) string {
	q := strings.Split(query, " ")
	switch q[0] {
	case LruGet:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.lruHandler.LRUGet(q[1])
	case LruSet:
		if len(q) < 3 {
			return InvalidOperation
		}
		return m.lruHandler.LRUSet(q[1], q[2])
	case LruRem:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.lruHandler.LRURemove(q[1])
	case LruPurge:
		return m.lruHandler.LRUPurge()
	default:
		return InvalidOperation
	}
}
