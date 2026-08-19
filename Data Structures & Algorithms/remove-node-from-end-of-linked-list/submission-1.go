/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	slow, fast, prev := head, head, head
	var start bool
	
	for i:= 1; i <= n; i++ {
		fast = fast.Next
		if fast == nil {
			break
		}
	}

	for fast != nil {
		
		if start {
			prev = prev.Next
		}

		fast = fast.Next
		slow = slow.Next

		if !start {
			start = true
		}
	}

	if slow != head {

		prev.Next = slow.Next
		slow.Next = nil

	} else {
		head = slow.Next
		slow.Next = nil
	}

	return head

}
