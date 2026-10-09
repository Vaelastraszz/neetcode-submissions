func maxProfit(prices []int) int {

	n := len(prices)
	dpHold := make([]int, n)
	dpCash := make([]int, n)

	dpHold[0] = -prices[0]
	dpCash[0] = 0

	for i:=1; i<n; i ++ {
		dpHold[i] = max(dpHold[i-1], dpCash[i-1] - prices[i])
		dpCash[i] = max(dpCash[i-1], dpHold[i-1] + prices[i])
	}

	return dpCash[n-1]
}
