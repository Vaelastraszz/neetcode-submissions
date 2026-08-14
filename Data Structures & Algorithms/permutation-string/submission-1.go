func checkInclusion(s1 string, s2 string) bool {
	
	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	windowSize := len(s1Rune)

	if windowSize > len(s2Rune) {
		return false
	}

	maps1 := make(map[rune]int)
	maps2 := make(map[rune]int)

	for _, letter := range s1Rune {
		maps1[letter]++
	}

	left := 0

	for right := 0; right < len(s2Rune); right++ {
		// Add the new character to the window
		maps2[s2Rune[right]]++

		// Keep the window size equal to s1
		if right-left+1 > windowSize {
			maps2[s2Rune[left]]--
			left++
		}

		// Check if the two frequency maps are equal
		if right-left+1 == windowSize {
			match := true

			for k, v := range maps1 {
				if maps2[k] != v {
					match = false
					break
				}
			}

			if match {
				return true
			}
		}
	}

	return false
}