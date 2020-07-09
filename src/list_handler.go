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
)

type listContainer struct {
	list    *List
	memUsed int
}

// ListHandler implements methods for managing lists
type ListHandler struct {
	listspace map[string]*listContainer
	manager   *MemoryManager
}

// NewListHandler instantiates a new list
func NewListHandler(m *MemoryManager) *ListHandler {
	l := ListHandler{}
	l.listspace = make(map[string]*listContainer)
	l.manager = m
	return &l
}

// LPush inserts items to the left of the list
func (l *ListHandler) LPush(name, data string) string {
	if l.manager.CurrentlyUsed+len(data) >= l.manager.MaxCapacity {
		return MemoryLimitExceeded
	}
	list, ok := l.listspace[name]
	if !ok {
		list = &listContainer{list: NewList(), memUsed: 0}
		l.listspace[name] = list
	}
	err := list.list.LPush(data)
	if err != nil {
		return ListInternalError
	}
	list.memUsed += len(data)
	l.manager.Allocate(len(data))
	return Ok
}

// RPush inserts items to the right of the list
func (l *ListHandler) RPush(name, data string) string {
	if l.manager.CurrentlyUsed+len(data) >= l.manager.MaxCapacity {
		return MemoryLimitExceeded
	}
	list, ok := l.listspace[name]
	if !ok {
		list = &listContainer{list: NewList(), memUsed: 0}
		l.listspace[name] = list
	}
	err := list.list.RPush(data)
	if err != nil {
		return fmt.Sprint(err)
	}
	list.memUsed += len(data)
	l.manager.Allocate(len(data))
	return Ok
}

// LPop removes items from the left of the list
func (l *ListHandler) LPop(name string) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	dataRaw, err := list.list.LPop()
	if err != nil {
		return fmt.Sprint(err)
	}
	data := fmt.Sprint(dataRaw)
	list.memUsed -= len(data)
	l.manager.Free(len(data))
	return data
}

// RPop removes items from the right of the list
func (l *ListHandler) RPop(name string) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	dataRaw, err := list.list.RPop()
	if err != nil {
		return fmt.Sprint(err)
	}
	data := fmt.Sprint(dataRaw)
	list.memUsed -= len(data)
	l.manager.Free(len(data))
	return data
}

// LLen return length of the string
func (l *ListHandler) LLen(name string) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	lenRaw := list.list.Len()
	len := fmt.Sprint(lenRaw)
	return len
}

// LClean empties the list
func (l *ListHandler) LClean(name string) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	memused := list.memUsed
	list.list.Purge()
	list.memUsed = 0
	l.manager.Free(memused)

	return Ok
}

//LDel removes a list from the listspace
func (l *ListHandler) LDel(name string) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	memused := list.memUsed
	list.list = nil
	list.memUsed = 0
	l.manager.Free(memused)
	delete(l.listspace, name)
	return Ok
}

// LRemIdx removes data from a particular index
func (l *ListHandler) LRemIdx(name string, idx int) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	dataRaw, err := list.list.RemIdx(idx)
	if err != nil {
		return fmt.Sprint(err)
	}
	data := fmt.Sprint(dataRaw)
	list.memUsed -= len(data)
	l.manager.Free(len(data))
	return data
}

// LGetIdx returns data from a particular index
func (l *ListHandler) LGetIdx(name string, idx int) string {
	list, ok := l.listspace[name]
	if !ok {
		return NoExist
	}
	dataRaw, err := list.list.GetIdx(idx)
	if err != nil {
		return fmt.Sprint(err)
	}
	data := fmt.Sprint(dataRaw)
	return data
}
