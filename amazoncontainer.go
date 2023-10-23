package main

func countStarsBetweenPipes(s string, indices [][]int) []int {
	var counts []int
	for _, pair := range indices {
		start, end := pair[0], pair[1]
		count := 0
		for i := start; i < end; i++ {
			if s[i] == '|' {
				for j := i+1; j < end; j++ {
					if s[j] == '*' {
						count++
					} else if s[j] == '|' {
						break
					}
					i = j
				}
			}
		}
		counts = append(counts, count)
	}
	return counts
}