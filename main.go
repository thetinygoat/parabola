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
	"fmt"
	"io"
	"net"

	"github.com/thetinygoat/DictX/protocol"
)

func main() {
	ln, err := net.Listen("tcp", ":9898")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	for {
		r := bufio.NewReader(conn)
		msg, err := protocol.Read(r)
		if err != nil && err != io.EOF {
			conn.Close()
			panic(err)
		}
		m := msg.Array()
		for i := range m {
			fmt.Println(m[i].String())
		}
	}
}
