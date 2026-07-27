func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	var sig string
	sigMap := make(map[string][]string)

	for _, str := range strs {
		sig = getSignature(str)
		if wordList, ok := sigMap[sig]; ok {
		sigMap[sig] = append(wordList, str)
			continue
		}
		sigMap[sig] = []string{str}
	}

	for _, wordList := range sigMap {
		result = append(result, wordList)
	}
	
	return result
}

func getSignature(s string) string {
	chars := []byte(s)

	sort.Slice(chars, func(i,j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}