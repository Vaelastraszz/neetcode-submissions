func longestPalindromeSubseq(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for seqSize := 1; seqSize <= n; seqSize++ {
		l := 0
		r := l + seqSize - 1

		for r < n {
			if seqSize == 1 {
				dp[l][r] = 1
			} else if s[l] == s[r] {
				if seqSize == 2 {
					dp[l][r] = 2
				} else {
					dp[l][r] = 2 + dp[l+1][r-1]
				}
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}

			l++
			r++
		}
	}

	return dp[0][n-1]
}