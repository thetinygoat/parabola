// Copyright 2020 Sachin Saini

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package lru provides core lru cache functionality
package lru

import (
	"time"
)

type node struct {
	next, prev *node
	expireTime time.Time
	key, value string
}

// Lru provides core lru functionality
type Lru struct {
	head, tail              *node
	keymap                  map[string]*node
	size, maxMemory, memory int64
}

// New instantiates a new lru
func New(maxMemory int64) *Lru {
	l := &Lru{}
	l.size = 0
	l.maxMemory = maxMemory
	l.memory = 0
	l.head = &node{}
	l.tail = &node{}
	l.head.next = l.tail
	l.tail.prev = l.head
	l.keymap = make(map[string]*node)
	return l
}

// Set inserts a key-value pair in to the lru
func (lru *Lru) Set(key, value string, ttl int64) {
	if lru.Contains(key) {
		node := lru.keymap[key]
		oldValue := node.value
		node.value = value
		node.expireTime = time.Now().Add(time.Duration(ttl) * time.Second)
		lru.reorder(node)
		lru.free(oldValue)
		lru.allocate(node.value)
		for lru.Size() > 0 && lru.memory > lru.maxMemory {
			value := lru.head.next.value
			lru.remove(lru.head.next)
			lru.size--
			lru.free(value)
		}
	} else {
		newNode := &node{}
		newNode.value = value
		newNode.expireTime = time.Now().Add(time.Duration(ttl) * time.Second)
		newNode.key = key
		lru.insert(newNode)
		lru.keymap[key] = newNode
		lru.size++
		lru.allocate(newNode.value)
		for lru.Size() > 0 && lru.memory > lru.maxMemory {
			value := lru.head.next.value
			lru.remove(lru.head.next)
			lru.size--
			lru.free(value)
		}
	}
}

// Get returns the value associated with a key
func (lru *Lru) Get(key string) (string, bool) {
	if !lru.Contains(key) {
		return "", false
	}
	node := lru.keymap[key]

	if lru.isExpired(node.expireTime) {
		lru.free(node.value)
		lru.remove(node)
		delete(lru.keymap, node.key)
		return "", false
	}
	lru.reorder(node)

	return node.value, true
}

// Contains returns true/false depending on whether the key is present in the cache
func (lru *Lru) Contains(key string) bool {
	_, found := lru.keymap[key]
	return found
}

// Size returns the size of the lru
func (lru *Lru) Size() int64 {
	return lru.size
}

func (lru *Lru) isExpired(expireTime time.Time) bool {
	return time.Now().Sub(expireTime) > 0
}

func (lru *Lru) reorder(node *node) {
	lru.remove(node)
	lru.insert(node)
}

func (lru *Lru) remove(node *node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil
}

func (lru *Lru) insert(node *node) {
	prev := lru.tail.prev
	node.next = lru.tail
	node.prev = prev
	prev.next = node
	lru.tail.prev = node
}

func (lru *Lru) free(data string) {
	lru.memory -= int64(len(data))
}

func (lru *Lru) allocate(data string) {
	lru.memory += int64(len(data))
}
