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
type Consumer interface {
	Allocate(*Manager)
	DeAllocate(*Manager)
}

// Manager implements global memory managemnt functions
type Manager struct {
	maxCapacity   int
	currentlyUsed int
	mutex         sync.Mutex
}

// NewMemoryManager instantiates a new memory manager
func NewMemoryManager() *Manager {
	m := Manager{}
	m.maxCapacity = MaxMemory
	m.currentlyUsed = 0
	return &m
}

// Allocate allocates memory
func (m *Manager) Allocate(c Consumer) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	c.Allocate(m)
}

// DeAllocate allocates memory
func (m *Manager) DeAllocate(c Consumer) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	c.DeAllocate(m)
}
