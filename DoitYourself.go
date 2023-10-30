package main

import (
	"math"
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

func ClosestThree(nums []int, target int) int {
	var closest int = math.MaxInt32
	ans := 0
	for i := 0; i < len(nums) -2; i++ {
		sum:= nums[i] + nums[i+1] + nums[i + 2]
		diff := target - sum
		if diff == 0 {
			return sum
		} 
		if diff < closest {
			closest = diff
			ans = sum
		}
	}
	return ans
}

// func GuessingGame(n int) int {
// 	start := 1
//     end := n
//     num := (start + end) /2

//     for start < end{
//         guessed := guess(num)
//         if guessed == 0{
//             return num
//         }else if guessed > 0{
//             start = num + 1
//             num = (start + end)/2
//         }else if guessed < 0{
//             end = num - 1
//             num = (start + end)/2
//         }
//     }

//     return num
// }
