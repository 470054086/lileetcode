package lettcodestring

import "fmt"

/**
罗马转化为整数的问题
https://leetcode-cn.com/problems/roman-to-integer/
*/

// 进行罗马数字的循环
func RomanToInt(s string) int {
	//从罗马后面的数字开始循环
	pre := 0
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := showRoman(s[i])
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}
	return r
}

func showRoman(r byte) int {
	// 罗马 对应的数字
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		panic(fmt.Sprintf("impossible: %q", string(r)))
	}
}
