package main

import "fmt"



func main() {
	result := ParseLogs([][]int{{1,2, 200}, {1,3, 300}, {1, 1, 600}, {2,2, 300}, {2,1, 200}}, 3)
	fmt.Println(result)
}



