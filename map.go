//  Copyright (c) 2019 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//  http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package script

import (
	"unsafe"
)

type RHMap struct {
	Parent    *RHMap
	hashCount uint32
	count     uint32
	hashItems []hashItem
	items     []Value
}

// hashItem represents an entry in the Map.
type hashItem struct {
	Key      Value
	Val      Value
	Distance int // How far item is from its best position.
}

func NewMap(size int) *RHMap {
	return &RHMap{hashItems: make([]hashItem, int64(size))}
}

func (m *RHMap) Len() int { return int(m.count) + int(m.hashCount) }

func (m *RHMap) MapLen() int { return int(m.hashCount) }

func (m *RHMap) ArrayLen() int { return int(m.count) }

// Clear clears Map, where already allocated memory will be reused.
func (m *RHMap) Clear() {
	m.hashItems = m.hashItems[:0]
	m.items = m.items[:0]
	m.count, m.hashCount = 0, 0
}

func (m *RHMap) GetString(k string) (v Value) {
	return m.Get(Str(k))
}

// Get retrieves the value for a given key.
func (m *RHMap) Get(k Value) (v Value) {
	if k == Nil {
		return Nil
	}
	if k.IsInt() {
		if idx := k.Int(); idx >= 0 && idx < int64(len(m.items)) {
			v = m.items[idx]
			goto FINAL
		}
	}
	if idx := m.findHash(k); idx >= 0 {
		v = m.hashItems[idx].Val
	} else if m.Parent != nil {
		v = m.Parent.Get(k)
	}
FINAL:
	if m.Parent != nil && v.Type() == FUNC {
		f := *v.Func()
		f.MethodSrc = m.Value()
		v = f.Value()
	}
	return v
}

func (m *RHMap) findHash(k Value) int {
	num := len(m.hashItems)
	if num <= 0 {
		return -1
	}
	idx := int(k.HashCode() % uint64(num))
	idxStart := idx

	for {
		e := &m.hashItems[idx]
		if e.Key == Nil {
			return -1
		}

		if e.Key.Equal(k) {
			return idx
		}

		idx++
		if idx >= num {
			idx = 0
		}

		if idx == idxStart { // Went all the way around.
			return -1
		}
	}
}

func (m *RHMap) Contains(k Value) bool {
	if k == Nil {
		return false
	}
	if k.IsInt() {
		if idx := k.Int(); idx >= 0 && idx < int64(len(m.items)) {
			return true
		}
	}
	return m.findHash(k) >= 0
}

func (m *RHMap) ParentContains(k Value) *RHMap {
	if k == Nil {
		return nil
	}
	if m.Parent != nil {
		p := m.Parent.ParentContains(k)
		if p != nil {
			return p
		}
	}
	if m.Contains(k) {
		return m
	}
	return nil
}

func (m *RHMap) SetString(k string, v Value) (prev Value) {
	return m.Set(Str(k), v)
}

// Set inserts or updates a key/val pair into the Map. If val == Nil, then key will get deleted
func (m *RHMap) Set(k, v Value) (prev Value) {
	if k == Nil {
		panicf("table set with nil key")
	}

	if m.Parent != nil && v.Type() != FUNC {
		if x := m.ParentContains(k); x != nil && x != m {
			return x.Set(k, v)
		}
	}

	if k.IsInt() {
		idx := k.Int()
		if idx >= 0 && idx < int64(len(m.items)) {
			prev, m.items[idx] = m.items[idx], v
			if v == Nil && prev != Nil {
				m.count--
			} else if v != Nil && prev == Nil {
				m.count++
			}
			return prev
		}
		if idx == int64(len(m.items)) {
			m.delHash(k)
			if v != Nil {
				m.items = append(m.items, v)
				m.count++
				return Nil
			}
			return Nil
		}
	}

	if v == Nil {
		return m.delHash(k)
	}

	if len(m.hashItems) <= 0 {
		m.hashItems = make([]hashItem, 8)
	}

	prev, _ = m.setHash(hashItem{Key: k, Val: v, Distance: 0})
	return
}

func (m *RHMap) setHash(incoming hashItem) (prev Value, growed bool) {
	num := len(m.hashItems)
	idx := int(incoming.Key.HashCode() % uint64(num))

	for idxStart := idx; ; {

		e := &m.hashItems[idx]
		if e.Key == Nil {
			m.hashItems[idx] = incoming
			m.hashCount++
			return Nil, false
		}

		if e.Key.Equal(incoming.Key) {
			prev = e.Val
			e.Val, e.Distance = incoming.Val, incoming.Distance
			return prev, false
		}

		// Swap if the incoming item is further from its best idx.
		if e.Distance < incoming.Distance {
			incoming, m.hashItems[idx] = m.hashItems[idx], incoming
		}

		incoming.Distance++ // One step further away from best idx.
		idx = (idx + 1) % num

		// Grow if distances become big or we went all the way around.
		if num < 8 {
			if idx == idxStart {
				m.resize(num + 1)
				prev, _ = m.setHash(incoming)
				return prev, true
			}
		} else {
			if int(m.hashCount) >= num/2 || idx == idxStart {
				m.resize(num*2 + 1)
				prev, _ = m.setHash(incoming)
				return prev, true
			}
		}
	}
}

func (m *RHMap) delHash(k Value) (prev Value) {
	idx := m.findHash(k)
	if idx < 0 {
		return Nil
	}
	prev = m.hashItems[idx].Val

	// Left-shift succeeding items in the linear chain.
	for {
		next := idx + 1
		if next >= len(m.hashItems) {
			next = 0
		}

		if next == idx { // Went all the way around.
			break
		}

		f := &m.hashItems[next]
		if f.Key == Nil || f.Distance <= 0 {
			break
		}

		f.Distance--

		m.hashItems[idx] = *f

		idx = next
	}

	m.hashItems[idx] = hashItem{}
	m.hashCount--
	return prev
}

func (m *RHMap) Foreach(f func(k, v Value) bool) {
	for k, v := m.Next(Nil); k != Nil; k, v = m.Next(k) {
		if !f(k, v) {
			return
		}
	}
}

func (m *RHMap) Next(k Value) (Value, Value) {
	nextHashPair := func(start int) (Value, Value) {
		for i := start; i < len(m.hashItems); i++ {
			if i := &m.hashItems[i]; i.Key != Nil {
				return i.Key, i.Val
			}
		}
		return Nil, Nil
	}
	if k == Nil {
		if len(m.items) == 0 {
			return nextHashPair(0)
		}
		return Int(0), m.items[0]
	}
	if k.IsInt() {
		n := k.Int()
		if n >= 0 && n < int64(len(m.items))-1 {
			for n++; n < int64(len(m.items)); n++ {
				if m.items[n] != Nil {
					return Int(n), m.items[n]
				}
			}
		}
		return nextHashPair(m.findHash(k) + 1)
	}
	idx := m.findHash(k)
	if idx < 0 {
		return Nil, Nil
	}
	return nextHashPair(idx + 1)
}

func (m *RHMap) Array() []Value {
	return m.items
}

func (m *RHMap) String() string {
	return m.Value().String()
}

func (m *RHMap) Value() Value {
	if m == nil {
		return Nil
	}
	return Value{v: uint64(MAP), p: unsafe.Pointer(m)}
}

func (m *RHMap) resize(newSize int) {
	tmp := RHMap{hashItems: make([]hashItem, newSize)}
	for _, e := range m.hashItems {
		if e.Key != Nil {
			e.Distance = 0
			tmp.setHash(e)
		}
	}
	m.hashItems = tmp.hashItems
}
