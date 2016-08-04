// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows that a slice slot is accessible when its index < len(s).
// Otherwise, a panic occurs.
package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5, 10)
	for i := 0; i < 6; i++ {
		s[i] = i
	}
	fmt.Printf("Slice: %v\tLen: %v\t\tCap: %v\n", s, len(s), cap(s))

	ss := s[6:]

	fmt.Printf("Slice: %v\t\tLen: %v\t\tCap: %v\n", ss, len(ss), cap(ss))
}
