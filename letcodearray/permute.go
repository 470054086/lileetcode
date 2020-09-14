package letcodearray

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**

 */
var res [][]int

func Permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	// 构建二维
	res = make([][]int, 0)
	// 调用变量
	dfs2(nums, 0, n)
	return res
}

func dfs2(nums []int, pos int, n int) {
	//递归调用 增加判断条件
	// 如果达到最后一个则为一次排列
	if pos == n-1 {
		x := make([]int, n)
		// 拷贝一个变量
		copy(x, nums)
		res = append(res, x)
		return
	}
	/**
	   [1 2 3]  pos = 0
		nums[0],nums[0]  = nums[0][0]
		nums[1],nums[0]  = nums[0],nums[1]   [2,1,3]
		nums[2],nums[0] = nums[0],nums[2]	 [3,1,2]
		// 变回来
		[1 2 3]  pos = 1
		nums[1],nums[1] = nums[1]nums[1]
		nums[2],nums[1] = num[2]nums[1]
		nums[2],nums[2] = nums[2]nums[2]
		[1,2,3] pos = 2
	*/
	for i := pos; i < n; i++ {
		// 进行位置交换
		if i != pos {
			nums[i], nums[pos] = nums[pos], nums[i]
		}
		// 进行递归调用
		dfs2(nums, pos+1, n)
		// 将位置置换回去
		if i != pos {
			nums[i], nums[pos] = nums[pos], nums[i]
		}
	}
}
