package tree

import "fmt"

func GetNetxt(node *TreeNode) {
	// 先找到节点的跟节点
	p := node
	// 找到的节点为根节点
	for p != nil {
		p = p.Next
	}
	MiddOrder(p)
	// 获取到所有的节点了 中序节点的下一个节点

}

func MiddOrder(p *TreeNode) {
	if p != nil {
		MiddOrder(p.Left)
		fmt.Println("我是当前的节点", p)
		MiddOrder(p.Right)
	}
}
