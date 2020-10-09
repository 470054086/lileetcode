package main

import (
	"fmt"
	"lileetcode.com/tree"
)

func main() {
	t := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val: 11,
				Left: &tree.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &tree.TreeNode{
			Val: 8,
			Left: &tree.TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &tree.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	palindrome := tree.PathSum(t, 22)
	fmt.Println(palindrome)
}
