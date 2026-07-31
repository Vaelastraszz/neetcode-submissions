func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1
	result := make([]int,0, 2)

	for left <= right {
		if numbers[left] + numbers[right] < target {
			left ++
			continue
		}

		if numbers[left] + numbers[right] > target {
			right --
			continue
		}

		result = append(result, left+1, right+1)
		break
	}

	return result
}
