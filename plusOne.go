package main

func PlusOne(digits []int) []int {
	// we need to check the last digit of the array
    for i := len(digits) -1; i >= 0; i-- {
		// if it's less than 9 simply add one and return it
		if digits[i] < 9 {
			digits[i]++
			return digits
		} 
		// if we reach this part than each digit was 9. We need to covert them all to 0 and append a 1
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}