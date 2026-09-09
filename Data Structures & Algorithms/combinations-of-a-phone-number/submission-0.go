func letterCombinations(digits string) []string {
	result := []string{}
	phoneMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	digitsRune := []rune(digits)

	current := []rune{}

	var backtrack func(int)

	backtrack = func(start int) {

		if len(current) == len(digitsRune) {
			result = append(result, string(current))
			return
		}

		letters, ok := phoneMap[digitsRune[start]]

		if ok {
			for _, letter := range letters {

				current = append(current, letter)
				backtrack(start+1)
				current = current[:len(current)-1]

			}

		}
	
	}

	backtrack(0)
	return result

}
