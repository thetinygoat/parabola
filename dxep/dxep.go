package dxep

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

// Data types supported by dxep
const (
	Int = iota
	Str
	Array
	Err
	Nil
)

var (
	delim    = []byte{'\r', '\n'}
	delimEnd = delim[len(delim)-1]
)

// Prefixes for data types
var (
	intPrefix   = []byte{'%'}
	strPrefix   = []byte{'$'}
	arrayPrefix = []byte{'#'}
	errPrefix   = []byte{'-'}
)

var (
	errParse    = errors.New("parse error")
	errBadType  = errors.New("wrong type")
	errBadStr   = errors.New("bad string")
	errStrLimit = errors.New("max string size is 512MB")
)

// Type represents data type
type Type int

// Message descibes a parsed message
type Message struct {
	Type
	val interface{}
	raw []byte
}

// Parse parses the incoming dxep query string
func Parse(reader io.Reader) (*Message, error) {
	r := bufio.NewReader(reader)
	return parseBufMessage(r)
}

func parseBufMessage(r *bufio.Reader) (*Message, error) {
	b, err := r.Peek(1)
	if err != nil {
		return nil, err
	}
	switch b[0] {
	case intPrefix[0]:
		return readInt(r)
	case strPrefix[0]:
		return readStr(r)
	case arrayPrefix[0]:
		return readArray(r)
	case errPrefix[0]:
		return readErr(r)
	default:
		return nil, nil
	}
}

func readInt(r *bufio.Reader) (*Message, error) {
	b, err := r.ReadBytes(delimEnd)
	if err != nil {
		return nil, err
	}
	i, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
	if err != nil {
		return nil, errParse
	}
	return &Message{Type: Int, val: i, raw: b}, nil
}

func readStr(r *bufio.Reader) (*Message, error) {
	b, err := r.ReadBytes(delimEnd)
	if err != nil {
		return nil, err
	}
	size, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
	if err != nil {
		return nil, errParse
	}
	if size < 0 {
		return &Message{Type: Nil, raw: b}, nil
	}
	if size > 512e6 {
		return nil, errStrLimit
	}
	total := make([]byte, size)
	b2 := total
	var n int

	for len(b2) > 0 {
		n, err = r.Read(b2)
		if err != nil {
			return nil, err
		}
		b2 = b2[n:]
	}

	trail := make([]byte, 2)
	for i := 0; i < 2; i++ {
		if c, err := r.ReadByte(); err != nil {
			return nil, err
		} else {
			trail[i] = c
		}
	}
	blens := len(b) + len(total)
	raw := make([]byte, blens+2)
	raw = append(raw, b...)
	raw = append(raw, total...)
	raw = append(raw, trail...)
	return &Message{Type: Str, val: total, raw: raw}, nil
}

func readArray(r *bufio.Reader) (*Message, error) {
	b, err := r.ReadBytes(delimEnd)
	if err != nil {
		return nil, err
	}
	size, err := strconv.ParseInt(string(b[1:len(b)-2]), 10, 64)
	if err != nil {
		return nil, errParse
	}
	arr := make([]*Message, size)
	for i := range arr {
		m, err := parseBufMessage(r)
		if err != nil {
			return nil, err
		}
		arr[i] = m
		b = append(b, m.raw...)
	}
	return &Message{Type: Array, val: arr, raw: b}, nil
}

func readErr(r *bufio.Reader) (*Message, error) {
	b, err := r.ReadBytes(delimEnd)
	if err != nil {
		return nil, err
	}
	return &Message{Type: Err, val: b[1 : len(b)-2], raw: b}, nil
}

// Bytes returns byte slice representation of val
func (m *Message) Bytes() ([]byte, error) {
	if b, ok := m.val.([]byte); ok {
		return b, nil
	}
	return nil, errBadType
}

// Str is a Convenience method around Bytes which converts the output to a
// string
func (m *Message) Str() (string, error) {
	b, err := m.Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Int returns an int64 representing the value of the Message. Only valid for
// Int messages
func (m *Message) Int() (int64, error) {
	if i, ok := m.val.(int64); ok {
		return i, nil
	}
	return 0, errBadType
}

// Err returns an error representing the value of the Message. Only valid for
// Err messages
func (m *Message) Err() (error, error) {
	if m.Type != Err {
		return nil, errBadType
	}
	s, err := m.Str()
	if err != nil {
		return nil, err
	}
	return errors.New(s), nil
}

// Array returns the Message slice encompassed by this Messsage, assuming the
// Message is of type Array
func (m *Message) Array() ([]*Message, error) {
	if a, ok := m.val.([]*Message); ok {
		return a, nil
	}
	return nil, errBadType
}
