/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var dfsGood func(*TreeNode, int) int

    dfsGood = func(node *TreeNode, maxSoFar int) int {
        if node == nil {
            return 0
        }

        if node.Val < maxSoFar {
            return dfsGood(node.Left, maxSoFar) +
                dfsGood(node.Right, maxSoFar)
        }

        return 1 +
            dfsGood(node.Left, node.Val) +
            dfsGood(node.Right, node.Val)
    }

    return dfsGood(root, root.Val)
}
