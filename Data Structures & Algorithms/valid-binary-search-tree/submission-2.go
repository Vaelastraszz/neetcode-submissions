/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

	var maxMinDFS func(*TreeNode, int, int) bool

	maxMinDFS = func(root *TreeNode, min, max int) bool {

		if root == nil {
			return true
		}

		if root.Val <= min || root.Val >= max {
			return false
		}

		return maxMinDFS(root.Left, min, root.Val) &&
			maxMinDFS(root.Right, root.Val, max)
	}

	return maxMinDFS(root, math.MinInt, math.MaxInt)
}
