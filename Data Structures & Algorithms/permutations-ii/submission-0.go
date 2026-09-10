func permuteUnique(nums []int) [][]int {
	
	result := [][]int{}
	current := []int{}
	positionMap := make(map[int]bool)
	
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)

	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i] < sortedNums[j]
	})
	var backtrack func()

	backtrack = func() {

		if len(current) == len(nums) {
			permutations := append([]int{}, current...)
			result = append(result, permutations)
			return
		}

		for i:= 0 ; i < len(sortedNums); i++ {

			if _, ok := positionMap[i]; ok{
				continue
			}

			if i > 0 && sortedNums[i] == sortedNums[i-1] {
				    if _, used := positionMap[i-1]; !used {
       				 continue
    				}
			}

			current = append(current, sortedNums[i])
			positionMap[i] = true
			backtrack()
			delete(positionMap, i)
			current = current[:len(current)-1]
			
		}
	}

	backtrack()
	return result
}
