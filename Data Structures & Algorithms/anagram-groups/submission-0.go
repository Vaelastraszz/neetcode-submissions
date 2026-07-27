func groupAnagrams(strs []string) [][]string {
	hasAnagrams := make(map[string]struct{})
	result := [][]string{}

	for i, word := range strs {
		if _, ok := hasAnagrams[word];ok {
			continue
		}
		anagSlice := []string{}
		anagSlice = append(anagSlice, word)
		hasAnagrams[word] = struct{}{}
		
		for _, word_2 := range strs[i+1:] {
			if ok := isAnagram(word, word_2); ok {
				hasAnagrams[word_2] = struct{}{}
				anagSlice = append(anagSlice, word_2)
			} 
		}
		result = append(result, anagSlice)
	}
	return result
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[rune]int)

	for _, c := range str1 {
		count[c]++
	}

	for _, c := range str2 {
		count[c]--

		if count[c] < 0 {
			return false
		}
	}

	return true
}


