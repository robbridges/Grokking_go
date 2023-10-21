package main



func MinimumDifference(nums []int, k int) int  {
     largest := -1
	 secondLargest := -1

	if (len(nums) == 1) {
		return 0
	}
	for _, num := range nums {
		currentLargest := largest
		if (largest != -1 && num > largest) {
			secondLargest = currentLargest
		}
		if num > largest {
			largest = num
		}
	}
	
	return largest - secondLargest
}