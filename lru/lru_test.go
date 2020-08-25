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

package lru

import "testing"

func TestSet(t *testing.T) {
	l := New(200)
	l.Set("key1", "value1", 20)
	l.Set("key2", "value2", 30)
	l.Set("key3", "value3", 40)
	if _, found := l.keymap["key1"]; !found {
		t.Errorf("failed for key1")
	}
	if _, found := l.keymap["key2"]; !found {
		t.Errorf("failed for key")
	}
	if _, found := l.keymap["key3"]; !found {
		t.Errorf("failed for key3")
	}
}
func TestGet(t *testing.T) {
}
