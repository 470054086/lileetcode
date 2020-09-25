package lettlink

/**
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。

示例：

输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-middle-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func DeleteNode(node *ListNode) *ListNode {
	//快指针
	right := node
	left := node
	// 如果快指针为空 则证明 走到了头 慢指针此时的节点就是中间节点
	// 但是获取到当前节点的中间节点呢
	for right != nil {
		if right.Next == nil {
			node.Next = node.Next.Next
		} else {
			right = right.Next.Next
			left = left.Next
		}
	}
	return left
}

/**
  删除链表的中间节点
*/
func DeleteMidNode(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return nil
	}
	if node.Next.Next == nil {
		return node.Next
	}
	// 快指针
	right := node.Next.Next
	// 慢指针
	left := node
	for right.Next != nil && right.Next.Next != nil {
		left = left.Next
		right = right.Next.Next
	}
	// 进行删除节点
	left.Next = left.Next.Next
	return node
}


