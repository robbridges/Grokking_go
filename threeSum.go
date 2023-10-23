package main

import "sort"

func ThreeSum(nums []int, target int) bool {
	// sort the array
	sort.Ints(nums)
	// iterate through the array
	for i := 0; i < len(nums); i++ {
		// set left and right pointers
		left := i + 1
		right := len(nums) - 1
		// while the left node is greater than the right node 
		for left < right {
			// check if the sum of the three numbers is equal to the target
			sum := nums[i] + nums[left] + nums[right]
			switch {
				// if the sum is equal to the target, return true
			case sum == target:
				return true
				// if the sum is less than the target, increment the left pointer
			case sum < target:
				left++

				
			case sum > target:
				right--
			}
		}
	}
	return false
}
