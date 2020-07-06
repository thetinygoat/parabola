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
	"sync"

	"github.com/willf/bloom"
)

// BloomFilter implements a bloom filter
type BloomFilter struct {
	filter *bloom.BloomFilter
	mutex  sync.Mutex
}

// NewBloomFilter instantiates a new bloom filter
func NewBloomFilter(m uint, k uint) *BloomFilter {
	b := BloomFilter{}
	b.filter = bloom.New(m, k)
	return &b
}

// Add inserts a key into the filter
func (b *BloomFilter) Add(key string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.filter.Add([]byte(key))
	return nil
}

// Get checks if a key is present in the filter
func (b *BloomFilter) Get(key string) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.filter.Test([]byte(key))
}

// Purge removes all keys from the filter
func (b *BloomFilter) Purge() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.filter.ClearAll()
	return nil
}
