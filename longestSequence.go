package main

func LongestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		// if previous number not found
		if !numSet[num-1] {
			// current num set to current array index
			currentNum := num
			// streak reset to one
			currentStreak := 1
			// while a new number is found increment current num by 1 and streak by one
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// in current streak is longer than longest set longest to current
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
