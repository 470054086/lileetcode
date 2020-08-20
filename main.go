package main

import (
	"fmt"
	"lileetcode.com/letcodearray"
)

func main() {
	s := []int{1,8,6,2,5,4,8,3,7}
	sum := letcodearray.MaxAreaDoublePoint(s)
	fmt.Println(sum)
}
