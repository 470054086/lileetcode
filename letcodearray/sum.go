package letcodearray

/**

https://leetcode-cn.com/problems/two-sum/

题目:在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
example:

Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

*/

/**
1. 使用暴力解法 使用两次循环 进行处理
*/
func TwoSum(nums []int, target int) []int {
	for k, v := range nums {
		diff := target - v
		for i := k + 1; i <= len(nums)-1; i++ {
			if diff == nums[i] {
				return []int{k, i}
			}
		}
	}
	return []int{0, 0}
}

/**
使用map进行解题

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，
如果找到了，直接返回 2 个数字的下标即可。如果找不到，
就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。
*/
func TwoSumUpgrade(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		diff := target - v
		if _, ok := m[diff]; ok {
			return []int{m[diff], k}
		}
		m[v] = k
	}
	return []int{0,0}
}
