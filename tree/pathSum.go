package tree

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

 

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func PathSum(root *TreeNode, sum int) [][]int {
	// 定义 一个二维数组
	res := make([][]int, 0)
	// 记录使用的数据
	used := make([]int, 0)
	treeSum(root, used, sum, &res)
	return res
}

func treeSum(node *TreeNode, used []int, sum int, r *[][]int) {
	// 如果为节点的nil话则进行提出
	if node == nil {
		return
	}
	// 将当前节点加入
	used = append(used, node.Val)
	// 减少当前的值
	sum = sum - node.Val
	// 如果等于0话 并且是叶子节点的话 则直接找到了一个路径
	if sum == 0 && node.Left == nil && node.Right == nil {
		// 复制一个变量
		t := make([]int, len(used))
		copy(t, used)
		*r = append(*r, t)
	}
	// 进行递归
	treeSum(node.Left, used, sum, r)
	treeSum(node.Right, used, sum, r)
	// used进行回溯
	used = used[:len(used)-1]
}
