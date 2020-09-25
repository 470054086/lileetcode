package lettlink

/**
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。

 

示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-lists-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	r := &ListNode{}
	result := r
	// 存放一个栈
	stack := make([]int, 1)
	var tmp int
	// 进行数组循环 两个数组都要循环到最后
	for l1 != nil || l2 != nil {
		if len(stack) == 0 {
			tmp = 0
		} else {
			tmp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		// 循环第一个链表
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		// 循环第二个链表
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		// 构造链表
		result.Next = &ListNode{
			Val: tmp % 10,
		}
		// 进位 存栈
		if tmp/10 > 0 {
			stack = append(stack, tmp/10)
		}
		// 向前移动
		result = result.Next
	}
	// 如果栈里面还有数据的话
	if len(stack) > 0 {
		result.Next = &ListNode{
			Val: stack[len(stack)-1],
		}
	}
	return r.Next
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	r := &ListNode{}
	// 存储一个最终值得链表
	resStack := []int{}
	// 进行数组循环 两个数组都要循环到最后
	for l1 != nil || l2 != nil {
		tmp := 0
		// 循环第一个链表
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		// 循环第二个链表
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		// 如果大于10 向上进位
		if tmp >= 10 {
			resStack[len(resStack)-1] = resStack[len(resStack)-1] + 1
			tmp = tmp % 10
		}
		// 存储当前的值
		resStack = append(resStack, tmp)
	}
	// 生成链表
	new := r
	for i := 0; i < len(resStack); i++ {
		new.Next = &ListNode{
			Val: resStack[i],
		}
		new= new.Next
	}

	return r.Next
}
