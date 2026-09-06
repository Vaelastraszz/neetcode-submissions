func combinationSum(nums []int, target int) [][]int {
    
	result := [][]int{}
	subRes := []int{}

	var dfs func(int, int)

	dfs = func(i int, sum int) {
		
		if sum == target {
			combinations := append([]int{}, subRes...)
			result = append(result, combinations)
			return
		}

		if i == len(nums) || sum > target {
			return 
		}

		subRes = append(subRes, nums[i])
		dfs(i, nums[i]+sum)
		subRes = subRes[:len(subRes)-1]

		dfs(i+1, sum)
	}

	dfs(0,0)

	return result
}
