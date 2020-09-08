package letcodearray

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/

/**
1. 通过两次循环比较  如果找到相同的 就改变当前获取的
2. 如果循环之后 相同的是0个 则停止循环
*/

func CommonPrefix(s []string) string {
	// 获取数组的第一个
	if len(s) == 0 {
		return ""
	}
	pref := s[0]
	for i := 1; i < len(s); i++ {
		pref = lcp(pref, s[i])
		if len(pref) == 0 {
			break
		}
	}
	return pref
}

func lcp(pre, s string) string {
	// 取两个字符串最小的长度
	lens := min(len(pre), len(s))
	index := 0
	for index < lens && pre[index] == s[index] {
		index++
	}
	return s[:index]
}

/**
纵向比较法
*/

func LongestCommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	// 外层是S[0] 第一个参数 将他和后面的数据进行纵向比较
	// 如果到最后 和 比较到不相等的时候 就进行返回即可
	// 如果循环完了 都是相等的 则返回S[0]即可
	for i := 0; i < len(s[0]); i++ {
		for j := 1; j < len(s); j++ {
			// 如果循环到最后
			// 1.如果当前的不等于s[0]的数据
			if i == len(s[j]) || s[j][i] != s[0][i] {
				return s[0][:i]
			}
		}
	}
	return s[0]
}
