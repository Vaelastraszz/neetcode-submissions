func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums) - 1
	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	maxMoneyFirstPath := dp[n-1]

	n = len(nums) - 1
	dp = make([]int, n)

	dp[0] = nums[1]
	dp[1] = max(nums[1], nums[2])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i+1]+dp[i-2])
	}

	maxMoneySecondPath := dp[n-1]

	return max(maxMoneyFirstPath, maxMoneySecondPath)
}