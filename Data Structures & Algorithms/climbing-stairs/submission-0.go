func climbStairs(n int) int {
    if n <= 2 {
		return n
	}

	dp := [2]int{1,2}

	for i:=1; i <= n - 2; i ++ {
		tmp := dp[1]
		dp[1] = dp[1] + dp[0]
		dp[0] = tmp
	}

	return dp[1] 
}
