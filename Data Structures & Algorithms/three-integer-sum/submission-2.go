func threeSum(nums []int) [][]int {
	
	sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j] })

	result := make([][]int,0)

	for i := range nums {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		for left, right := i + 1, len(nums) - 1; right > left; {
						
			if nums[i] + nums[left] + nums[right] < 0 {
				left ++
				continue
			}

			if nums[i] + nums[left] + nums[right] > 0 {
				right --
				continue
			}
				
				good_sum := []int{nums[i],nums[left],nums[right]}
				result = append(result, good_sum)
				
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				
				right -- 
				for right > left && nums[right] == nums[right+1] {
					right--
				}		
	}

	}

	return result

}
