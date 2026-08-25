/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	curr *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	curr := root
	stack := []*TreeNode{}

	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	return BSTIterator{
		curr: curr, 
		stack: stack,
	}
}

func (this *BSTIterator) Next() int {
	
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	val := node.Val

	this.curr = node.Right

	for this.curr != nil {
		this.stack = append(this.stack, this.curr)
		this.curr = this.curr.Left
	}

	return val
}

func (this *BSTIterator) HasNext() bool {
	if len(this.stack) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 