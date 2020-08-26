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
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/thetinygoat/DictX/protocol"
)

func encode(query string) []byte {
	query = query[:len(query)-1]
	tokens := strings.Split(query, " ")
	return protocol.EncodeArray(tokens)
}

func main() {
	port := flag.String("port", "9898", "Port on which DictX server is running")
	flag.Parse()

	conn, err := net.Dial("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		query, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		enc := encode(query)
		conn.Write(enc)
		r = bufio.NewReader(conn)
		res, err := protocol.Read(r)
		if err != nil && err != io.EOF {
			fmt.Println(err)
		} else if res.Type() == protocol.Nil {
			fmt.Println("Nil")
		} else {
			msg, _ := res.String()
			fmt.Println(msg)
		}
	}
}
