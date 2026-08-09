func containsNearbyDuplicate(nums []int, k int) bool {
	left := 0 

	if k <= 0 {
		return false
	}

	for right:= left + 1 ; right < len(nums) ; right ++ {

		if right - left > k {
			left ++ 
		}

		for i:= left; i < right; i ++ {
			
			if nums[i] == nums[right] {
				return true
			}
		}

	}

	return false
}
