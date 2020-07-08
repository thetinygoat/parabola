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

import "sync"

// MemoryManager implements global memory managemnt functions
type MemoryManager struct {
	MaxCapacity   int
	CurrentlyUsed int
	mutex         sync.Mutex
}

// NewMemoryManager instantiates a new memory manager
func NewMemoryManager() *MemoryManager {
	m := MemoryManager{}
	m.MaxCapacity = MaxMemory
	m.CurrentlyUsed = 0
	return &m
}

// Allocate allocates memory
func (m *MemoryManager) Allocate(chunk int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.CurrentlyUsed += chunk
}

// Free allocates memory
func (m *MemoryManager) Free(chunk int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.CurrentlyUsed -= chunk
}
