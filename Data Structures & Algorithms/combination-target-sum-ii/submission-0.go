func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}

	sortCand := append([]int{}, candidates...)

	sort.Slice(sortCand, func(i, j int) bool {
		return sortCand[i] < sortCand[j]
	})

	var backtrack func(int, int)

	backtrack = func(start int, sum int) {

		if sum == target {
			combinations := append([]int{}, current...)
			result = append(result, combinations)
			return
		}

		if sum > target || start >= len(sortCand) {
			return
		}

		for i := start; i < len(sortCand); i++ {

			if i > start && sortCand[i] == sortCand[i-1] {
				continue
			}

			if sum+sortCand[i] > target {
				break
			}

			current = append(current, sortCand[i])

			backtrack(i+1, sum+sortCand[i])

			current = current[:len(current)-1]
		}
	}

	backtrack(0, 0)

	return result
}