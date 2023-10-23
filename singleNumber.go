package main

import (

)

func SingleNumber(nums []int) int {
	numsMap := make(map[int]int)
	
	for _, v := range nums {
		numsMap[v]++
	}

	for k,v := range numsMap {
		if v != 2 {
			return k
		}
	}
	return 0

	//TODO look into XOR
	// for i := 1; i < len(nums); i++ {
    //     nums[0] ^= nums[i]
    // }
    // return nums[0]
}