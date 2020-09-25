package lettcodestring

import (
	"strings"
)

func Substirngs(s string) int {
	// 判断字符串的长度
	if len(s) == 0 {
		return 0
	}
	// 定义双指针
	left := 0  // 左指针
	right := 0 // 有指针
	// 定义滑动窗口
	windows := s[left:right]
	res := 0
	// 循环字符串
	for ; right < len(s); right++ {
		// 如果当前的字符串在滑动窗口中 则把left指针像前移动
		if index := strings.IndexByte(windows, s[right]); index != -1 {
			left += index + 1
		}
		// 截取left到right的中间数组 即是不重复的
		windows = s[left : right+1]
		// 获取不重复的最大值
		if len(windows) > res {
			res = len(windows)
		}
	}
	return res
}
