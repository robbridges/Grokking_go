package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	ans := math.MinInt
	runningSum := 0
	for idx, val := range nums {
		runningSum += val
		if idx >= k-1 {
			if idx != k-1 {
				runningSum = runningSum - nums[idx-k]
			}
			if ans < runningSum {
				ans = runningSum
			}
		}
	}
	return float64(ans) / float64(k)
}