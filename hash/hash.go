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
	hash   map[interface{}]*hashmap.Map
	memory int64
}

// New instantiates a new hash
func New() *Hash {
	return &Hash{hash: make(map[interface{}]*hashmap.Map), memory: 0}
}

// Get returns the value associated with a key and a bool to indicate if the value was found
func (h *Hash) Get(superKey, key interface{}) (value interface{}, found bool) {
	// get hash from the hash space
	hash, found := h.hash[superKey]
	if !found {
		return nil, false
	}
	return hash.Get(key)
}

// Put adds or updates a field in a specified hash
func (h *Hash) Put(superKey, key, value interface{}) {
	hash, found := h.hash[superKey]
	if !found {
		h.hash[superKey] = hashmap.New()
	} else {
		value, found := hash.Get(key)
		if found {
			h.free(value)
			h.free(key)
		}
	}
	h.hash[superKey].Put(key, value)
	h.allocate(key)
	h.allocate(value)
}

// Remove removes a field from the specified hash
func (h *Hash) Remove(superKey, key interface{}) {
	hash, found := h.hash[superKey]
	if !found {
		return
	}
	value, found := hash.Get(key)
	if found {
		h.free(key)
		h.free(value)
		hash.Remove(key)
	}
}

// Size returns the size of the specified hash
func (h *Hash) Size(superKey, key interface{}) int64 {
	hash, found := h.hash[superKey]
	if !found {
		return -1
	}
	return int64(hash.Size())
}

// Empty returns if the specified hash is empty or not
func (h *Hash) Empty(superKey, key interface{}) int {
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

func (h *Hash) allocate(data interface{}) {
	h.memory += int64(len(str(data)))
}

func (h *Hash) free(data interface{}) {
	h.memory -= int64(len(str(data)))
}

func str(data interface{}) string {
	return data.(string)
}
