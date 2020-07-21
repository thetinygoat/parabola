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

import "testing"

func TestNewList(t *testing.T) {
	list := NewList()
	if list.mem != 0 {
		t.Errorf("mem = %d, want 0", list.mem)
	}
	if list.Length() > 0 {
		t.Errorf("length = %d, want 0", list.Length())
	}
}

func TestListAllocate(t *testing.T) {
	list := NewList()
	d1 := "sachin"
	list.Append(d1)
	allocSize := list.allocate(d1)
	if allocSize != 6 {
		t.Errorf("wrong allocation, allocSize = %d, want 6", allocSize)
	}
	d2 := "sachin2"
	list.Append(d2)
	allocSize = list.allocate(d2)
	if allocSize != 7 {
		t.Errorf("wrong allocation, allocSize = %d, want 7", allocSize)
	}
	if list.mem != 13 {
		t.Errorf("wrong mem footprint, mem = %d, want 13", list.mem)
	}
}

func TestListFree(t *testing.T) {
	list := NewList()
	d1 := "sachin"
	d2 := "saini"
	list.Append(d1)
	list.allocate(d1)
	list.Append(d2)
	list.allocate(d2)
	data, _ := list.RemoveLast()
	deallocSize := list.free(data)
	if deallocSize != 5 {
		t.Errorf("wrong deallocation, deallocSize = %d, want 5", deallocSize)
	}
	data, _ = list.RemoveLast()
	deallocSize = list.free(data)
	if deallocSize != 6 {
		t.Errorf("wrong deallocation, deallocSize = %d, want 6", deallocSize)
	}
	if list.mem != 0 {
		t.Errorf("wrong mem footprint, mem = %d, want 0", list.mem)
	}
}

func TestAppend(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Append(d1)
	list.Append(d2)
	list.Append(d3)
	res1, _ := list.Get(0)
	res2, _ := list.Get(1)
	res3, _ := list.Get(2)
	if res1 != d1 {
		t.Errorf("wrong data appended, data = %s, want %s", res1, d1)
	}
	if res2 != d2 {
		t.Errorf("wrong data appended, data = %s, want %s", res2, d2)
	}
	if res3 != d3 {
		t.Errorf("wrong data appended, data = %s, want %s", res3, d3)
	}
}

func TestPrepend(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Prepend(d1)
	list.Prepend(d2)
	list.Prepend(d3)
	res1, _ := list.Get(0)
	res2, _ := list.Get(1)
	res3, _ := list.Get(2)
	if res1 != d3 {
		t.Errorf("wrong data prepended, data = %s, want %s", res1, d3)
	}
	if res2 != d2 {
		t.Errorf("wrong data prepended, data = %s, want %s", res2, d2)
	}
	if res3 != d1 {
		t.Errorf("wrong data prepended, data = %s, want %s", res3, d1)
	}
}

func TestRemoveLast(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Append(d1)
	list.Append(d2)
	list.Append(d3)
	data, _ := list.RemoveLast()
	if data != d3 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d3)
	}
	data, _ = list.RemoveLast()
	if data != d2 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d2)
	}
	data, _ = list.RemoveLast()
	if data != d1 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d1)
	}
	if _, ok := list.RemoveLast(); ok {
		t.Errorf("wrong error, error = %t, want false", ok)
	}
}

func TestRemoveFirst(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Append(d1)
	list.Append(d2)
	list.Append(d3)
	data, _ := list.RemoveFirst()
	if data != d1 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d1)
	}
	data, _ = list.RemoveFirst()
	if data != d2 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d2)
	}
	data, _ = list.RemoveFirst()
	if data != d3 {
		t.Errorf("wrong data removed, data = %s, want %s", data, d3)
	}
	if _, ok := list.RemoveFirst(); ok {
		t.Errorf("wrong error, error = %t, want false", ok)
	}
}

func TestLength(t *testing.T) {
	list := NewList()
	list.Append("sachin")
	list.Append("saini")
	size := list.Length()
	if size != 2 {
		t.Errorf("wrong length, length = %d, want 2", size)
	}
	list.RemoveLast()
	size = list.Length()
	if size != 1 {
		t.Errorf("wrong length, length = %d, want 1", size)
	}
	list.RemoveLast()
	size = list.Length()
	if size != 0 {
		t.Errorf("wrong length, length = %d, want 0", size)
	}
}

func TestGet(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Append(d1)
	list.Append(d2)
	list.Append(d3)
	res1, _ := list.Get(0)
	res2, _ := list.Get(1)
	res3, _ := list.Get(2)
	if res1 != d1 {
		t.Errorf("got wrong data, data = %s, want %s", res1, d1)
	}
	if res2 != d2 {
		t.Errorf("got wrong data, data = %s, want %s", res2, d2)
	}
	if res3 != d3 {
		t.Errorf("got wrong data, data = %s, want %s", res3, d3)
	}
	if _, ok := list.Get(9999); ok {
		t.Errorf("wrong error, error = %t, want false", ok)
	}
}

func TestRemove(t *testing.T) {
	list := NewList()
	d1 := "Sachin"
	d2 := "saini"
	d3 := "dictX"
	list.Append(d1)
	list.Append(d2)
	list.Append(d3)
	res1, _ := list.Remove(0)
	if res1 != d1 {
		t.Errorf("got wrong data, data = %s, want %s", res1, d1)
	}
	res2, _ := list.Remove(1)
	res3, _ := list.Remove(0)
	if res2 != d3 {
		t.Errorf("got wrong data, data = %s, want %s", res2, d2)
	}
	if res3 != d2 {
		t.Errorf("got wrong data, data = %s, want %s", res3, d3)
	}
	if _, ok := list.Remove(9999); ok {
		t.Errorf("wrong error, error = %t, want false", ok)
	}
}
