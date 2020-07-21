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

func TestNewHash(t *testing.T) {
	h := NewHash()
	mem := h.mem
	if mem != 0 {
		t.Errorf("wrong mem, mem = %d, want 0", mem)
	}
}

func TestHashAllocate(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	h.Put(k1, v1)
	h.allocate(v1)
	if h.mem != 5 {
		t.Errorf("wrong mem, mem = %d, want 5", h.mem)
	}
	h.Put(k2, v2)
	h.allocate(v2)
	if h.mem != 11 {
		t.Errorf("wrong mem, mem = %d, want 11", h.mem)
	}
}

func TestHashFree(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	h.Put(k1, v1)
	h.allocate(v1)
	h.Put(k2, v2)
	h.allocate(v2)
	value, _ := h.Get(k1)
	h.Remove(k1)
	h.free(value)
	if h.mem != 6 {
		t.Errorf("wrong mem, mem = %d, want 6", h.mem)
	}
	value, _ = h.Get(k2)
	h.Remove(k2)
	h.free(value)
	if h.mem != 0 {
		t.Errorf("wrong mem, mem = %d, want 0", h.mem)
	}
}

func TestHashGet(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	k3 := ""
	v3 := ""
	h.Put(k1, v1)
	h.Put(k2, v2)
	h.Put(k3, v3)
	res1, _ := h.Get(k1)
	res2, _ := h.Get(k2)
	res3, _ := h.Get(k3)
	if res1 != v1 {
		t.Errorf("got wrong value, value = %s, want %s", res1, v1)
	}
	if res2 != v2 {
		t.Errorf("got wrong value, value = %s, want %s", res2, v2)
	}
	if res3 != v3 {
		t.Errorf("got wrong value, value = %s, want %s", res3, v3)
	}
}

func TestHashPut(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	k3 := ""
	v3 := ""
	h.Put(k1, v1)
	h.Put(k2, v2)
	h.Put(k3, v3)
	res1, _ := h.Get(k1)
	res2, _ := h.Get(k2)
	res3, _ := h.Get(k3)
	if res1 != v1 {
		t.Errorf("wrong value inserted, value = %s, want %s", res1, v1)
	}
	if res2 != v2 {
		t.Errorf("wrong value inserted, value = %s, want %s", res2, v2)
	}
	if res3 != v3 {
		t.Errorf("wrong value inserted, value = %s, want %s", res3, v3)
	}
}

func TestHashRemove(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	h.Put(k1, v1)
	h.Put(k2, v2)
	h.Remove(k1)
	ok := h.Contains(k1)
	if ok {
		t.Errorf("%s not removed, exists = %t, want false", k1, ok)
	}
	h.Remove(k2)
	ok = h.Contains(k2)
	if ok {
		t.Errorf("%s not removed, exists = %t, want false", k2, ok)
	}
	h.Remove(k1)
	ok = h.Contains(k1)
	if ok {
		t.Errorf("%s not removed, exists = %t, want false", k1, ok)
	}
}

func TestHashContains(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	h.Put(k1, v1)
	h.Put(k2, v2)
	ok := h.Contains(k1)
	if !ok {
		t.Errorf("%s is false negative, exists = %t, want true", k1, ok)
	}
	ok = h.Contains(k2)
	if !ok {
		t.Errorf("%s is false negative, exists = %t, want true", k2, ok)
	}
	ok = h.Contains("dictxserver")
	if ok {
		t.Errorf("%s is false positve, exists = %t, want false", "dictxserver", ok)
	}
}

func TestHashSize(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	h.Put(k1, v1)
	h.Put(k2, v2)
	if h.Size() != 2 {
		t.Errorf("wrong size, size = %d, want 2", h.Size())
	}
	h.Remove(k1)
	if h.Size() != 1 {
		t.Errorf("wrong size, size = %d, want 1", h.Size())
	}
	h.Remove(k2)
	if h.Size() != 0 {
		t.Errorf("wrong size, size = %d, want 0", h.Size())
	}
}

func TestHashValues(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	v := h.Values()
	if len(v) != 0 {
		t.Errorf("wrong size, size = %d, want 0", h.Size())
	}
	h.Put(k1, v1)
	h.Put(k2, v2)
	v = h.Values()
	if len(v) != 2 {
		t.Errorf("wrong size, size = %d, want 0", h.Size())
	}
}

func TestHashKeys(t *testing.T) {
	h := NewHash()
	k1 := "sachin"
	v1 := "saini"
	k2 := "dictX"
	v2 := "server"
	v := h.Keys()
	if len(v) != 0 {
		t.Errorf("wrong size, size = %d, want 0", h.Size())
	}
	h.Put(k1, v1)
	h.Put(k2, v2)
	v = h.Keys()
	if len(v) != 2 {
		t.Errorf("wrong size, size = %d, want 0", h.Size())
	}
}
