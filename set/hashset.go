// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set

type HashSet struct {
}

func NewHashSet() *HashSet {
	return &HashSet{}
}

func (s *HashSet) Add(v interface{}) {
}

func (s *HashSet) Contains(v interface{}) bool {
	return false
}
