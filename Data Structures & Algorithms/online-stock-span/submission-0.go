type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	
	span := 1
	
	if len(this.prices) > 0 {
		end := len(this.prices) - 1

		for i := end ; i >= 0; i -- {
			if this.prices[i] <= price {
				span ++
			} else {
				break
			}
		}
	}

	this.prices = append(this.prices, price)

	return span

}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 