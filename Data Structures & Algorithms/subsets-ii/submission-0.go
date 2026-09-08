func subsetsWithDup(nums []int) [][]int {

	result := [][]int{}
	current := []int{}
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)

	sort.Slice(sortedNums, func(i,j int) bool {return sortedNums[i] < sortedNums[j]})

	var backtrack func(int)

	backtrack = func(start int) {
		if start >= len(nums) {
			subset := append([]int{}, current...)
			result = append(result, subset)
			return
		}

		current = append(current, sortedNums[start])
		backtrack(start+1)
		current = current[:len(current)-1]

		for start + 1 < len(sortedNums) && sortedNums[start] == sortedNums[start+1] {
			start ++ 
		}

		backtrack(start+1)
	}

	backtrack(0)

	return result
}
