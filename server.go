// Copyright 2020 Sachin Saini

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/thetinygoat/dictX/dxep"
)

const (
	bufSize = 4096
)

// Server describes dictX server
type Server struct {
	listener net.Listener
	timeout  time.Duration
}

// InitServer initializes the Server struct
func InitServer(network, addr string, timeout time.Duration) (*Server, error) {
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	srv := &Server{}
	srv.listener = ln
	srv.timeout = timeout
	return srv, nil
}

// Listen listens for new connections
func (srv *Server) Listen() error {
	defer srv.listener.Close()

	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			return err
		}
		go srv.read(conn)
	}
}

func (srv *Server) read(conn net.Conn) {
	defer conn.Close()
	for {
		// conn.SetDeadline(time.Now().Add(time.Second * 30))
		buf := bufio.NewReaderSize(conn, bufSize)
		msg, err := dxep.Parse(buf)
		if err != nil {
			log.Fatal(err)
		}
		arr, _ := msg.Array()
		for _, m := range arr {
			fmt.Println(m.Str())
		}
	}
}
