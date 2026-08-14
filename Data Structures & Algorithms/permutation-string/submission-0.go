func checkInclusion(s1 string, s2 string) bool {
	
	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	windowSize := len(s1Rune)

	maps1 := make(map[rune]int)
	maps2 := make(map[rune]int)

	left := 0

	for _, letter := range s1Rune {
		maps1[letter] ++
	}

	outerLoop : for right := (left + windowSize - 1); right < len(s2Rune); right ++ {
		
		if right - left + 1 < windowSize {
			return false
		}
		
		for k, v := range maps1 {
			maps2[k] = v
		}

		for _, letter := range s2Rune[left:right+1] {
			
			maps2[letter] --

			if maps2[letter] < 0 {
				clear(maps2)
				left ++
				continue outerLoop
			}


		}

		return true

	}

	return false
}
