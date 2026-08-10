func minSubArrayLen(target int, nums []int) int {
	left, total := 0, 0
	length := len(nums) + 1

	for right := 0; right < len(nums); right ++ {

		total += nums[right]

		if total >= target {
			for total >= target {
				length = min(length, (right - left) + 1 )
				total -= nums[left]
				left ++
			}
		}
	}

	if length > len(nums) {
		return 0
	} else {
		return length
	}
}
