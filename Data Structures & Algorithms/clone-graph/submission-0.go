/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]bool)
	copyMap := make(map[int]*Node)

	queue := []*Node{node}
	visited[node.Val] = true

	for len(queue) > 0 {
		vertice := queue[0]
		queue = queue[1:]

		copyVertice, ok := copyMap[vertice.Val]

		if !ok {
			copyVertice = &Node{
				Val: vertice.Val,
			}
			copyMap[vertice.Val] = copyVertice
		}

		for _, neighbor := range vertice.Neighbors {
			copyNeighbor, ok := copyMap[neighbor.Val]

			if !ok {
				copyNeighbor = &Node{
					Val: neighbor.Val,
				}
				copyMap[neighbor.Val] = copyNeighbor
			}

			copyVertice.Neighbors = append(
				copyVertice.Neighbors,
				copyNeighbor,
			)

			if !visited[neighbor.Val] {
				visited[neighbor.Val] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return copyMap[node.Val]
}