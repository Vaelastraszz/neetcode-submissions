func combinationSum(nums []int, target int) [][]int {

   result := [][]int{}
   current := []int{}

   var backtrack func(int, int)
   backtrack = func(start, sum int) {
	
	if sum == target {
		comb := append([]int{}, current...)
		result = append(result, comb)
		return
	}

	if sum > target {
		return
	}

	for i:= start; i < len(nums); i++ {
		
		current = append(current, nums[i])
		backtrack(i, sum + nums[i])
		current = current[:len(current)-1]
	}

   }

   backtrack(0, 0)
   return result
}
