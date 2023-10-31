package main

func MoveZeroes(nums []int) {
	// zeroCounts := 0 
	// ans := []int{}
	// for _, num := range nums {
	// 	if num != 0 {
	// 		ans = append(ans, num)
	// 	} else {
	// 		zeroCounts++
	// 	}
	// }

	// for zeroCounts > 0 {
	// 	ans = append(ans, 0)
	// 	zeroCounts--
	// }
	// return ans

		// this one is kind of tricky, sorting an array in place, we set two pointers, l and r to 0
		// increment r, it is our fast pointer
		// move r by one until it counters a non zero value, since l only moves once we found one r will be pushed to the front of the array
		// this will continue until r reaches the end, and l moves up every time we find a non zero value, we basically keep a pointer at the front of our array and iterate
		// with r
		for l, r := 0, 0; r < len(nums); r++ {
			if nums[r] != 0 {
				nums[l], nums[r] = nums[r], nums[l]
				l++
			}
		}
		
	
}