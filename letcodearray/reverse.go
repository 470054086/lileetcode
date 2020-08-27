package letcodearray

import (
	"fmt"
	"math"
	"strconv"
)
/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

链接：https://leetcode-cn.com/problems/reverse-integer
 */


// 整数反转
// 使用字符串进行反转
func Reverse(x int) int {
	isNeg := false // 是否是负数
	if x < 0 {
		isNeg = true
		x = -x
	}
	// 转化为字符串
	s := fmt.Sprintf("%v", x)
	s2 := reverseString(s)
	// 字符串转化为数字
	atoi, err := strconv.Atoi(s2)
	if err == nil {
		if isNeg == true {
			atoi = -atoi
		}
		if atoi > math.MaxInt32 {
			return 0
		}
		if atoi < math.MinInt32 {
			return 0
		}
		return atoi
	} else {
		return 0
	}
}

/**
反转整数的方法可以与反转字符串进行类比。

我们想重复“弹出” xx 的最后一位数字，并将它“推入”到 \text{rev}rev 的后面。最后，\text{rev}rev 将与 xx 相反。

要在没有辅助堆栈 / 数组的帮助下 “弹出” 和 “推入” 数字，我们可以使用数学方法。


//pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop;
rev = temp;

为了防止溢出 我们需要检测溢出的问题

链接：https://leetcode-cn.com/problems/reverse-integer/solution/zheng-shu-fan-zhuan-by-leetcode/

 */
func reverseString(s string) string {
	// 转化为Rune
	r := []rune(s)
	// 进行字符串反转
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 使用 %10 就是当前的数字

func ReverseUpgrade(x int) int {
	rev := 0
	for x != 0 {
		// 判断是否溢出
		if temp := int32(rev); (temp*10)/10 != temp {
			return 0
		}
		// 弹出
		pop := x % 10
		x = x / 10
		// 加这个数字加入
		rev = rev*10 + pop
	}
	return rev
}
