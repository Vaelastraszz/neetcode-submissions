/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	res []int
}

func Constructor(root *TreeNode) BSTIterator {

	stack := []*TreeNode{}
	curr := root
	res := []int{}

	var node *TreeNode

	for curr !=nil || len(stack) > 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)
		
		curr = node.Right
		
	}

	return BSTIterator{
		res : res,
	}

}

func (this *BSTIterator) Next() int {
	val := this.res[0]
	this.res = this.res[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	if len(this.res) == 0 {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 