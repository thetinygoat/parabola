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
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", ":9898")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		txt, _ := reader.ReadString('\n')
		fmt.Fprintf(c, txt+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(message)
	}
}
