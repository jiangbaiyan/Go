package class10

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value interface{}
}

// Pre 先序
func Pre(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	Pre(head.left)
	Pre(head.right)
}

// In 中序
func In(head *TreeNode) {
	if head == nil {
		return
	}
	In(head.left)
	fmt.Println(head.value)
	In(head.right)
}

// Pos 后序
func Pos(head *TreeNode) {
	if head == nil {
		return
	}
	Pos(head.left)
	Pos(head.right)
	fmt.Println(head.value)
}
