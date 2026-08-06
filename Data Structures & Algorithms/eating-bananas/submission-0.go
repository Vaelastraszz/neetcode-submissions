func minEatingSpeed(piles []int, h int) int {

	max := 0

	for _, stackSize := range piles {

		if stackSize > max {
			max = stackSize
		}
	}

	left, right := 1, max
	var mid int 
	var hours int
	var k int


	for left <= right {
		mid = (left + right)/ 2
		hours = 0

		for _, stackSize := range piles {
			hours += (mid + stackSize -1) / mid
		}

		if hours > h {
			left = mid + 1

		} else if hours <= h {
			k = mid
			right = mid - 1
		} 

	}

	return k	

}
