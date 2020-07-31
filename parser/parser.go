//Copyright 2020 Sachin Saini

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//this package implements request parser

package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thetinygoat/dictX/dxep"
)

// Data types supported by dxep
const (
	Int = iota
	Str
	Array
	Err
	Nil
)
const (
	Set  = "SET"
	Hash = "HASH"
	List = "LIST"
	Key  = "KEY"
)

var (
	errValidate = errors.New("invalid query")
)

type ParsedQuery struct {
	cmd    string
	subcmd string
	data   []string
}

func Parse(query []*dxep.Message) {
	err := validate(query)
	if err != nil {
		return
	}
	tokens := make([]string, len(query))
	for i, token := range query {
		tokens[i], _ = token.Str()
	}
	cmd := strings.ToUpper(tokens[0])
	switch cmd {
	case Hash:
		fmt.Println("reached hash query parser")
		parseHashQuery(tokens)
	case Set:
		fmt.Println("reached set query parser")

	case List:
		fmt.Println("reached list query parser")

	case Key:
		fmt.Println("reached key query parser")
	default:
		fmt.Println("invalid command")

	}
}

//list of hash commands
//hash put key field value
//hash get key field
//hash del key field

func parseHashQuery(tokens []string) (string, []string, error) {
	if len(tokens) < 3 {
		return "", nil, errValidate
	}
	subcmd := tokens[1]
	switch subcmd {
	case "put":
		fmt.Println("reached hash put")
	case "del":
		fmt.Println("reached hash del")
	case "get":
		fmt.Println("reached hash get")
	default:
		return "", nil, errValidate
	}
	return "", nil, errValidate
}

func validate(query []*dxep.Message) error {
	if len(query) < 1 {
		return errValidate
	}
	for _, m := range query {
		if m.Type != Str {
			return errValidate
		}
	}
	return nil
}
