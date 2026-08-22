/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {

	result := []int{}
	stack := []*TreeNode{}
	current := root

	var node *TreeNode

	for current != nil || len(stack) > 0 {

		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		current = node.Right
		

	}

	return result

}
