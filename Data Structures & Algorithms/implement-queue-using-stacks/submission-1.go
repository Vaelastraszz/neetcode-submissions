type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			n := len(this.stack1) - 1
			val := this.stack1[n]

			this.stack1 = this.stack1[:n]
			this.stack2 = append(this.stack2, val)
		}
	}

	n := len(this.stack2) - 1
	val := this.stack2[n]
	this.stack2 = this.stack2[:n]

	return val
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			n := len(this.stack1) - 1
			val := this.stack1[n]

			this.stack1 = this.stack1[:n]
			this.stack2 = append(this.stack2, val)
		}
	}

	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}