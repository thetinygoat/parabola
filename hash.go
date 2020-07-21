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

// Hash contains data related to hashes
// hash - core hashmap ds
// mem - memory footprint of the map
type Hash struct {
	hash map[string]string
	mem  uint64
}

// NewHash instantiates a new hash
func NewHash() *Hash {
	h := Hash{hash: make(map[string]string), mem: 0}
	return &h
}

func (h *Hash) allocate(data string) uint64 {
	allocSize := uint64(len(data))
	h.mem += allocSize
	return allocSize
}

// free frees the virtual memory
func (h *Hash) free(data string) uint64 {
	if h.mem == 0 {
		return 0
	}
	deallocSize := uint64(len(data))
	h.mem -= deallocSize
	return deallocSize
}

// Put inserts a key-value pair
func (h *Hash) Put(key, value string) bool {
	h.hash[key] = value
	return true
}

// Get returns value for a specific key
func (h *Hash) Get(key string) (string, bool) {
	value, exists := h.hash[key]
	return value, exists
}

// Remove removes a key from the hash
func (h *Hash) Remove(key string) bool {
	delete(h.hash, key)
	return true
}

// Contains determines if a key exists
func (h *Hash) Contains(key string) bool {
	_, exists := h.hash[key]
	return exists
}

// Size returns size of the hash
func (h *Hash) Size() int {
	return len(h.hash)
}

// Values returns all the values in the hash
func (h *Hash) Values() []string {
	values := make([]string, h.Size())
	count := 0
	for _, value := range h.hash {
		values[count] = value
		count++
	}
	return values
}

// Keys returns all the keys in the hash
func (h *Hash) Keys() []string {
	keys := make([]string, h.Size())
	count := 0
	for key := range h.hash {
		keys[count] = key
		count++
	}
	return keys
}
