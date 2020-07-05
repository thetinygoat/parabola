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

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func parseMemory(memory string) int {
	kbRegex := regexp.MustCompile("kb$")
	mbRegex := regexp.MustCompile("mb$")
	gbRegex := regexp.MustCompile("gb$")
	var suffix int
	if memory[len(memory)-2] == 'k' {
		memory = kbRegex.ReplaceAllString(memory, "")
		suffix = 1024
	} else if memory[len(memory)-2] == 'm' {
		memory = mbRegex.ReplaceAllString(memory, "")
		suffix = 1024 * 1024
	} else if memory[len(memory)-2] == 'g' {
		memory = gbRegex.ReplaceAllString(memory, "")
		suffix = 1024 * 1024 * 1024
	}
	prefix, err := strconv.Atoi(memory)
	if err != nil {
		panic(err)
	}

	return prefix * suffix
}

// set global vars
func setVar(pair []string) {
	switch pair[0] {
	case "port":
		Port = pair[1]
	case "password":
		Password = pair[1]
	case "eviction-policy":
		EvictionPolicy = pair[1]
	case "operation-mode":
		OperationMode = pair[1]
	case "max-memory":
		MaxMemory = parseMemory(pair[1])
	}
}

// ReadConfig reads and parses config file
func ReadConfig() {
	path, _ := filepath.Abs(configFile)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	r.Comma = ' '
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		setVar(record)
	}
}
