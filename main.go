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
	"encoding/json"
	"log"
	"os"

	"github.com/tidwall/evio"
)

type config struct {
	Port      string
	Maxmemory uint64 `json:",string"`
}

func readConfig() *config {
	f, e := os.Open("dictX.json")
	if e != nil && os.IsNotExist(e) {
		log.Fatal("error: config file not found.")
		return nil
	} else if e != nil {
		log.Fatal(e)
		return nil
	}
	decoder := json.NewDecoder(f)
	var c config
	decoder.Decode(&c)

	return &c
}

func main() {
	config := readConfig()
	var ev evio.Events
	ev.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("dictX server started on port %s \n", config.Port)
		return
	}
	ev.Opened = func(conn evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		log.Printf("opened: %s\n", conn.RemoteAddr())
		return
	}
	ev.Data = func(conn evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = in
		return
	}
	if err := evio.Serve(ev, "tcp://localhost:"+config.Port); err != nil {
		panic(err)
	}
}
