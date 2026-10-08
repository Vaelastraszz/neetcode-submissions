func longestConsecutive(nums []int) int {
	
	n := len(nums)
	
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	mapNums := make(map[int]bool)
	longSeq := 1
	var streak int

	for _, num := range nums {
		mapNums[num] = true
	}

	for num := range mapNums {
		
		streak = 1
		_, isNotFirst := mapNums[num-1]

		if !isNotFirst {		
			for i:=num+1; ;i ++ {
				if _, ok := mapNums[i]; ok {
					streak ++
					continue
				} 
				break
			}
		}

		longSeq = max(streak, longSeq)
	}

	return longSeq
}
