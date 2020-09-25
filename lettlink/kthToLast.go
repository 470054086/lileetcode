package lettlink
/*
面试题 02.02. 返回倒数第 k 个节点
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。
 */
func KthToLast(head *ListNode, k int) int {
	// 如果是空链接表 直接返回
	if head == nil {
		return 0
	}
	// 计算链表的长度
	p := head
	n := 0
	for p != nil {
		n++
		p = p.Next
	}
	//循环的节点
	diffN := n - k + 1
	n = 0
	for head != nil {
		n++
		if diffN == n {
			return head.Val
		}
		head = head.Next
	}
	return 0
}

func DoublueLine(head *ListNode, k int) int {
	// 使用快慢指针  记录K个节点
	// 1. 距离K点 只需要快指针先走k个点 快指针走到最后了即是记录的节点
	left := head  //慢指针
	right := head //快指针
	// 快指针先移动
	for i := 0; i < k; i++ {
		right = right.Next
	}
	// 快指针达到最后即可
	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left.Val

}
