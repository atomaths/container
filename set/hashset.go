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

// Add is blah blah..
func (h *HashSet) Add(v interface{}) {
	h.data[v] = true
}

// Contains blah blah..
func (h *HashSet) Contains(v interface{}) bool {
	_, ok := h.data[v]
	return ok
}

// 아래 구조가 필요가 없을래나...?
var hs = NewHashSet()

func init() {
	hs.data = make(map[interface{}]bool)
}
