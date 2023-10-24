package main

import (
	"sort"
)

func ThreeSumTwo(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	// we need to go through every nums except subtract the length because we're using two pointers
	for i := 0; i < len(nums)-2; i++ {
		// if we've already gone through the loop once and the current numbers and one before it are the same , continue to the next step
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// initialize a left and right pointer
		left, right := i+1, len(nums)-1
		// while left is smaller than right
		for left < right {
			// check the value
			target := nums[i] + nums[left] + nums[right]
			// if zero append it to our array and increment the left and right 
			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1
				// if the left is less than the right and the the current iteraton and one before is the same, move the left forward to eliminate duplicates
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// if the right is the same move the right forward one
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			// if the result is too big, subtract the right
			} else if target > 0 {
				right--
			// otherwise move the left forward
			} else {
				left++
			}
		}
	}
	// return answer
	return res
}