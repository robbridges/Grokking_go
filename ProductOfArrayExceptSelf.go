package main

func ProductExceptSelf(nums []int) []int {
	// create our array
	out := make([]int, len(nums))
	prefix := 1
	// get all prefixes and add them to the response we're preparing to send
	for i := range nums {
		out[i] = prefix
		prefix *= nums[i]
	}
	
	postfix := 1
	// get the postfixes and multiply them by the values already in the array and increment the postfix
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= postfix
		postfix *= nums[i]
	}
	return out
}