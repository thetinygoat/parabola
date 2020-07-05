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
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

const (
	maxKeys = 1000000 // max keys lru can hold
)

// Lru implements lru caching
type Lru struct {
	cache    *lru.Cache
	size     int
	capacity int
}

// NewLru instantiates new Lru cache
func NewLru(capacity int) *Lru {
	l := Lru{}
	l.cache, _ = lru.New(maxKeys)
	l.capacity = capacity
	l.size = 0
	return &l
}

// Get handles getting the values from the cache
func (l *Lru) Get(key string) string {
	if ok := l.cache.Contains(key); !ok {
		return noExist
	}
	value, ok := l.cache.Get(key)
	if !ok {
		return noExist
	}
	return fmt.Sprint(value)
}

// Set handles adding values to the cache
func (l *Lru) Set(key string, value string) string {
	// remove oldest until we have enough room for new value
	for l.size+len(value) > l.capacity {
		_, value, _ := l.cache.GetOldest()
		sizeOfOldest := len(fmt.Sprint(value))
		l.cache.RemoveOldest()
		l.size -= sizeOfOldest
	}
	l.cache.Add(key, value)
	l.size += len(value)
	return ok
}

// Remove removes a key from the keyspace
func (l *Lru) Remove(key string) string {
	if ok := l.cache.Contains(key); !ok {
		return noExist
	}
	value := l.Get(key)
	l.cache.Remove(key)
	l.size -= len(value)
	return ok
}

// Purge completley clears the cache
func (l *Lru) Purge() string {
	l.cache.Purge()
	return ok
}
