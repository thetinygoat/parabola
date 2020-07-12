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
	NoExist             = "Nil"
	Ok                  = "Ok"
	Disconnected        = "Disconnected"
	InvalidOperation    = "InvalidOperation"
	InvalidIdx          = "InvalidIdx"
	MemoryLimitExceeded = "MemoryLimitExceeded"
	configFile          = "dictX.conf"
)

// Lru constants
const (
	LruMaxKeys       = 1000 // max keys lru can hold
	LruInternalError = "LruInternalError"
	LruNoKeyError    = "LruNoKeyError"

	LruGet   = "LRUGET"
	LruSet   = "LRUSET"
	LruRem   = "LRUREM"
	LruPurge = "LRUPURGE"
)

// List constants
const (
	ListEmptyError      = "ListEmptyError"
	ListInternalError   = "ListInternalError"
	ListInvalidIdxError = "ListInvalidIdxError"

	LPush   = "LLPUSH"
	RPush   = "LRPUSH"
	LPop    = "LLPOP"
	RPop    = "LRPOP"
	LLen    = "LLEN"
	LClean  = "LCLEAN"
	LDel    = "LDEL"
	LGetIdx = "LGETIDX"
	LRemIdx = "LREMIDX"
)

// HMap constants
const (
	HMapNoKeyError    = "HMapNoKeyError"
	HMapInternalError = "HMapInternalError"
)

// Bloom filter constants
const (
	BFAdd   = "BFADD"
	BFTest  = "BFTEST"
	BFPurge = "BFPURGE"
)

// Port is the port dictX runs on
var Port string

// Password is the dictX password
var Password string

// EvictionPolicy is the cache eviction policy
var EvictionPolicy string

// OperationMode specifies mode of operation ie clustereor normal
var OperationMode string

// MaxMemory is the max memory available to dictX
var MaxMemory int

// BloomEnable tells if bloom filter functionality is enabled or not
var BloomEnable bool

// BloomBits is the number of bits in a bloom filter
var BloomBits uint

// BloomHashes is the no. of hashing functions
var BloomHashes uint

// Logo is dictX logo with info
const Logo = `
  ooooooooooooooooooooo
 ooooooooooooooooooooooo	DictX v%s %s
ooos   ooos   sooo   sooo	
ooos   ooos   sooo   sooo	Running in %s mode
ooos   ooos   sooo   sooo	Port: %s
ooooooooooooooooooooooooo
hhhhhhhhhhhhhhhhhhhhhhhhh	https://github.com/dictX/dictX
hhosyohoshhhdhhhsohoysohh
 h++s/y++hd   dh++y/s++h

`
