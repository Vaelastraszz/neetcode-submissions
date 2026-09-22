func maxArea(heights []int) int {
	
	maxWater := 0
	left, right := 0, len(heights) - 1

	for left < right {
		minHeight := min(heights[left], heights[right])
		maxWater = max(maxWater, minHeight * (right-left))

		if heights[left] < heights[right] {
			left ++
			continue
		}

		right -- 
	}

	return maxWater

}
