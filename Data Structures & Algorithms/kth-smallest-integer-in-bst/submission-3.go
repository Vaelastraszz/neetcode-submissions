/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    cnt, res := k, 0

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if cnt == 0 {
			return
		}

		dfs(node.Left)
		cnt --

		if cnt == 0 {
			res = node.Val
		}

		dfs(node.Right)
	}

	dfs(root)

	return res
}
