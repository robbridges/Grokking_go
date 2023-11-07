package main

func maxProfit2(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for i := 1; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}
		profit := prices[i] - buyPrice
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[i]
		}
	}
	return maximumProfit
}