package letcodearray

/**
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

 

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 只能在原数组进行处理 那就进行循环每个数字进行比较
// 速度应该慢在每次循环都要对数组进行偏移处理 这个的操作是oN2
func RemoveDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	// 进行数组循环
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
	}
	return len(nums)
}
func RemoveDuplicatesUpgrade(nums []int) int {
	if nums == nil {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		// 如果相等的话 则移动指针
		// 交换数组
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
