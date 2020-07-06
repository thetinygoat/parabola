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
	"errors"
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

// Lru implements lru caching
type Lru struct {
	cache *lru.Cache
	mutex sync.Mutex
}

// NewLru instantiates new Lru cache
func NewLru() *Lru {
	l := Lru{}
	l.cache, _ = lru.New(LruMaxKeys)
	return &l
}

// Get handles getting the values from the cache
func (l *Lru) Get(key string) (interface{}, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	value, ok := l.cache.Get(key)
	if !ok {
		return nil, errors.New(LruNoKeyError)
	}
	return value, nil
}

// Set handles adding values to the cache
func (l *Lru) Set(key string, value string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.cache.Add(key, value)
	return nil
}

// Remove removes a key from the keyspace
func (l *Lru) Remove(key string) (interface{}, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if !l.cache.Contains(key) {
		return nil, errors.New(LruNoKeyError)
	}
	value, _ := l.cache.Get(key)
	l.cache.Remove(key)
	return value, nil
}

// Purge completley clears the cache
func (l *Lru) Purge() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.cache.Purge()
	return nil
}

// RemoveOldest removes oldest entry in the lru
func (l *Lru) RemoveOldest() (interface{}, interface{}, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	key, value, ok := l.cache.RemoveOldest()
	if !ok {
		return nil, nil, errors.New(LruInternalError)
	}
	return key, value, nil
}
