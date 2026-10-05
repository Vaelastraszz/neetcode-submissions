func maxProduct(nums []int) int {
    
	n := len(nums)
	dpMin := make([]int, n)
	dpMax := make([]int, n)

	if len(nums) == 1 {
		return nums[0]
	}

	dpMin[0], dpMax[0] = nums[0], nums[0]
	var answer int

	for i:= 1; i<n; i++ {
		
		newMin := min(dpMin[i-1]*nums[i],dpMax[i-1]*nums[i])
		newMax := max(dpMin[i-1]*nums[i],dpMax[i-1]*nums[i])
		
		dpMin[i] = min(newMin, nums[i])
		dpMax[i] = max(newMax, nums[i])

		answer = max(answer, dpMax[i])
	}

	return answer
}
