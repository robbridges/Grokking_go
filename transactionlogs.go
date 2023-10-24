package main


func ParseLogs(logs [][]int, threshold int ) []int {
	var result = []int{}
	userMap := make(map[int]int)

	for _, log := range logs {
		if log[0] != log[1] {
			userMap[log[0]]++
			userMap[log[1]]++
		} else {
			userMap[log[0]]++
		}
	}
	for k, v := range userMap {
		if v >= threshold {
			result = append(result, k)
		}

	}

	return result
}