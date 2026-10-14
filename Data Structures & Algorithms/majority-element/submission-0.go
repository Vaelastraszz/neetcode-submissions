func majorityElement(nums []int) int {
    freqMap := make(map[int]int)
	n := len(nums)
	var majElem int

	for _, num := range nums {
		freqMap[num]++
	}

	for k,v := range freqMap {
		if v > n / 2 {
			majElem = k
		}
	}

	return majElem
}
