func maxProfit(prices []int) int {
	price := prices[0]
	total := 0
	for _, p := range prices[1:] {
		if p > price {
			total += p - price
			price = p
		} else if p < price {
			price = p
		}
	}
	return total
}