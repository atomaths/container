// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HashSet implements blah blah..
package set

// HashSet is...
type HashSet struct {
	data map[interface{}]bool
}

// NewHashSet returns a new hashset.
func NewHashSet() *HashSet {
	return &HashSet{
		data: make(map[interface{}]bool),
	}
}

// Add is blah blah.. 이미 값이 있으면 overwrite 됨
func (h *HashSet) Add(v interface{}) {
	h.data[v] = true
}

// Contains blah blah..
func (h *HashSet) Contains(v interface{}) bool {
	_, ok := h.data[v]
	return ok
}

func (h *HashSet) Length() uint {
	return uint(len(h.data))
}

func (h *HashSet) Remove(v interface{}) bool {
	delete(h.data, v)
	return true
}

func (h *HashSet) Clear() {
	for k := range h.data {
		delete(h.data, k)
	}
}
