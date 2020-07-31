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

package key

import (
	"time"

	"github.com/emirpasic/gods/maps/hashmap"
)

type value struct {
	val      string
	addedAt  time.Time
	expireAt time.Time
}

type Key struct {
	key    *hashmap.Map
	memory int64
}

func New() *Key {
	return &Key{key: hashmap.New(), memory: 0}
}

func (k *Key) Set(key, val string, ttl int) {
	ov, found := k.key.Get(key)
	if found {
		k.free(ov.(*value).val)
	}
	v := &value{val: val, addedAt: time.Now(), expireAt: getExpireTime(ttl)}
	k.key.Put(key, v)
	k.allocate(val)
}

func (k *Key) Get(key string) (v string, ok bool) {
	val, found := k.key.Get(key)
	if !found {
		return "", false
	}
	if isExpired(val.(*value).expireAt) {
		k.free(val.(*value).val)
		k.key.Remove(key)
		return "", false
	}
	return val.(*value).val, true
}

func (k *Key) Del(key string) {
	val, found := k.key.Get(key)
	if !found {
		return
	}
	k.free(val.(*value).val)
	k.key.Remove(key)
}

func (k *Key) Memory() int64 {
	return k.memory
}

func getExpireTime(ttl int) time.Time {
	return time.Now().Add(time.Duration(ttl) * time.Second)
}
func isExpired(expireAt time.Time) bool {
	if time.Now().Sub(expireAt) >= 0 {
		return true
	}
	return false
}

func (k *Key) allocate(data string) {
	k.memory += int64(len(data))
}

func (k *Key) free(data string) {
	k.memory -= int64(len(data))
}
