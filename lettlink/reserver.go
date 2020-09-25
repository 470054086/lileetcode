package lettlink

import (
	"strconv"
)

/**
反转链表
*/
func Reserver(head *ListNode) *ListNode {
	// 定义一个空节点
	var curr *ListNode
	p := head
	for p != nil {
		// 保留当前的下一个节点
		temp := p.Next
		//  当前节点的下一个为curr
		p.Next = curr
		//  curr为当前节点
		curr = p
		// p 指向下一个下一个节点
		p = temp
	}
	return curr
}

func Reserver2(node *ListNode) *ListNode {
	// 定义一个前指针(空指针)
	var pref *ListNode
	// 定义当前指针
	curr := node
	// 只要curr走到最后一个指针的时候 即走到头了
	for curr != nil {
		// 记录当前指针的下一个
		temp := curr.Next
		// 将当前指针的下一个指向前指针
		curr.Next = pref
		//  前一个指针当前指针
		pref = curr
		// 当前指针指向下一个
		curr = temp
	}
	return pref
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 计算数字
	num1 := Sum(l1)
	num2 := Sum(l2)
	num := num1 + num2
	// 构建链表
	return lreateList(num)
}
func lreateList(num int) *ListNode {
	var curr *ListNode
	// 转化字符串
	n := strconv.Itoa(num)
	for i := 0; i < len(n); i++ {
		// 转化为数字
		atoi, _ := strconv.Atoi(string(n[i]))
		// 头插法构建链表
		t := &ListNode{
			Val:  atoi,
			Next: curr,
		}
		curr = t
	}
	return curr
}

func Sum(l *ListNode) int {
	res := 0
	n := 1
	for l != nil {
		res += l.Val * n
		n = n * 10
		l = l.Next
	}
	return res
}


