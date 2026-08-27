/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	
	if root == nil {
		return false
	}

	var equalDFS func(root, subRoot *TreeNode) bool

	equalDFS = func(root, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}

		if root == nil || subRoot == nil {
			return false
		}

		if root.Val != subRoot.Val {
			return false
		}

		return equalDFS(root.Left, subRoot.Left) &&
			equalDFS(root.Right, subRoot.Right)
	}

	if equalDFS(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}