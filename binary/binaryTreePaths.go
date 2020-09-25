package binary

import "strconv"

/**
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
通过次数77,095提交次数116,796
*/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**

 */
var res []string

func BinaryTreePaths(node *TreeNode) []string {
	if node == nil {
		return nil
	}
	res = make([]string, 0)
	dfs4(node, "")
	return res
}

func dfs4(node *TreeNode, s string) {
	// 如果数组是空的话 则递归结束
	if node == nil {
		return
	}
	// 如果两个都是子节点都是空的话 说明是叶子节点 递归结束
	if node.Left == nil && node.Right == nil {
		res = append(res, s+strconv.Itoa(node.Val))
		return
	}
	// 进行左右节点的递归
	dfs4(node.Left, s+strconv.Itoa(node.Val)+"->")
	dfs4(node.Right, s+strconv.Itoa(node.Val)+"->")
}
