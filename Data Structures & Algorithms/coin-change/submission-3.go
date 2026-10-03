func coinChange(coins []int, amount int) int {
    
	dp := make([]int, amount+1)
	dp[0] = 0

	for i:= 1; i <= amount; i ++ {

		dp[i] = math.MaxInt

		for _, coin := range coins {
			
			if coin > i || dp[i-coin] == math.MaxInt {
				continue
			}

			dp[i] = min(dp[i-coin]+1, dp[i])

		}

	}

	if dp[amount] == math.MaxInt {
			return - 1
		}

	return dp[amount]
}
