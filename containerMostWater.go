package main

func maxArea(height []int) int {
	//brute force

	// n := len(height)

	// var result int

	// for left := 0; left < n-1; left++ {
	// 	for right := left + 1; right < n; right++ {
	// 		currArea := min(height[left], height[right]) * (right - left) // min height * width
	// 		result = maxInt(result, currArea)
	// 	}
	// }
	// return result

	res := int32(0)
	l, r := 0, len(height) -1
	for l < r {
		// get the current area, you get the min height, as water can't over flow one side, times the distance of the two points
		currArea := int32(min(height[l], height[r])) * int32(r - l)
		// check to see if the current area is greater than our current response
		res = int32(maxInt(int(res), int(currArea)))
		// 
		if height[l] < height[r] {
			l++
		}  else  {
			r--
		} 
	}
	return int(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}