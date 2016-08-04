// All material is licensed under the Apache License Version 2.0, August 2016
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows that the initial cap allocation varies based on the data type.
// Afterward, it doubles.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sl []bool
	fmt.Printf("bool slice begining\tcap:%d\n", cap(sl))
	for i := 0; i < 10; i++ {
		sl = append(sl, true)
		fmt.Printf("bool slice append %d\tcap:%d\n", (i + 1), cap(sl))
	}

	var sb []byte
	fmt.Println()
	fmt.Printf("byte slice begining\tcap:%d\n", cap(sb))
	for i := 0; i < 10; i++ {
		sb = append(sb, byte(i))
		fmt.Printf("byte slice append %v\tcap:%v\n", (i + 1), cap(sb))
	}

	var si []int
	fmt.Println()
	fmt.Printf("int slice begining\tcap:%d\n", cap(si))
	for i := 0; i < 10; i++ {
		si = append(si, i)
		fmt.Printf("int slice append %d\tcap:%d\n", (i + 1), cap(si))
	}

	var ss []string
	fmt.Println()
	fmt.Printf("string slice begining\tcap:%d\n", cap(ss))
	for i := 0; i < 10; i++ {
		ss = append(ss, strconv.Itoa(i))
		fmt.Printf("string slice append %d\tcap:%d\n", (i + 1), cap(ss))
	}

	var sf []float64
	fmt.Println()
	fmt.Printf("string slice begining\tcap:%d\n", cap(sf))
	for i := 0; i < 10; i++ {
		sf = append(sf, float64(i))
		fmt.Printf("float slice append %v\tcap:%v\n", (i + 1), cap(sf))
	}

}
