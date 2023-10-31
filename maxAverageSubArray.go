package main

import "fmt"

func FindMaxAverage(nums []int, k int) float64 {
	
	for i := 0; i < len(nums) - k; i++ {
		sum := nums[i] + nums[i +1] + nums[i + 2] + nums[i + 3]
		fmt.Println(sum)
	}

	return 0.0
}