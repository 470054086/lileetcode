package main

import (
	"fmt"
	"lileetcode.com/tree"
)

func main() {
	t := &tree.TreeNode{
		Val:3,
		Left:&tree.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
			Next:  nil,
		},
		Right:&tree.TreeNode{
			Val:   20,
			Left:  &tree.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &tree.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next:  nil,
		},
	}
	//palindrome := tree.MirrorTreeStack(t)
	////palindrome1 := tree.MirrorTree(t)
	queue := tree.LeftRightQueue(t)
	fmt.Println(queue)
	//fmt.Println()
	//poster(palindrome1)
}

func poster(n *tree.TreeNode) {
	// 定义一个队列
	stack := []*tree.TreeNode{}
	stack = append(stack, n)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
}
