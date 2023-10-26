package main

import "fmt"

func FindDifference(nums1 []int, nums2 []int) [][]int {
    ans := [][]int{{}, {}}
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, num := range nums1 {
		if map1[num] {
			continue
		}
		map1[num] = true
	}
	for _, num := range nums2 {
		if map2[num] {
			continue
		}
		map2[num] = true
	}
	for k := range map1 {
		if !map2[k] {
				// Append to the existing subarray
				ans[0] = append(ans[0], k)
			
		}
	}
	for k := range map2 {
		if !map1[k] {
			
				ans[1] = append(ans[1], k)
			}
		
	}
	fmt.Println(map1, map2)
	return ans

}