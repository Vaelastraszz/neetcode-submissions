func search(nums []int, target int) int {
	
	offset := 0
	res := -1

	var bs func([]int, int) 

	bs = func(nums []int, target int) {
		
		if len(nums) == 0 {
			return
		}
		
		left, right := 0, len(nums) - 1
		mid := (left+right)/2

		if nums[mid] == target {
			res = offset + mid
			return
		}

		if target < nums[mid] {
			bs(nums[:mid], target)
		} else {
			offset += mid + 1
			bs(nums[mid+1:], target)
		}


	}

	bs(nums, target)

	return res

}
