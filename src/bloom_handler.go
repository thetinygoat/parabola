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

// BloomHandler holds data for managing bloom filter
type BloomHandler struct {
	filter  *BloomFilter
	manager *MemoryManager
}

// NewBloomHandler instantiates an new bloom filter handler
func NewBloomHandler(manager *MemoryManager) *BloomHandler {
	b := BloomHandler{}
	b.filter = NewBloomFilter(BloomBits, BloomHashes)
	b.manager = manager

	return &b
}

// BFAdd inserts data in the filter
func (b *BloomHandler) BFAdd(key string) string {
	err := b.filter.Add(key)
	if err != nil {
		return "internalerror"
	}
	return Ok
}

// BFTest inserts data in the filter
func (b *BloomHandler) BFTest(key string) string {
	res := b.filter.Test(key)
	if res {
		return "True"
	}
	return "False"
}

// BFPurge removes all data from filter
func (b *BloomHandler) BFPurge() string {
	b.filter.Purge()
	return Ok
}
