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

import "fmt"

func main() {
	lru := NewLru(12)
	val := lru.Get("sachin")
	fmt.Println(val)
	fmt.Println(lru.size)
	lru.Set("sachin", "saini")
	val = lru.Get("sachin")
	fmt.Println(val)
	fmt.Println(lru.size)
	lru.Set("peeyush", "pessss")
	val = lru.Get("peeyush")
	fmt.Println(val)
	fmt.Println(lru.size)
	lru.Set("vrinda", "ves")
	val = lru.Get("vrinda")
	fmt.Println(val)
	fmt.Println(lru.size)
	val = lru.Get("sachin")
	fmt.Println(val)
	fmt.Println(lru.cache.Len())
}
