package main

import "sort"

type Interval struct {
	start int
	end int
}

func canAttendMeetings(intervals []Interval) bool {
	starts := []int{}
	ends := []int{}
	for _, interval := range intervals {
		starts = append(starts, interval.start)
		ends = append(ends, interval.end)
	}
	sort.Ints(starts)
	sort.Ints(ends)

	for i:= 0; i < len(intervals) -1; i++ {
		if starts[i + 1] < ends[i] {
			return false
		}
	}

	return true
}