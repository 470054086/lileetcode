package letcodearray

import (
	"fmt"
	"sort"
)

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
 

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

/**
有些朋友可能会疑惑什么时候使用 used 数组，什么时候使用 begin 变量。这里为大家简单总结一下：

排列问题，讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为不同列表时），需要记录哪些数字已经使用过，此时用 used 数组；
组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为相同列表时），需要按照某种顺序搜索，此时使用 begin 变量。
注意：具体问题应该具体分析， 理解算法的设计思想 是至关重要的，请不要死记硬背。

解题详细解答
作者：liweiwei1419
链接：https://leetcode-cn.com/problems/combination-sum/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
var res2 [][]int

func CombinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return nil
	}
	res2 = make([][]int, 0)
	path := make([]int, 0)
	// 对数组进行排序
	sort.Ints(candidates)

	dfs3(candidates, 0, n, path, target)
	return res2
}

func dfs3(candidates []int, deep int, n int, path []int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// 使用拷贝 生成一个全新的变量
		// 不会因为加入了 印象另外一个变量的使用
		t := make([]int, len(path))
		copy(t, path)
		res2 = append(res2, t)
		return
	}
	for i := deep; i < n; i++ {
		// 因为对数组进行了排序 如果target-当前的值小于0的话 则后面的输是不用递归的
		if target-candidates[i] < 0 {
			break
		}
		path = append(path, candidates[i])
		fmt.Println("递归之前i的值为", i)
		fmt.Println("递归之前---:", path, "剩余数字是:", target-candidates[i])
		// 这个i 如果上面的循环不改变的话 每次传入的值还是不会改变 这样的话就搜索了自己
		dfs3(candidates, i, n, path, target-candidates[i])
		path = path[:len(path)-1]
		fmt.Println("递归之后", path)
	}
}
