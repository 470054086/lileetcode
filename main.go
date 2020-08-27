package main

import (
	"fmt"
	"lileetcode.com/letcodearray"
)

type Name struct {
	name string
}
func main() {
	reverse := letcodearray.ReverseUpgrade(2147483412)
	fmt.Println(reverse)
}
