type MyStack struct {
	data []int
}

func Constructor() MyStack {
	mydata := []int{}
	return MyStack{
		data: mydata,
	}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]

	return val
}

func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
