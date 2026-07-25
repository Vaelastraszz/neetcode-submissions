func twoSum(nums []int, target int) []int {
    seenNum := make(map[int]int, len(nums))

	for i, num := range nums {
		if pos_first_num, ok := seenNum[target-num]; ok {
			return []int{pos_first_num, i}
		} else{
			seenNum[num] = i
		}
	}
	return nil
}
