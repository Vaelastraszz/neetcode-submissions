func copyRandomList(head *Node) *Node {
	curr := head

	var newHead, newPrev *Node
	nodes := make(map[*Node]*Node)

	for curr != nil {
		newNode := newNode(curr.Val)

		if newPrev != nil {
			newPrev.Next = newNode
		} else {
			newHead = newNode
		}

		newPrev = newNode
		nodes[curr] = newNode
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		if randomNode, ok := nodes[curr.Random]; ok {
			nodes[curr].Random = randomNode
		}

		curr = curr.Next
	}

	return newHead
}

func newNode(val int) *Node {
	return &Node{
		Val: val,
	}
}