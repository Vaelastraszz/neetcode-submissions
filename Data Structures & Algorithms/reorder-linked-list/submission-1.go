func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    middle, fast := head, head.Next

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        middle = middle.Next
    }

    current := middle.Next
    var node, next *ListNode

    for current.Next != nil {
        next = middle.Next
        middle.Next = current.Next
        node = middle.Next
        current.Next = node.Next
        node.Next = next
    }

	node = middle.Next
	middle.Next = nil
	current = head

	for node != nil {
		next = node.Next

		node.Next = current.Next
		current.Next = node

		current = node.Next
		node = next
	}

}