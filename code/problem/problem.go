// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows that appending items to the pre-allocated slice.
package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("Slice:%v\tLen:%d\tCap:%d\n", s, len(s), cap(s))
}
