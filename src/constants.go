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

// Misc constants
const (
	NoExist        = "nil"
	Ok             = "ok"
	MemoryOverLoad = "MemoryOverLoad"
	configFile     = "memdb.conf"
)

// Lru constants
const (
	LruMaxKeys       = 1000000 // max keys lru can hold
	LruInternalError = "LruInternalError"
	LruNoKeyError    = "lruNoKeyError"
)

// List constants
const (
	ListNokeyError    = "ListNokeyError"
	ListInternalError = "ListInternalError"
)

// HMap constants
const (
	HMapNoKeyError    = "HMapNoKeyError"
	HMapInternalError = "HMapInternalError"
)

// Port is the port memdb runs on
var Port string

// Password is the memdb password
var Password string

// EvictionPolicy is the cache eviction policy
var EvictionPolicy string

// OperationMode specifies mode of operation ie clustereor normal
var OperationMode string

// MaxMemory is the max memory available to memdb
var MaxMemory int
