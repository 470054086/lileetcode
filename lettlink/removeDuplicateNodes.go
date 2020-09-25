package lettlink

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 如果下一个节点为空 也直接返回
	if head.Next == nil {
		return head
	}

	// 使用一个map用于存储
	used := make(map[int]bool)
	// 使用双指针
	left := head
	right := head.Next
	used[left.Val] = true
	for left.Next != nil {
		// 如果存在的话
		if _, ok := used[right.Val]; ok {
			left.Next = right.Next
			right = right.Next
		} else {
			// 如果不存在的话
			left = left.Next
			used[right.Val] = true
		}
	}
	return head
}
