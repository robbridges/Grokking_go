package main

func UniqueOccurrences(arr []int) bool {
 numMap := make(map[int]int)
 seen := make(map[int]bool)
 for _, num := range arr {
	numMap[num]++
 }
 for _, val := range numMap {
	if seen[val] {
		return false
	}
	seen[val] = true

 }
 return true
 
}