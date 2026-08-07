func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	var mid, valRight, valMid int


	for left != right {

		mid = (left + right) / 2
		valRight = nums[right]
		valMid = nums[mid]

		if valMid > valRight {
			left = mid + 1
		} else if valMid < valRight {
			right = mid
		}

	}
	return nums[left]
}
