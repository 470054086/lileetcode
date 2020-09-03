package main

import (
	"fmt"
	"lileetcode.com/letcodearray"
)

type Name struct {
	name string
}
func main() {
	//toInt := lettcodestring.RomanToInt("LVIII")
	//fmt.Println(toInt)
	s:= []string{"flower","f","flow","floight"}
	prefix := letcodearray.CommonPrefix(s)
	fmt.Println(prefix)
}
