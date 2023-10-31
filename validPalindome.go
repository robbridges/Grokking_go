package main

import (
	"strings"
	"unicode"
)

func IsPaldindome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func CleanPalindome(s string) bool {
	var result strings.Builder

	for _, char := range s {
		if unicode.IsLetter(char) {
			result.WriteRune(unicode.ToLower(char))
		} else if unicode.IsNumber(char) {
			result.WriteRune(char)	
		}
	}

	cleaned := result.String()


	return IsPaldindome(cleaned)
	
}
