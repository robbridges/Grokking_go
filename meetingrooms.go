package main

import "sort"

type Interval struct {
	start int
	end int
}

func canAttendMeetings(intervals [][]int) bool {
	starts := []int{}
	ends := []int{}
	for _, interval := range intervals {
    starts = append(starts, interval[0])
    ends = append(ends, interval[1])
}
sort.Ints(starts)
sort.Ints(ends)

for i := 0; i < len(intervals) - 1; i++ {
    // check to see if the next meeting starts before this meeting ends if so, return false
    if starts[i+1] < ends[i] {
        return false
    }
}

return true
}