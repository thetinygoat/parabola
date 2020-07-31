//Copyright 2020 Sachin Saini

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package set

import "github.com/emirpasic/gods/sets/hashset"

type Set struct {
	set    map[string]*hashset.Set
	memory int64
}

func New() *Set {
	return &Set{set: make(map[string]*hashset.Set), memory: 0}
}

func (s *Set) Add(key, value string) {
	_, found := s.set[key]
	if !found {
		s.set[key] = hashset.New()
	}
	s.allocate(value)
	s.set[key].Add(value)
}

func (s *Set) Contains(key, value string) bool {
	set, found := s.set[key]
	if !found {
		return false
	}

	return set.Contains(value)
}

func (s *Set) Remove(key, value string) {
	set, found := s.set[key]
	if !found {
		return
	}
	if set.Contains(value) {
		set.Remove(value)
		s.free(value)
	}
}

func (s *Set) Size(key string) int64 {
	set, found := s.set[key]
	if !found {
		return -1
	}

	return int64(set.Size())
}

func (s *Set) allocate(data string) {
	s.memory += int64(len(data))
}

func (s *Set) free(data string) {
	s.memory -= int64(len(data))
}
