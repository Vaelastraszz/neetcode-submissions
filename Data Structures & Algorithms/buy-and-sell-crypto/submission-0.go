func maxProfit(prices []int) int {
	left := 0
	profit := 0

	for right := left + 1; right < len(prices); right ++ {

		if prices[right] > prices[left] {
			profit = max(profit, prices[right] - prices[left])
		} else {
			left = right
		}
	}

	return profit
	
	}
