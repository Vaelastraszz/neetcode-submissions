func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	left := 0
	length := 0
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

		delete(mapDuplicates, runes[left])
		left++

		mapDuplicates[letter] = true

		length = max(length, len(mapDuplicates))
	}

	return length
}