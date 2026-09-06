func subsets(nums []int) [][]int {


	var backtrack func(start int)
	current := []int{}
	result := [][]int{}

	backtrack = func(start int) {
		
		subset := append([]int{}, current...)
		result = append(result, subset)

		for i := start; i < len(nums); i ++ {			
			current = append(current, nums[i])
			backtrack(i+1)
			current = current[:len(current)-1] 

		}
	}

	backtrack(0)

	return result

}
