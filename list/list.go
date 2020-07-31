//Copyright 2020 Sachin Saini

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package list

import dll "github.com/emirpasic/gods/lists/doublylinkedlist"

type List struct {
	list   map[string]*dll.List
	memory int64
}

func New() *List {
	return &List{list: make(map[string]*dll.List), memory: 0}
}

func (l *List) Append(key, value string) {
	_, found := l.list[key]
	if !found {
		l.list[key] = dll.New()
	}
	l.list[key].Append(value)
	l.allocate(value)
}

func (l *List) Prepend(key, value string) {
	_, found := l.list[key]
	if !found {
		l.list[key] = dll.New()
	}
	l.list[key].Prepend(value)
	l.allocate(value)
}

func (l *List) RemoveFirst(key string) (string, bool) {
	list, found := l.list[key]
	if !found {
		return "", false
	}
	v, ok := list.Get(0)
	if !ok {
		return "", false
	}
	sv := str(v)
	l.free(sv)
	return sv, true
}

func (l *List) RemoveLast(key string) (string, bool) {
	list, found := l.list[key]
	if !found {
		return "", false
	}
	v, ok := list.Get(list.Size() - 1)
	if !ok {
		return "", false
	}
	sv := str(v)
	l.free(sv)
	return sv, true
}

func (l *List) Size(key string) int64 {
	list, found := l.list[key]
	if !found {
		return -1
	}
	return int64(list.Size())
}

func (l *List) allocate(data string) {
	l.memory += int64(len(data))
}

func (l *List) free(data string) {
	l.memory -= int64(len(data))
}

func str(data interface{}) string {
	return data.(string)
}
