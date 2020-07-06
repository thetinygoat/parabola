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

// Consumer specifes methods to be used for allocation and deallocation of memory

// Manager implements global memory managemnt functions
type Manager struct {
	MaxCapacity   int
	CurrentlyUsed int
	mutex         sync.Mutex
}

// NewMemoryManager instantiates a new memory manager
func NewMemoryManager() *Manager {
	m := Manager{}
	m.MaxCapacity = MaxMemory
	m.CurrentlyUsed = 0
	return &m
}

// Allocate allocates memory
func (m *Manager) Allocate(chunk int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.CurrentlyUsed += chunk
}

// Free allocates memory
func (m *Manager) Free(chunk int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.CurrentlyUsed -= chunk
}
