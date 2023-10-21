package main

import (
	"fmt"
	"strings"
)



func LengthOfLastWord(s string) int {
	// length := 0

    // // Skip trailing spaces
    // i := len(s) - 1
    // for i >= 0 && s[i] == ' ' {
    //     i--
    // }

    // // Count the characters in the last word
    // for i >= 0 && s[i] != ' ' {
    //     length++
    //     i--
    // }

    // return length
	trimmed := strings.TrimSpace(s)
	fmt.Println(trimmed)
	cur := strings.Split(trimmed, " ")
	return len(cur[len(cur) -1])
	
}