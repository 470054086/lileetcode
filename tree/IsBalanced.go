package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
判断平衡树
*/
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	IsBalanced(root.Left)
	IsBalanced(root.Right)
	if math.Abs(float64(depp(root.Left)-depp(root.Right))) <= 1 {
		return true
	}else{
		return false
	}
}

/**
计算树的深度
*/
func depp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depp(root.Left), depp(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
