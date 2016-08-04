// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows that a slice is accessible when its index is between 0 to len(s)-1.
package main

import (
	"fmt"
)

func main() {
	// initialize s with length 5 and capacity 10.
	s := make([]int, 5, 10)

	// loop over slice s.
	for i := 0; i < 6; i++ {
		s[i] = i
	}
	fmt.Printf("Slice: %v\tLen: %v\t\tCap: %v\n", s, len(s), cap(s))

	// create a new slice ss.
	ss := s[6:]
	fmt.Printf("Slice: %v\t\tLen: %v\t\tCap: %v\n", ss, len(ss), cap(ss))
}
