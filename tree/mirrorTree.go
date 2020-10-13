package tree

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MirrorTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	// 交换位置
	node.Left, node.Right = MirrorTree(node.Right), MirrorTree(node.Left)
	return node
}

func MirrorTreeStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 定义一个栈
	stack := []*TreeNode{}
	// 先把根节点加入
	stack = append(stack, root)
	for len(stack) != 0 {
		//弹出当前的节点
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		t.Left, t.Right = t.Right, t.Left
	}
	return root
}
