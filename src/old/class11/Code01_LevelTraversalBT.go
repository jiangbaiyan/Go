package class11

import (
	"fmt"
	"github.com/jiangbaiyan/go/src/old/class03"
)

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

// LevelTraversalBT 层序遍历
func LevelTraversalBT(head *Node) {
	if head == nil {
		return
	}
	queue := &class03.MyRingQueue{}
	queue.Push(head)
	for !queue.IsEmpty() {
		cur := queue.Pop().(*Node)
		fmt.Println(cur.value)
		if cur.left != nil {
			queue.Push(cur.left)
		}
		if cur.right != nil {
			queue.Push(cur.right)
		}
	}
}
