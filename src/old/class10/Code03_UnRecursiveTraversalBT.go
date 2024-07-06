package class10

import (
	"fmt"
	"github.com/jiangbaiyan/go/src/old/class03"
)

// PreU 非递归先序遍历
func PreU(head *TreeNode) {
	if head == nil {
		return
	}
	stack := class03.MyStack{}
	stack.Push(head)
	for !stack.IsEmpty() {
		node := stack.Pop().(*TreeNode)
		fmt.Println(node.value)
		// 先压右再压左
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
	return
}

// PosU 非递归后序遍历
func PosU(head *TreeNode) {
	if head == nil {
		return
	}
	// 准备两个栈
	s1 := class03.MyStack{}
	s2 := class03.MyStack{}
	s1.Push(head)
	for !s1.IsEmpty() {
		// 头 右 左 反过来就是 左 右 头 => 后序
		node := s1.Pop().(*TreeNode)
		s2.Push(node)
		// 先压左再压右
		if node.left != nil {
			s1.Push(node.left)
		}
		if node.right != nil {
			s1.Push(node.right)
		}
	}
	for !s2.IsEmpty() {
		fmt.Println(s2.Pop().(*TreeNode).value)
	}
	return
}

// InU 非递归中序遍历
func InU(cur *TreeNode) {
	if cur == nil {
		return
	}
	stack := class03.MyStack{}
	for cur != nil {
		stack.Push(cur)
		cur = cur.left
	}
	for !stack.IsEmpty() {
		node := stack.Pop().(*TreeNode)
		fmt.Println(node.value)
		node = node.right
		for node != nil {
			stack.Push(node)
			node = node.left
		}
	}
	return
}
