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
	key    string
	data   []string
}

func Parse(query []*dxep.Message) (*ParsedQuery, error) {
	err := validate(query)
	if err != nil {
		return nil, err
	}
	tokens := make([]string, len(query))
	for i, token := range query {
		tokens[i], _ = token.Str()
	}
	cmd := strings.ToUpper(tokens[0])
	switch cmd {
	case Hash:
		subcmd, key, data, err := parseHashQuery(tokens)
		if err != nil {
			return nil, err
		}
		return &ParsedQuery{cmd: cmd, subcmd: subcmd, key: key, data: data}, nil
	case Set:
		subcmd, key, data, err := parseSetQuery(tokens)
		if err != nil {
			return nil, err
		}
		return &ParsedQuery{cmd: cmd, subcmd: subcmd, key: key, data: data}, nil
	case List:
		subcmd, key, data, err := parseListQuery(tokens)
		if err != nil {
			return nil, err
		}
		return &ParsedQuery{cmd: cmd, subcmd: subcmd, key: key, data: data}, nil
	case Key:
		subcmd, key, data, err := parseKeyQuery(tokens)
		if err != nil {
			return nil, err
		}
		return &ParsedQuery{cmd: cmd, subcmd: subcmd, key: key, data: data}, nil
	default:
		return nil, errValidate

	}
}

//list of hash commands
//hash put key field value
//hash get key field
//hash del key field

func parseHashQuery(tokens []string) (string, string, []string, error) {
	if len(tokens) <= 3 {
		return "", "", nil, errValidate
	}
	subcmd := tokens[1]
	key := tokens[2]
	data := tokens[3:]
	switch subcmd {
	case "put":
		if len(data) >= 2 && len(data)%2 == 0 {
			return subcmd, key, data, nil
		}
		return "", "", nil, errValidate

	case "del":
		if len(data) > 1 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	case "get":
		if len(data) > 1 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	default:
		return "", "", nil, errValidate
	}
}

//list of set queries
//set add key value
//set contains key value
//set del key value
func parseSetQuery(tokens []string) (string, string, []string, error) {
	if len(tokens) < 4 {
		return "", "", nil, errValidate
	}
	subcmd := tokens[1]
	key := tokens[2]
	data := tokens[3:]
	switch subcmd {
	case "add":
		if len(data) <= 0 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	case "contains":
		if len(data) > 1 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	case "del":
		if len(data) > 1 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	default:
		return "", "", nil, errValidate
	}

}

//list of key queries
//key put key value ex
//key get key
//key del key
func parseKeyQuery(tokens []string) (string, string, []string, error) {
	if len(tokens) < 3 {
		return "", "", nil, errValidate
	}
	subcmd := tokens[1]
	key := tokens[2]
	data := tokens[3:]
	switch subcmd {
	case "put":
		if len(data) < 2 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	case "get", "del":
		if len(data) > 0 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	default:
		return "", "", nil, errValidate
	}
}

//list of list commands
//list fpush key value
//list rpush key value
//list fpop key
//list rpop key
func parseListQuery(tokens []string) (string, string, []string, error) {
	if len(tokens) < 3 {
		return "", "", nil, errValidate
	}
	subcmd := tokens[1]
	key := tokens[2]
	data := tokens[3:]
	switch subcmd {
	case "fpush", "rpush":
		if len(data) <= 0 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	case "fpop", "rpop":
		if len(data) > 0 {
			return "", "", nil, errValidate
		}
		return subcmd, key, data, nil
	default:
		return "", "", nil, errValidate
	}
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

func (p *ParsedQuery) Cmd() string {
	return p.cmd
}

func (p *ParsedQuery) Subcmd() string {
	return p.subcmd
}
func (p *ParsedQuery) Key() string {
	return p.key
}

func (p *ParsedQuery) Data() []string {
	return p.data
}
