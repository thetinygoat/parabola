// Copyright 2020 Sachin Saini

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hash

import (
	"github.com/emirpasic/gods/maps/hashmap"
)

// Hash describes a hashmap
type Hash struct {
	hash   map[string]*hashmap.Map
	memory int64
}

// New instantiates a new hash
func New() *Hash {
	return &Hash{hash: make(map[string]*hashmap.Map), memory: 0}
}

// Get returns the value associated with a key and a bool to indicate if the value was found
func (h *Hash) Get(superKey, key string) (value string, found bool) {
	// get hash from the hash space
	hash, found := h.hash[superKey]
	if !found {
		return "", false
	}
	val, found := hash.Get(key)
	return str(val), found
}

// Put adds or updates a field in a specified hash
func (h *Hash) Put(superKey, key, value string) {
	hash, found := h.hash[superKey]
	if !found {
		h.hash[superKey] = hashmap.New()
	} else {
		val, found := hash.Get(key)
		if found {
			h.free(str(val))
			h.free(key)
		}
	}
	h.hash[superKey].Put(key, value)
	h.allocate(key)
	h.allocate(value)
}

// Remove removes a field from the specified hash
func (h *Hash) Remove(superKey, key string) {
	hash, found := h.hash[superKey]
	if !found {
		return
	}
	value, found := hash.Get(key)
	if found {
		h.free(key)
		h.free(str(value))
		hash.Remove(key)
	}
}

// Size returns the size of the specified hash
func (h *Hash) Size(superKey, key string) int64 {
	hash, found := h.hash[superKey]
	if !found {
		return -1
	}
	return int64(hash.Size())
}

// Empty returns if the specified hash is empty or not
func (h *Hash) Empty(superKey, key string) int {
	hash, found := h.hash[superKey]
	if !found {
		return -1
	}
	if hash.Empty() {
		return 1
	}
	return 0
}

// Memory returns total memory used by all the hashes
func (h *Hash) Memory() int64 {
	return h.memory
}

func (h *Hash) allocate(data string) {
	h.memory += int64(len(data))
}

func (h *Hash) free(data string) {
	h.memory -= int64(len(data))
}

func str(data interface{}) string {
	return data.(string)
}
