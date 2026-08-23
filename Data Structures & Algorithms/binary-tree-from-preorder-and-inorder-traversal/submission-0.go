/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    
	if len(preorder) == 0 || len(inorder) == 0 {
    return nil
	}
	
	rootVal := preorder[0]
	rootIndex := 0

	for i, val := range inorder {
		
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	rootNode := &TreeNode{
		Val: rootVal,
	}

	inordLeftPart := inorder[:rootIndex]
	inordRightPart := inorder[rootIndex+1:]

	preOrdLeft := preorder[1:1+len(inordLeftPart)]
	preOrdRight := preorder[1+len(inordLeftPart):]

	rootNode.Left = buildTree(preOrdLeft, inordLeftPart)
	rootNode.Right = buildTree(preOrdRight, inordRightPart)

	return rootNode
}
