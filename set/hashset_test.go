// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set

import (
	"fmt"
)

func ExampleHashSet() {
	hs := NewHashSet()
	hs.Add("foo")
	hs.Add("bar")
	fmt.Printf("%v\n", hs.Contains("foo"))
	fmt.Printf("%v", hs.Contains("baz"))
	// Output:
	// false
	// false
}
