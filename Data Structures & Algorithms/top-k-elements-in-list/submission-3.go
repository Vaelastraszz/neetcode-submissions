func topKFrequent(nums []int, k int) []int {
	
	freqMap := make(map[int]int, len(nums))
	bucketMap := make(map[int][]int)
	result := []int{}

	for _, num := range nums {
		freqMap[num]++
	}

	for num, freq := range freqMap {
		bucketMap[freq] = append(bucketMap[freq], num)
	}

	for i:=len(nums); i>=0; i -- {
		if top_k, ok := bucketMap[i] ; ok {
			result = append(result, top_k...)
		}

		if len(result) >= k {
			return result[:k]
		} 
	}

	return result
}
