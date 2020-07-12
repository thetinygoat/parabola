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
	"strconv"
	"strings"
)

// RequestManager holds handlers for different data structures
type RequestManager struct {
	lruHandler   *LruHandler
	listHandler  *ListHandler
	bloomHandler *BloomHandler
}

// NewRequestManager instantiates a new request manager
func NewRequestManager(memoryManager *MemoryManager) *RequestManager {
	m := RequestManager{}
	m.lruHandler = NewLruHandler(memoryManager)
	m.listHandler = NewListHandler(memoryManager)
	m.bloomHandler = NewBloomHandler(memoryManager)
	return &m
}

// ParseRequestString parses request string and passes the request to
// suitable handler
func (m *RequestManager) ParseRequestString(req string) string {
	if len(req) <= 0 {
		return InvalidOperation
	}
	req = req[:len(req)-1]
	listRegex := regexp.MustCompile("^L")
	lruRegex := regexp.MustCompile("^LRU")
	bloomRegex := regexp.MustCompile("^BF")
	if lruRegex.MatchString(req) {
		return m.parseLruQuery(req)
	} else if listRegex.MatchString(req) {
		return m.parseListQuery(req)
	} else if bloomRegex.MatchString(req) {
		return m.parseBFQuery(req)
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

func (m *RequestManager) parseListQuery(query string) string {
	q := strings.Split(query, " ")
	switch q[0] {
	case LPush:
		if len(q) < 3 {
			return InvalidOperation
		}
		return m.listHandler.LPush(q[1], q[2])
	case RPush:
		if len(q) < 3 {
			return InvalidOperation
		}
		return m.listHandler.RPush(q[1], q[2])
	case LPop:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.listHandler.LPop(q[1])
	case RPop:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.listHandler.RPop(q[1])
	case LClean:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.listHandler.LClean(q[1])
	case LDel:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.listHandler.LDel(q[1])
	case LGetIdx:
		if len(q) < 3 {
			return InvalidOperation
		}
		idx, err := strconv.Atoi(q[2])
		if err != nil {
			return InvalidIdx
		}
		return m.listHandler.LGetIdx(q[1], idx)
	case LRemIdx:
		if len(q) < 3 {
			return InvalidOperation
		}
		idx, err := strconv.Atoi(q[2])
		if err != nil {
			return InvalidIdx
		}
		return m.listHandler.LRemIdx(q[1], idx)
	case LLen:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.listHandler.LLen(q[1])
	default:
		return InvalidOperation
	}
}

func (m *RequestManager) parseBFQuery(query string) string {
	q := strings.Split(query, " ")
	switch q[0] {
	case BFAdd:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.bloomHandler.BFAdd(q[1])
	case BFTest:
		if len(q) < 2 {
			return InvalidOperation
		}
		return m.bloomHandler.BFTest(q[1])
	case BFPurge:
		return m.bloomHandler.BFPurge()
	default:
		return InvalidOperation
	}
}
