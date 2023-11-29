package main

func Trap(height []int) int {
	if height == nil {
		return 0
	}
	l,r := 0, len(height) -1
	// we need to get track of the left max and right max
	leftMax, rightMax := height[l], height[r]
	res := 0

	for l < r {
		// if leftmax is less than right max
		if leftMax < rightMax {
			l++
			leftMax = maxInt(leftMax, height[l])
			res = res + (leftMax - height[l])
		// if rightmax is less than left max
		} else {
			r--
			rightMax = maxInt(rightMax, height[l])
			res  = res + (rightMax - height[r])
		}
	}
	return res
}