func validPalindrome(s string) bool {
	runes := []rune(s)
	n := len(runes)

	if isPal(runes, -1) {
		return true
	}

	for i:=0; i<n;i++ {
		
		if isPal(runes, i) {
			return true
		}
	}

	return false
}

func isPal(runes []rune, skipPos int) bool{
	left, right := 0, len(runes) - 1

	for left < right {

		if skipPos >= 0 {

			if left == skipPos {
				left ++
			}

			if right == skipPos {
				right --
			}

		}
		
		if runes[left] != runes[right] {
			return false

		}

		left ++
		right --
	}

	return true
}
