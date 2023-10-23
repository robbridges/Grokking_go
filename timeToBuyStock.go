package main

import("math")
func maxProfit(prices []int) int {
	// set to pointers left and right
	left, right := 0, 1
	maxP := 0.0

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			maxP = math.Max(float64(maxP), float64(profit))
		} else {
			left = right
		}
		right += 1
	}

	return int(maxP)
}