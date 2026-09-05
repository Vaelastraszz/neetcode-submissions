/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {

	var cumulSum func(root *TreeNode, targetSum, sum int) bool 

	cumulSum = func(root *TreeNode, targetSum, sum int) bool {
		
		if root == nil {
			return false
		}

		newSum := sum + root.Val

		if root.Left == nil && root.Right == nil && newSum == targetSum {
			return true
		}

		if cumulSum(root.Left, targetSum, newSum) {
			return true
		}

		if cumulSum(root.Right, targetSum, newSum) {
			return true
		}

		return false
	}

	return cumulSum(root, targetSum, 0)
}
