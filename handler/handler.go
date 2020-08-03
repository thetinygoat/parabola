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

package handler

import (
	"errors"
	"fmt"

	"github.com/thetinygoat/dictX/hash"
	"github.com/thetinygoat/dictX/key"
	"github.com/thetinygoat/dictX/list"
	"github.com/thetinygoat/dictX/parser"
	"github.com/thetinygoat/dictX/set"
)

const (
	Set  = "SET"
	Hash = "Hash"
	Key  = "KEY"
	List = "LIST"
)

const (
	resOk = "Ok"
)

const (
	Str   = "$"
	Int   = "%"
	Array = "#"
)

var (
	errMemoryFull  = errors.New("memory limit exceeded")
	errInvalidKey  = errors.New("invalid key")
	errServerError = errors.New("server error")
)

type Handler struct {
	hash     *hash.Hash
	set      *set.Set
	key      *key.Key
	list     *list.List
	capacity int64
}

type Response struct {
	data string
	typ  string
	null bool
}

//New returns a handler
func New() *Handler {
	return &Handler{hash: hash.New(), set: set.New(), key: key.New(), list: list.New(), capacity: 1024}
}

func (h *Handler) Handle(p *parser.ParsedQuery) (*Response, error) {
	switch p.Cmd() {
	case Set:
		fmt.Println("set query handler")
		return h.handleSet(p)
	}
	return nil, nil
}

func (h *Handler) handleSet(p *parser.ParsedQuery) (*Response, error) {

	switch p.Subcmd() {
	case "add":
		if !h.canInsert() {
			return nil, errMemoryFull
		}
		data := p.Data()
		for _, v := range data {
			h.set.Add(p.Key(), v)
		}
		return &Response{data: resOk, typ: Str, null: false}, nil
	case "del":
		h.set.Remove(p.Key(), p.Data()[0])
		return &Response{data: resOk, typ: Str, null: false}, nil
	case "contains":
		f := h.set.Contains(p.Key(), p.Data()[0])
		if f {
			return &Response{data: "1", typ: Int, null: false}, nil
		}
		return &Response{data: "0", typ: Int, null: false}, nil
	default:
		return nil, errServerError
	}
}

func (h *Handler) canInsert() bool {
	if h.set.Memory()+h.hash.Memory()+h.list.Memory()+h.key.Memory() < h.capacity {
		return true
	}
	return false
}
