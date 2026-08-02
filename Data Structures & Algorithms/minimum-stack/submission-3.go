type MinStack struct {
	data []int
	all_mins []int
}

func Constructor() MinStack {

	data := make([]int, 0)
	all_mins := make([]int, 0)

	return MinStack{data : data,
					all_mins : all_mins}
}

func (this *MinStack) Push(val int) {

	this.data = append(this.data, val)
	nb_mins := len(this.all_mins)

	if nb_mins == 0 || (nb_mins > 0 && val <= this.all_mins[nb_mins - 1]) {
		this.all_mins = append(this.all_mins, val)
		return
	} 

}

func (this *MinStack) Pop() {

	if len(this.data) > 0 {
		end := len(this.data) - 1

		end_min := len(this.all_mins)-1
		min := this.all_mins[end_min]

		val := this.data[end]
		this.data = this.data[:end]

		if min == val {
			this.all_mins = this.all_mins[:end_min]
		}

	}

}

func (this *MinStack) Top() int {
	
	if len(this.data) > 0 {
		
		end := len(this.data) - 1

		return this.data[end]
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.all_mins) > 0 {
		
		min := this.all_mins[len(this.all_mins)-1]
		return min
	}

	return 0
}
