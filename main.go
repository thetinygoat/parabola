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

import "flag"

func main() {
	port := flag.String("port", "9898", "Port to run DictX server on")
	maxMemory := flag.Int64("mem", 1073741824, "Max memory limit for the DictX server")
	flag.Parse()
	srv, _ := NewServer(*port, *maxMemory)
	srv.Listen()
}
