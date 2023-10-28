package main

import (
	"sort"
	"strings"
)

func ReverseStrings(s string) string {
	broken :=  strings.Fields(s)
	left, right := 0, len(broken) -1
	for left < right {
		broken[left], broken[right] = broken[right], broken[left]
		left++
		right--
	}
	s = strings.Join(broken, " ")
	return s
}

func ThreeSumYourself(nums []int, target int) []int {
  sort.Ints(nums)
  ans := []int{}
  left, right := 1, len(nums) -1
  for left < right {
	total := nums[0] + nums[left] + nums[right]
	if total == target {
		ans = append(ans, nums[0], nums[left], nums[right])
		return ans
	} else if total > target {
		right--
	} else {
		left++
	}
	}
	return ans
  }
