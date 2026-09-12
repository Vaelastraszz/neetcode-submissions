func partition(s string) [][]string {
	
	result := [][]string{}
	part := []string{}

	var backtrack func(start, end int)

	backtrack = func(start, end int) {
		if start >= len(s) {
			result = append(result, append([]string{}, part...))
			return
		}
		if end >= len(s) {
			return
		}

		if isPali(s, start, end) {
			part = append(part, s[start:end+1])
			backtrack(end+1, end + 1)
			part = part[:len(part)-1]
		}

		backtrack(start, end+1)
	}

	backtrack(0,0)
	return result
}

func isPali(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}