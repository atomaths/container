// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set

import (
	"fmt"
	"testing"
)

// TODO
func TestBasics(t *testing.T) {
}

func ExampleRemove() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add("bar")
	hs.Remove("bar")

	fmt.Printf("%v\n", hs.Length())
	hs.Clear()
	fmt.Printf("%v\n", hs.Length())
	// Output:
	// 1
	// 0
}

func ExampleHashSet() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add("bar")
	fmt.Printf("%v\n", hs.Contains("foo"))
	fmt.Printf("%v\n", hs.Contains("baz"))
	fmt.Printf("%v\n", hs.Length())
	// Output:
	// true
	// false
	// 2
}
