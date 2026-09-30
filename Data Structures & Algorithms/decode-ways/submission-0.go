func numDecodings(s string) int {
	runes := []rune(s)
	n := len(runes)

	if n == 0 {
		return 0
	}

	if n == 1 {
		if runes[0] == '0' {
			return 0
		}
		return 1
	}

	dp := make([]int, n)

	// Premier caractère
	if runes[0] != '0' {
		dp[0] = 1
	}

	// Deux premiers caractères
	num := int(runes[0]-'0')*10 + int(runes[1]-'0')

	if runes[1] != '0' {
		dp[1] += dp[0]
	}

	if num >= 10 && num <= 26 {
		dp[1] += 1
	}

	for i := 2; i < n; i++ {
		num := int(runes[i-1]-'0')*10 + int(runes[i]-'0')

		// Chiffre seul
		if runes[i] != '0' {
			dp[i] += dp[i-1]
		}

		// Deux chiffres
		if num >= 10 && num <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n-1]
}