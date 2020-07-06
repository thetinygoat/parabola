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

	"github.com/emirpasic/gods/maps/hashmap"
)

// HMap implements hashset for key, value storage
type HMap struct {
	hmap  *hashmap.Map
	mutex sync.Mutex
}

// NewHMap instantiates new hashmap
func NewHMap() *HMap {
	h := HMap{}
	h.hmap = hashmap.New()

	return &h
}

// Put inserts a new key, value pair
func (h *HMap) Put(key, value interface{}) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.hmap.Put(key, value)
	return nil
}

// Get retrieves value from the map
func (h *HMap) Get(key interface{}) (interface{}, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	value, ok := h.hmap.Get(key)
	if !ok {
		return nil, errors.New(HMapNoKeyError)
	}
	return value, nil
}

// Purge clears the map
func (h *HMap) Purge() error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.hmap.Clear()
	return nil
}
