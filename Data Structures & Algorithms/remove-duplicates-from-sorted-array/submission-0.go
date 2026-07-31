func removeDuplicates(nums []int) int {
	left, right := 0,1

	for right < len(nums) {
						
		if nums[left] == nums[right] {
			right ++
			continue
		}

		if right > left + 1 {
			nums[left + 1] = nums[right]
		}

		left ++
		right ++
	}
	nums = nums[:left+1]
	return len(nums)
}
