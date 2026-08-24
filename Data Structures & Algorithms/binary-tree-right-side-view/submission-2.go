/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type LevelNode struct {
	*TreeNode
	level int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := []*LevelNode{
		{
			TreeNode: root,
			level:    0,
		},
	}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if len(result) == node.level {
				result = append(result, node.Val)
			}

			if node.Right != nil {
				queue = append(queue, &LevelNode{
					TreeNode: node.Right,
					level:    node.level + 1,
				})
			}

			if node.Left != nil {
				queue = append(queue, &LevelNode{
					TreeNode: node.Left,
					level:    node.level + 1,
				})
			}
		}
	}

	return result
}
