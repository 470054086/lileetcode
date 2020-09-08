package match

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func huiwen(itoa string, lens int) bool {
	var left string
	var right string
	// 如果是偶数的话 则不用舍弃中间的数字
	// 如果是基数的话 需要舍弃中间的数字
	if lens%2 == 0 {
		left = itoa[:lens/2]
		right = itoa[lens/2:]
	} else {
		left = itoa[:lens/2]
		right = itoa[lens/2+1:]
	}
	// 循环比较 left 从左到右循环 left的左边数字等于right右边数字
	for i := 0; i < len(right); i++ {
		if right[i] != left[len(right)-i-1] {
			return false
		}

	}
	return true
}

// 使用数字的方法 解决问题
/**
 	1. 先计算出最大位数 313 通过 循环 知道最大位数为3 则 除数是100
	2. 第一位为  313 / 100  最后一位为 313 % 10
	3. 将33从数字中去掉  x = (x%100) / 10 则剩余1了 少了两位 需要将div/100
*/

func HuiwenNum(x int) bool {
	// 如果小于0的话 则证明是负数 负数则不是回文
	if x < 0 {
		return false
	}
	// 计算div最大的位数
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x > 0 {
		// 前面的数字
		left := x / div
		// 后面的数字
		right := x % 10
		if left != right {
			return false
		}
		// 去除回文的第一位和最后一位
		x = (x % div) / 10
		div = div / 100
	}
	return true

}
