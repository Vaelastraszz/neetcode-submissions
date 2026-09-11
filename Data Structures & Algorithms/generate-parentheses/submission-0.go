func generateParenthesis(n int) []string {
	result := []string{}
	current := make([]rune, 0, 2*n-2)

	fullSlice := make([]rune, 2*n)
	fullSlice[0], fullSlice[2*n-1] = '(', ')'

	parenthesisLeft := map[rune]int{
		'(': n - 1,
		')': n - 1,
	}

	open := 1
	close := 0

	var backtrack func()

	backtrack = func() {

		if len(current) == n*2-2 {

			for i, validPar := range current {
				fullSlice[i+1] = validPar
			}

			validStr := string(fullSlice)
			result = append(result, validStr)
			return
		}

		for par, nbLeft := range parenthesisLeft {

			if nbLeft > 0 {

				if par == ')' && close >= open {
					continue
				}

				current = append(current, par)
				parenthesisLeft[par]--

				if par == '(' {
					open++
				} else {
					close++
				}

				backtrack()

				parenthesisLeft[current[len(current)-1]]++

				if par == '(' {
					open--
				} else {
					close--
				}

				current = current[:len(current)-1]
			}
		}
	}

	backtrack()
	return result
}