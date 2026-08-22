/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
		
	if root == nil {
		return true
	}
	
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if abs(leftDepth - rightDepth) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)

	}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func maxDepth(root *TreeNode) int {
	
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}

