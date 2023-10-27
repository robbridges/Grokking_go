package main

import (
	"strings"
)

func ReverseStrings(s string) string {
	broken :=  strings.Fields(s)
	left, right := 0, len(broken) -1
	for left < right {
		broken[left], broken[right] = broken[right], broken[left]
		left++
		right--
	}
	s = strings.Join(broken, " ")
	return s
}