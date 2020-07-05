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
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

// Conn handles communication
type Conn struct {
	port    string
	clients map[net.Addr]bool
	mutex   sync.Mutex
}

// NewConnection instantiates a new connection
func NewConnection(port string) *Conn {
	c := Conn{}
	c.port = port
	c.clients = make(map[net.Addr]bool)
	return &c
}

// StartConnection starts the connection
func (c *Conn) StartConnection() {
	ln, err := net.Listen("tcp", ":"+c.port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		c.addClient(conn)
		fmt.Println(c.clients)
		go c.handleConnection(conn)
	}
}

func (c *Conn) handleConnection(conn net.Conn) {
	for {
		buf, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(buf) > 0 {
			conn.Write([]byte(buf))
		} else {
			c.removeClient(conn)
			fmt.Println(c.clients)
			conn.Close()
			return
		}
	}
}

func (c *Conn) addClient(conn net.Conn) {
	c.mutex.Lock()
	c.clients[conn.RemoteAddr()] = true
	c.mutex.Unlock()
}

func (c *Conn) removeClient(conn net.Conn) {
	c.mutex.Lock()
	delete(c.clients, conn.RemoteAddr())
	c.mutex.Unlock()
}
