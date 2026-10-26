type StockSpanner struct {
	prices []Price
}

type Price struct {
	quote int
	span  int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: []Price{},
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	for i := len(this.prices) - 1; i >= 0; {
		if this.prices[i].quote > price {
			break
		}

		span += this.prices[i].span
		i -= this.prices[i].span
	}

	this.prices = append(this.prices, Price{
		quote: price,
		span:  span,
	})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 