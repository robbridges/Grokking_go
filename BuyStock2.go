package main

func maxProfit2(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for i := 1; i < len(prices); i++ {
		// we set buy price to the today's price if it's lower
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}
		// we get profit as if we tried to sell every day
		profit := prices[i] - buyPrice
		// if there is a profit, we add it to our max profit, and set the buyprice as we have sold the stock and now we must buy it again. We'll add to our max profit
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[i]
		}
	}
	return maximumProfit
}