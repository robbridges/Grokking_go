package main

func LongestConsecutiveLetters(s string) int {
	totalCount, count := 1, 1

	if len(s)==1 {
		return 1
	}

	for i := 0; i < len(s) -1; i++ {
		nextChar := i + 1
		if s[i] == s[nextChar] {
			count++
		} else {
			count = 1
		}

		if count > totalCount {
			totalCount = count
		}
	}
	return totalCount
}