// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO:
//   HashSet이 Set을 구현했는지 안 했는지 테스트케이스 추가하기.
package set

import (
	"fmt"
	"testing"
)

// TODO
func TestBasics(t *testing.T) {
}

func ExampleHashSet() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add("bar")
	fmt.Printf("%v\n", hs.Contains("foo"))
	fmt.Printf("%v\n", hs.Contains("baz"))
	fmt.Printf("%v\n", hs.Len())
	// Output:
	// true
	// false
	// 2
}

func ExampleRemove() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add("bar")
	hs.Remove("bar")

	fmt.Printf("%v\n", hs.Len())
	hs.Clear()
	fmt.Printf("%v\n", hs.Len())
	// Output:
	// 1
	// 0
}

func ExampleIter() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add(1)
	hs.Add(3.14)

	for v := range hs.Iter() {
		fmt.Println(v)
	}
	// Output:
	// foo
	// 1
	// 3.14
}
