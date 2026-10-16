/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    

	maxHeight := 0

	var computeHeight func(root *TreeNode) int 
	
	computeHeight = func(root *TreeNode) int {
		
		if root == nil {
			return 0
		}

		leftH := computeHeight(root.Left)
		rightH := computeHeight(root.Right)
		maxHeight = max(maxHeight, leftH + rightH)

		return 1 + max(leftH, rightH)
	}

	computeHeight(root)
	return maxHeight
}
