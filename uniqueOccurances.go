package main

func UniqueOccurrences(arr []int) bool {
 seen := make(map[int]bool) 

 for _, v := range arr {
	if seen[v] {
		return false
	}

	seen[v] = true
 }
 
 
 return true
 
}