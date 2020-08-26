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
	"errors"
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

var (
	errParse       = errors.New("error parsing query")
	errInvalidType = errors.New("invalid type")
	errEncode      = errors.New("encoding error")
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
	if err != nil && err != io.EOF {
		return nil, err
	}

	switch h {
	case stringPrefix:
		return readString(r)
	case arrayPrefix:
		return readArray(r)
	case integerPrefix:
		return readInteger(r)
	case nilPrefix:
		return readString(r)
	case errorPrefix:
		return readError(r)
	default:
		return nil, errParse
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

// Type returns type of the data
func (m *Message) Type() int {
	return m.t
}

// Bytes returns underlying string
func (m *Message) Bytes() ([]byte, error) {
	if v, ok := m.v.([]byte); ok {
		return v, nil
	}
	return nil, errInvalidType
}

// String returns underlying string
func (m *Message) String() (string, error) {
	b, err := m.Bytes()
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

// Integer returns underlying integer(int64)
func (m *Message) Integer() (int64, error) {
	if v, ok := m.v.(int64); ok {
		return v, nil
	}
	return 0, errInvalidType
}

// Array returns underlying message slice
func (m *Message) Array() ([]*Message, error) {
	if v, ok := m.v.([]*Message); ok {
		return v, nil
	}
	return nil, errInvalidType
}

// EncodeStr encodes the given data to a string
func EncodeStr(data string) []byte {
	var buf []byte
	buf = append(buf, stringPrefix)
	buf = strconv.AppendInt(buf, int64(len(data)), 10)
	buf = append(buf, sep...)
	buf = append(buf, data...)
	buf = append(buf, sep...)

	return buf
}

// EncodeErr encodes the given data to an error
func EncodeErr(data string) []byte {
	var buf []byte
	buf = append(buf, errorPrefix)
	buf = append(buf, data...)
	buf = append(buf, sep...)

	return buf
}

// EncodeInt encodes the given data to an int
func EncodeInt(data int) []byte {
	var buf []byte
	buf = append(buf, integerPrefix)
	buf = strconv.AppendInt(buf, int64(data), 10)
	buf = append(buf, sep...)

	return buf
}

// EncodeNil encodes data to nil
func EncodeNil() []byte {
	var buf []byte
	buf = append(buf, nilPrefix)
	buf = strconv.AppendInt(buf, int64(-1), 10)
	buf = append(buf, sep...)

	return buf
}

// EncodeArray encodes data to an array
func EncodeArray(data []string) []byte {
	var buf []byte
	buf = append(buf, arrayPrefix)
	buf = strconv.AppendInt(buf, int64(len(data)), 10)
	buf = append(buf, sep...)
	for i := range data {
		buf = append(buf, EncodeStr(data[i])...)
	}

	return buf
}
