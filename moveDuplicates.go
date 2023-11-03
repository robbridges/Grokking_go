package main

func removeDuplicates(nums []int) int {
	prev := 0
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[prev] {
			nums[l] = nums[prev]
			l++
		}
	}
	return l
   }