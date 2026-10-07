func lengthOfLIS(nums []int) int {
    /*
	1. Qu'est-ce que dp[i] représente ?
	2. Si je suis à i, d'où puis-je venir ?
	3. Quelles sont toutes les possibilités ?
	4. Est-ce qu'un seul état suffit ?
	5. Quel est le cas de base ? 
	*/



	n := len(nums)

	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1
	maxLen := 1

	for i:= 1; i<n; i++ {
		dp[i] = 1
		for j:=0; j<i; j ++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxLen = max(maxLen,dp[i])
	}

	return maxLen
}
