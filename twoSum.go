package main

func TwoSum(nums []int, target int) []int {
	seenMap := make(map[int]int)

	for i := 0; i < len(nums ); i++ {
		want := target - nums[i]
		if _, ok := seenMap[want]; ok {
			return []int{seenMap[want], i}
		}
		seenMap[nums[i]] = i
	}
	return []int{0,0}
}

func TwoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums) -1

	for left < right {
		result := nums[left] + nums[right]
		if result == target {
			return []int{left + 1, right + 1}
		} else if result > target {
			right--
		} else  {
			left++
		}
	}
	return []int{-1,-1}
}