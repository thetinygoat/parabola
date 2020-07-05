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

	"github.com/willf/bloom"
)

// BloomFilter implements a bloom filter
type BloomFilter struct {
	maxSize uint
	filter  *bloom.BloomFilter
}

// NewBloomFilter instantiates a new bloom filter
func NewBloomFilter(m uint, k uint) *BloomFilter {
	b := BloomFilter{}
	b.maxSize = m
	b.filter = bloom.New(m, k)
	return &b
}

// Add inserts a key into the filter
func (b *BloomFilter) Add(key string) string {
	b.filter.Add([]byte(key))
	return ok
}

// Get checks if a key is present in the filter
func (b *BloomFilter) Get(key string) string {
	return fmt.Sprint(b.filter.Test([]byte(key)))
}

// Purge removes all keys from the filter
func (b *BloomFilter) Purge() string {
	b.filter.ClearAll()
	return ok
}
