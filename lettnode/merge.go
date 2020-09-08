package lettnode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
迭代法
组合两个链表  最小的在最前面
只要两个链表都为空的话 则证明链表循环完毕  循环的时候 判断最小的在前面即可
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果都是null的话 则直接退出
	if l1 == nil && l2 == nil {
		return nil
	}
	// 头结点
	pthread := &ListNode{
		Val:  -1,
		Next: nil,
	}
	// 当节点
	res := pthread
	for l1 != nil && l2 != nil {
		// 如果小于的话 则下一个链接是l1
		//  l1 再次向下移动
		if l1.Val < l2.Val {
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		// res向前移动一个节点
		res = res.Next
	}
	// 到最后 可能会存在还有一个节点为移动完毕
	// 未移动完毕的话 说明后面都是已经排序好了的 直接放入当前的next即可
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}
	// 返回投节点的下一个节点即可
	return pthread.Next
}

func MergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}else if l1.Val<l2.Val {
		l1.Next = MergeTwoListsRecursion(l1.Next,l2)
		return l1
	}else {
		l2.Next = MergeTwoListsRecursion(l1,l2.Next)
		return l2
	}
}


