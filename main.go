package main

import "fmt"



func main() {
	result := canAttendMeetings([]Interval{{start: 1, end: 3}, {start: 5, end:10}, {start:15, end: 20}})
	fmt.Println(result)
}



