package main

func Rotate( nums[]int, k int) []int {
	nums = append(nums[k + 1:], nums[:k + 1]...)
	return nums
}