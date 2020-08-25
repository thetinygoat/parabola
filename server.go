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

package main

import (
	"bufio"
	"io"
	"net"

	"github.com/thetinygoat/DictX/lru"

	"github.com/thetinygoat/DictX/parser"

	"github.com/thetinygoat/DictX/protocol"
)

// Server provides core server functionality
type Server struct {
	ln    net.Listener
	port  string
	cache *lru.Lru
}

// NewServer instantiates a new tcp server
func NewServer(port string, maxMemory int64) (*Server, error) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return &Server{ln: ln, port: port, cache: lru.New(maxMemory)}, nil
}

// Listen starts the server and listens for connections on the port
func (srv *Server) Listen() error {
	for {
		conn, err := srv.ln.Accept()
		if err != nil {
			return err
		}
		go srv.handleConnection(conn)
	}
}

func (srv *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		r := bufio.NewReader(conn)
		m, err := protocol.Read(r)
		if err != nil && err != io.EOF {
			conn.Write([]byte(err.Error()))
			return
		}
		qArr, err := m.Array()
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
		q, err := parser.Parse(qArr)
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
		switch q.Cmd {
		case "GET":
			res, ok := srv.cache.Get(q.Key)
			if !ok {
				conn.Write([]byte("nil\n"))
			} else {
				conn.Write([]byte(res + "\n"))
			}
		case "SET":
			srv.cache.Set(q.Key, q.Value, q.TTL)
			conn.Write([]byte("ok\n"))
		case "DEL":
			srv.cache.Del(q.Key)
			conn.Write([]byte("ok\n"))
		default:
			conn.Write([]byte("no cmd\n"))
		}
	}
}
