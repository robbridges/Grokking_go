package main

import "fmt"



func main() {
	result := ConcurrentReadAndWrite([]int{1,2,3,4,5,6,7,8,9,10})
	fmt.Println(result)
}



