// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package set implements the Set data structure in Java or C++.

package set

type Set interface {
	Add(v interface{})
	Contains(v interface{}) bool
	Length() uint
	Remove(v interface{}) bool
	Clear()
}
