func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	left, length := 0, 0
	mapDuplicates := make(map[rune]bool)

	for _, letter := range runes {

		if _, ok := mapDuplicates[letter]; !ok {
			mapDuplicates[letter] = true
			length = max(length, len(mapDuplicates))
			continue
		}

		for runes[left] != letter {
			delete(mapDuplicates, runes[left])
			left++
		}

		left++
	}

	return length
}