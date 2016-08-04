// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows that the valid index is in the range of x <= y <= cap(s) for slice s[x:y].
package main

import (
	"fmt"
)

func main() {
	// initialize s with length 5 and capacity 10.
	s := make([]int, 5, 10)

	// loop over slice s.
	for i := 0; i < 5; i++ {
		s[i] = i
	}
	fmt.Printf("Slice: %v\tLen: %v\t\tCap: %v\n", s, len(s), cap(s))

	// create a new slice ss.
	ss := s[5:]
	fmt.Printf("Slice: %v\t\tLen: %v\t\tCap: %v\n", ss, len(ss), cap(ss))
}
