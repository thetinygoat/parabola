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

import dll "github.com/emirpasic/gods/lists/doublylinkedlist"

// List contains data related to a particular list
type List struct {
	list    *dll.List
	memSize uint64
}

// ListSpace contains all the lists
type ListSpace struct {
	listSpace map[string]*List
	listCount uint64
}

// InitListSpace initializes the listspace
func InitListSpace() *ListSpace {
	ls := ListSpace{}
	ls.listCount = 0
	ls.listSpace = make(map[string]*List)
	return &ls
}

// Append appends data to the specified list
func (ls *ListSpace) Append(listName, data string) {
	_, exists := ls.listSpace[listName]
	if !exists {
		newList := List{list: dll.New(), memSize: 0}
		ls.listSpace[listName] = &newList
	}

	list, _ := ls.listSpace[listName]
	list.list.Append(data)

}
