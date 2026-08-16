func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	curr := list2
	head := list1

	compareNode := list1

	var temp, prev *ListNode

	for curr != nil {

		if curr.Val < compareNode.Val {

			temp = curr.Next
			curr.Next = compareNode

			if curr.Next == head {
				head = curr
			} else {
				prev.Next = curr
			}

			prev = curr
			curr = temp
			continue
		}

		prev = compareNode

		if compareNode.Next == nil {
			compareNode.Next = curr
			break
		}

		compareNode = compareNode.Next
	}

	return head
}