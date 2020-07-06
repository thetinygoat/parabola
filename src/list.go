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
	"sync"

	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
)

// List implements list data structure
type List struct {
	list  *dll.List
	mutex sync.Mutex
}

// NewList instantiates a new list
func NewList() *List {
	l := List{}
	l.list = dll.New()

	return &l
}

// LPush adds items to the left of the list
func (l *List) LPush(data string) string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list.Append(data)
	return ok
}

// RPush adds items to the right of the list
func (l *List) RPush(data string) string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list.Prepend(data)
	return ok
}

// LPop removes items from the left of the list
func (l *List) LPop() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	size := l.list.Size()
	last, ok := l.list.Get(size - 1)
	if !ok {
		return noExist
	}
	l.list.Remove(size - 1)
	return fmt.Sprint(last)
}

// RPop removes items from the right of the list
func (l *List) RPop() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	first, ok := l.list.Get(0)
	if !ok {
		return noExist
	}
	l.list.Remove(0)
	return fmt.Sprint(first)
}

// LLen returns length of the list
func (l *List) LLen() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.list.Size()
}

// LPurge remove all items from the list
func (l *List) LPurge() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list.Clear()
	return ok
}
