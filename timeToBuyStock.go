package main



func MaxProfit(prices []int) int {
	// set to pointers left and right
	left, right := 0, 1
	maxProfit := 0 
	for right < len(prices) {
	 if prices[left] < prices[right] {
		tempProfit := prices[right] - prices[left]
		if tempProfit > maxProfit {
			maxProfit = tempProfit
		}
	} else {
	   left = right
	}
	 right++
	}
	return maxProfit
}


