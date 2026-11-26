/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getLonelyNodes(root *TreeNode) []int {
    
	// Creation de liste des resultats
	results := []int{}

	// On peut faire une dfs pour chaque noeud
	// si nil on return
	// Si il a un enfant non nil et un autre nil on ajoute le node non nil

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		
		if root == nil {
			return
		}

		if root.Left != nil && root.Right == nil {
			results = append(results, root.Left.Val)
		}

		if root.Right != nil && root.Left == nil {
			results = append(results, root.Right.Val)
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	
	dfs(root)
	
	return results
}
