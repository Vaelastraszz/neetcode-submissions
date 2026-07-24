func getConcatenation(nums []int) []int {
    ans := make([]int, 2 * len(nums))
	pos_ins := 0

	for i := 0; i < 2; i ++ {	
		for j := range nums {
			ans[pos_ins + j] = nums[j]
		}
		pos_ins = len(nums)
	}
	return ans
}
