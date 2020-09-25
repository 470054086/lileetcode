package main

import (
	"fmt"
	"lileetcode.com/lettlink"
)

func main() {
	//l := &lettlink.ListNode{
	//	Val: 1,
	//	Next: &lettlink.ListNode{
	//		Val: 2,
	//		Next: &lettlink.ListNode{
	//			Val: 3,
	//			Next: &lettlink.ListNode{
	//				Val: 3,
	//				Next: &lettlink.ListNode{
	//					Val: 2,
	//					Next: &lettlink.ListNode{
	//						Val:  1,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	l := &lettlink.ListNode{
		Val: 1,
		Next: &lettlink.ListNode{
			Val:  2,
			Next: &lettlink.ListNode{
				Val:  3,
				Next:&lettlink.ListNode{
					Val:  2,
					Next: &lettlink.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	//r := &lettlink.ListNode{
	//	Val: 2,
	//	Next: &lettlink.ListNode{
	//		Val:  9,
	//		Next: &lettlink.ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//
	//}

	palindrome := lettlink.IsPalindrome(l)
	fmt.Println(palindrome)
}
