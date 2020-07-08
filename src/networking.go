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

	"github.com/emirpasic/gods/sets/hashset"
)

// Conn handles communication
type Conn struct {
	port     string
	listener net.Listener
	clients  *hashset.Set
	mutex    sync.Mutex
}

// NewConnection instantiates a new connection
func NewConnection() *Conn {
	c := Conn{}
	c.port = Port
	c.clients = hashset.New()
	ln, err := net.Listen("tcp", ":"+c.port)
	if err != nil {
		panic(err)
	}
	c.listener = ln
	return &c
}

// StartConnection starts the connection
func (c *Conn) StartConnection() {

	for {
		conn, err := c.listener.Accept()
		defer c.listener.Close()
		if err != nil {
			panic(err)
		}
		c.addClient(conn)
		go c.handleConnection(conn)

	}
}

// StopConnection signals clients that it has been disconnected and removes them
func (c *Conn) StopConnection() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	conns := c.clients.Values()
	for _, conn := range conns {
		conn.(net.Conn).Write([]byte(Disconnected + "\n"))
		c.clients.Remove(conn)
	}
}

func (c *Conn) handleConnection(conn net.Conn) {
	for {
		defer conn.Close()
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
	defer c.mutex.Unlock()
	c.clients.Add(conn)
}

func (c *Conn) removeClient(conn net.Conn) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clients.Remove(conn)
}
