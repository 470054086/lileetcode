package lettlink

/**
编写一个函数，检查输入的链表是否是回文的。

 

示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 定义快慢指针
	left := head
	right := head
	var pref *ListNode

	// 快指针走到最后 即使慢指针的中间点  将链表反转即可
	for right != nil && right.Next != nil {
		// 前一个为当前指针
		temp := left
		left = left.Next
		right = right.Next.Next
		temp.Next = pref
		pref = temp
	}
	// 如果还没循环玩 说明回文串是奇数 不需要中间一个
	if right != nil {
		left = left.Next
	}

	for left != nil {
		if left.Val != pref.Val {
			return false
		}
		left = left.Next
		pref = pref.Next
	}
	return true
}
