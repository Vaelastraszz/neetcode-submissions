func topKFrequent(nums []int, k int) []int {
	
	freqMap := make(map[int]int, len(nums))
	bucket := make([][]int, len(nums)+1)
	result := []int{}

	for _, num := range nums {
		freqMap[num]++
	}

	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	for i:=len(nums); i>=0; i -- {
		result = append(result, bucket[i]...)

		if len(result) >= k {
			return result[:k]
		} 
	}

	return result
}
