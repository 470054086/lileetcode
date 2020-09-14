package main

import (
	"fmt"
	"lileetcode.com/letcodearray"
)

var s = 0

func main() {
	//nums := []int{0,0,1,1,1,2,2,3,3,4}
	//duplicates := letcodearray.RemoveDuplicatesUpgrade(nums)
	//fmt.Println(nums)
	//fmt.Println(duplicates)
    // 总长度是3
	// [3 0] ->[0 0]
	// [3 1] ->[1 0]
	// [3 2] ->[2 0]

	// [1 0] -> [0 1]
	// [1 1] -> [1 1]
	// [1 2] -> [2 1 ]

	s := []int{1,2,3}
	permute := letcodearray.Permute(s)
	fmt.Println(permute)
}
