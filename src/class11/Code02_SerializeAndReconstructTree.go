package class11

// PreMarshal 先序方式序列化
func PreMarshal(head *Node) {
	ans := make([]interface{}, 0)
	preMarshal(head, ans)
}

func preMarshal(head *Node, ans []interface{}) {
	if head == nil {
		ans = append(ans, nil)
		return
	}
	ans = append(ans, head.value)
	preMarshal(head.left, ans)
	preMarshal(head.right, ans)
}

// PreUnmarshal 先序方式反序列化
func PreUnmarshal(preList []interface{}) *Node {
	if len(preList) == 0 {
		return nil
	}
	return preUnmarshal(preList)
}

func preUnmarshal(preList []interface{}) *Node {
	node, preList := preList[0], preList[1:]
	if node == nil {
		return nil
	}
	head := &Node{value: node}
	head.left = preUnmarshal(preList)
	head.right = preUnmarshal(preList)
	return head
}
