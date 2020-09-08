package match

/**
括号匹配问题 所有栈 先进先出
如果遇到( 就进栈 如果遇到) 就出栈 最后判断栈的大小 如果为0则进行匹配成功

https://leetcode-cn.com/problems/valid-parentheses/

*/

func IsValid(s string) bool {
	// 使用slice 实现队列
	queen := []byte{}
	flag := false
	for i := 0; i < len(s); i++ {
		isvaild := false
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			queen = append(queen, s[i])
			flag = true
		}
		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(queen) == 0 {
				return false
			}
			brack := queen[len(queen)-1]
			if s[i] == ')' && brack == '(' {
				isvaild = true
				if len(queen) == 1 {
					queen = []byte{}
				} else {
					queen = queen[:len(queen)-1]
				}
			}
			if s[i] == ']' && brack == '[' {
				isvaild = true
				if len(queen) == 1 {
					queen = []byte{}
				} else {
					queen = queen[:len(queen)-1]
				}
			}
			if s[i] == '}' && brack == '{' {
				isvaild = true
				if len(queen) == 1 {
					queen = []byte{}
				} else {
					queen = queen[:len(queen)-1]
				}
			}
			if isvaild == false {
				return false
			}
		}
	}
	if flag == false {
		return false
	}
	if len(queen) == 0 {
		return true
	}
	return false
}

// 使用map进行优化
func IsValidUpgrade(s string) bool {
	n := len(s)
	// 如果是基数的话 说明括号肯定不匹配
	if n%2 == 1 {
		return false
	}
	brakMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	queue := []byte{}
	for i := 0; i < n; i++ {
		if brakMap[s[i]] > 0 {
			if len(queue) == 0 || queue[len(queue)-1] != brakMap[s[i]] {
				return false
			}
			queue = queue[:len(queue)-1]
		} else {
			queue = append(queue, s[i])
		}
	}
	return len(queue) == 0

}
