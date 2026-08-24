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
		&LevelNode{
			root,
			0,
			},
	}
	
	for len(queue) > 0 {

		for range queue {
			
			node := queue[0]
			queue = queue[1:]

			if result == nil || len(result)-1 < node.level {
				result = append(result, node.Val)
			}

			if node.Right != nil {
				queue = append(queue, &LevelNode{
					node.Right,
					node.level + 1,
				})
			}

			
			if node.Left != nil {
				queue = append(queue, &LevelNode{
					node.Left,
					node.level + 1,
				})
			}

		}

	}

	return result

}
