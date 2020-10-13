package tree
/**
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]
 */
func LeftRightQueue(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 定义保存val的数组
	stack := []int{}
	// 定义一个队列
	queueNode := []*TreeNode{}
	queueNode = append(queueNode, root)
	for len(queueNode) != 0 {
		// 出队列
		node := queueNode[0]
		queueNode = queueNode[1:]
		// 添加数值到数组
		stack = append(stack, node.Val)
		// 入队列
		if node.Left != nil {
			queueNode = append(queueNode, node.Left)
		}
		if node.Right != nil {
			queueNode = append(queueNode, node.Right)
		}
	}
	return stack
}
