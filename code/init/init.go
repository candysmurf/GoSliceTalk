// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows slice len and cap at initial states.
package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10)
	fmt.Printf("Slice:%v\tLen:%d\t\tCap:%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Printf("Slice:%v\t\tLen:%d\t\tCap:%d\n", s2, len(s2), cap(s2))

	s5 := []int{1, 2, 3}
	fmt.Printf("Slice:%v\t\t\tLen:%d\t\tCap:%d\n", s5, len(s5), cap(s5))

	var s3 []int
	fmt.Printf("Slice:%v\t\t\tLen:%d\t\tCap:%d\n", s3, len(s3), cap(s3))

	s4 := []int{}
	fmt.Printf("Slice:%v\t\t\tLen:%d\t\tCap:%d\n", s4, len(s4), cap(s4))

}
