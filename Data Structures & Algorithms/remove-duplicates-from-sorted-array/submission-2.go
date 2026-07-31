func removeDuplicates(nums []int) int {
	left, right := 0,1

	for right < len(nums) {
						
		if nums[left] == nums[right] {
			right ++
			continue
		}

		nums[left + 1] = nums[right]
		left ++
		right ++
	}
	
	return left + 1 
}
