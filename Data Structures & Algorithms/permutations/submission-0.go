func permute(nums []int) [][]int {

	result := [][]int{}
	current := []int{}
	visited := make(map[int]bool)

	var backtrack func()

	backtrack = func() {
		
		if len(current) == len(nums) {
			permutations := append([]int{}, current...)
			result = append(result, permutations)
			return
		}

		for i:=0; i < len(nums); i ++ {

			if _, ok := visited[nums[i]]; ok {
				continue
			}

			current = append(current, nums[i])
			visited[nums[i]] = true
			backtrack()
			delete(visited, current[len(current)-1])
			current = current[:len(current)-1]
		}
	}

	backtrack()
	return result
}
