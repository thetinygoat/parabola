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

// Package protocol provides core encoding and decoding
// for the communication protocol
package protocol

import (
	"bufio"
	"io"
	"strconv"
)

// data types
const (
	String = iota
	Integer
	Array
	Error
	Nil
)

var (
	stringPrefix  = byte('$')
	integerPrefix = byte('%')
	arrayPrefix   = byte('#')
	errorPrefix   = byte('!')
	nilPrefix     = byte('-')
)

var (
	sep   = []byte{'\r', '\n'}
	delim = sep[len(sep)-1]
)

// Message is the struct that contains the decoded/encoded data
type Message struct {
	t int
	v interface{}
}

// Read reads from io.reader
func Read(reader io.Reader) (*Message, error) {
	r := bufio.NewReader(reader)
	return read(r)
}

func read(r *bufio.Reader) (*Message, error) {
	h, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	switch h {
	case stringPrefix:
		return readString(r)
	case arrayPrefix:
		return readArray(r)
	default:
		return nil, nil
	}
}

func readString(r *bufio.Reader) (*Message, error) {
	buf, err := r.ReadBytes(delim)
	if err != nil {
		return nil, err
	}
	size, err := strconv.ParseInt(string(buf[:len(buf)-2]), 10, 64)
	if size < 0 {
		return &Message{t: Nil}, nil
	}
	parsed := make([]byte, size)
	buf = parsed
	for len(buf) > 0 {
		n, err := r.Read(buf)
		if err != nil {
			return nil, err
		}
		buf = buf[n:]
	}
	for i := 0; i < 2; i++ {
		if _, err := r.ReadByte(); err != nil {
			return nil, err
		}
	}
	return &Message{t: String, v: parsed}, nil
}

func readArray(r *bufio.Reader) (*Message, error) {
	buf, err := r.ReadBytes(delim)
	if err != nil {
		return nil, err
	}
	size, err := strconv.ParseInt(string(buf[:len(buf)-2]), 10, 64)
	if size < 0 {
		return &Message{t: Nil}, nil
	}
	parsed := make([]*Message, size)
	for i := range parsed {
		m, err := read(r)
		if err != nil {
			return nil, err
		}
		parsed[i] = m
	}
	return &Message{t: Array, v: parsed}, nil
}

func readInteger(r *bufio.Reader) (*Message, error) {
	buf, err := r.ReadBytes(delim)
	if err != nil {
		return nil, err
	}
	num, err := strconv.ParseInt(string(buf[:len(buf)-2]), 10, 64)
	if err != nil {
		return nil, err
	}
	return &Message{t: Integer, v: num}, nil
}

func readError(r *bufio.Reader) (*Message, error) {
	buf, err := r.ReadBytes(delim)
	if err != nil {
		return nil, err
	}
	return &Message{t: Error, v: buf[:len(buf)-2]}, nil
}

// Bytes returns underlying string
func (m *Message) Bytes() []byte {
	return m.v.([]byte)
}

// String returns underlying string
func (m *Message) String() string {
	b := m.Bytes()
	return string(b)
}

// Integer returns underlying integer(int64)
func (m *Message) Integer() int64 {
	return m.v.(int64)
}

// Array returns underlying message slice
func (m *Message) Array() []*Message {
	return m.v.([]*Message)
}
