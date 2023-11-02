package main

func RemoveDuplicates(nums []int) []int {
	numMap := make(map[int]bool)
	res := []int{}

	for _, num := range nums {
		if !numMap[num] {
			res = append(res, num)
		}
		numMap[num] = true
	}

	return res
}