package tree

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4 
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	// 如果有一个是空的话 则说明没有子结构
	if A == nil || B == nil {
		return false
	}
	// 只要满足一个条件的话 就是子树
	return isSub2(A, B) || IsSubStructure(A.Left,B) || IsSubStructure(A.Right,B)
}

func isSub2(A *TreeNode, B *TreeNode) bool {
	// 如果B树先达到根节点了 说明是子树
	if B == nil {
		return true
	}
	// 如果A树先达到根节点的话 则B树不是A树的子树
	// 如果A B树不相同的话 则B树不是A树的子树
	if A == nil || A.Val != B.Val{
		return false
	}
	// 左右子树开始递归
	return isSub2(A.Left, B.Left) && isSub2(A.Right, B.Right)
}
