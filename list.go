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
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
)

const ()

// List contains data related to a list
// list - is a pointer the the list data structre
// mem - is the memory footprint of the list
type List struct {
	list *dll.List
	mem  uint64
}

// NewList instantiates a new list
func NewList() *List {
	l := List{list: dll.New(), mem: 0}
	return &l
}

// allocate allocates virtual memory from the mempool
func (l *List) allocate(data string) uint64 {
	allocSize := uint64(len(data))
	l.mem += allocSize
	return allocSize
}

// free frees the virtual memory
func (l *List) free(data string) uint64 {
	if l.mem == 0 {
		return 0
	}
	deallocSize := uint64(len(data))
	l.mem -= deallocSize
	return deallocSize
}

// Append appends data to the list
func (l *List) Append(data string) bool {
	l.list.Append(data)
	return true
}

// Prepend prepends data to the list
func (l *List) Prepend(data string) bool {
	l.list.Prepend(data)
	return true
}

// RemoveLast removes last element from the list
func (l *List) RemoveLast() (string, bool) {
	size := l.list.Size()
	data, exists := l.list.Get(size - 1)
	if !exists {
		return "", false
	}
	l.list.Remove(size - 1)
	return data.(string), true
}

// RemoveFirst removes first element from the list
func (l *List) RemoveFirst() (string, bool) {
	data, exists := l.list.Get(0)
	if !exists {
		return "", false
	}
	l.list.Remove(0)
	return data.(string), true
}

// Length returns the length of the list
func (l *List) Length() int {
	return l.list.Size()
}

// Get returns element at any index
func (l *List) Get(index int) (string, bool) {
	data, exists := l.list.Get(index)
	if !exists {
		return "", false
	}
	return data.(string), true
}

// Remove removes element from any index
func (l *List) Remove(index int) (string, bool) {
	data, exists := l.list.Get(index)
	if !exists {
		return "", false
	}
	l.list.Remove(index)
	return data.(string), true
}
