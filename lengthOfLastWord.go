package main

import (
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
	// trimmed := strings.TrimSpace(s)
	// fmt.Println(trimmed)
	// cur := strings.Split(trimmed, " ")
	// return len(cur[len(cur) -1])

    words := []string{}
    var builder strings.Builder
    for _, char := range s {
        if char == ' ' {
            if builder.Len() != 0 {
                words = append(words, builder.String())
                builder.Reset()
            }
        } else {
            builder.WriteRune(char)
        }
    }
    if builder.Len() != 0 {
        words = append(words, builder.String())
    }
    if builder.Len() != 0 {
        words = append(words, builder.String())
    }
	return len(words[len(words) - 1])
}