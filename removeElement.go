package main

func RemoveElement(nums []int, val int) int {
  for i := 0; i < len(nums); i++ {
	if nums[i] == val {
		// redefine num as new slice, appending all elements before i and all elements after i
	  nums = append(nums[:i], nums[i+1:]...)
	  i--
	}
  }
  return len(nums)   
}