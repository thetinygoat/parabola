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
	"fmt"
	"runtime"
)

func main() {
	ReadConfig()
	// manager := NewMemoryManager()
	// if _, isDaemon := os.LookupEnv("MEMDB_DAEMON"); !isDaemon {
	// 	daemonEnv := []string{"MEMDB_DAEMON=true"}
	// 	childPid, _ := syscall.ForkExec(os.Args[0], os.Args, &syscall.ProcAttr{
	// 		Env: append(os.Environ(), daemonEnv...),
	// 		Sys: &syscall.SysProcAttr{
	// 			Setsid: true,
	// 		},
	// 	})
	// 	fmt.Printf("running daemon with pid %d", childPid)
	// 	return
	// }

	fmt.Printf(Logo, "0.1.0", runtime.GOARCH, OperationMode, Port)
	conn := NewConnection("8080")
	conn.StartConnection()
}
