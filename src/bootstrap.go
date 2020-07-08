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
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func bootstrap() {
	daemonptr := flag.Bool("d", false, "start as a daemon")
	flag.Parse()
	if *daemonptr {
		if _, isDaemon := os.LookupEnv("MEMDB_DAEMON"); !isDaemon {
			daemonEnv := []string{"MEMDB_DAEMON=true"}
			childPid, _ := syscall.ForkExec(os.Args[0], os.Args, &syscall.ProcAttr{
				Env: append(os.Environ(), daemonEnv...),
				Sys: &syscall.SysProcAttr{
					Setsid: true,
				},
			})
			fmt.Printf("running memdb daemon with pid %d\n", childPid)
			return
		}
	}
	ReadConfig()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	// manager := NewMemoryManager()
	var conn *Conn
	go func() {
		conn = NewConnection()
		conn.StartConnection()
	}()
	if !*daemonptr {
		fmt.Printf(Logo, "0.1.0", runtime.GOARCH, OperationMode, Port)
	}
	intr := <-exit
	conn.StopConnection()
	if !*daemonptr {
		fmt.Println()
		fmt.Println(intr)
		fmt.Println("bye.")
	}
}
