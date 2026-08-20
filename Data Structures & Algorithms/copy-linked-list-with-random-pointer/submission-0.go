/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
	curr := head
	var headNew, node, prev *Node
	mapNode := make(map[*Node]*Node)
	

	for curr != nil {
		
		node = newNode(curr.Val)
		
		if prev != nil {
			prev.Next = node
		}
		
		if headNew == nil {
			headNew = node
		}
		
		prev = node
		mapNode[curr] = node
		curr = curr.Next
		
	}

	curr = head

	for curr != nil {

		if node, ok  := mapNode[curr.Random]; ok {
			mapNode[curr].Random = node
		} else {
			mapNode[curr].Random = nil
		}

		curr = curr.Next
	}

	return headNew

}

func newNode(val int) *Node {
	return &Node{
		Val : val,
	}
}
