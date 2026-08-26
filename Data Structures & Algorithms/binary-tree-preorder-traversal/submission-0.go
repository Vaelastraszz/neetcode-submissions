/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    
	res := []int{}

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode){
		if root == nil {
			return
		}
		
		res = append(res, root.Val)
		
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return res
}
