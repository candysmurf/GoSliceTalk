// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows append works for unknown length data
package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("Slice:%v\t\t\t\tLen:%d\t\tCap:%d\n", s, len(s), cap(s))

	// append items one by one
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("Slice:%v\t\tLen:%d\t\tCap:%d\n", s, len(s), cap(s))

	ss := []int{10, 20, 30}
	s = append(s, ss...)
	fmt.Printf("Slice:%v\tLen:%d\t\tCap:%d\n", s, len(s), cap(s))
}
