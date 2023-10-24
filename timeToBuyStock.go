package main

import("math")

func MaxProfit(prices []int) int {
	// set to pointers left and right
	left, right := 0, 1
	maxP := 0.0

	for right < len(prices) {
		// if there is a profit 
		if prices[left] < prices[right] {
			// check what the profit is, and reassign maxProfit if it's more than the current max profit
			profit := prices[right] - prices[left]
			maxP = math.Max(float64(maxP), float64(profit))
		} else {
			// if no profit move the left to the right
			left = right
		}
		// increment right
		right += 1
	}

	return int(maxP)
}