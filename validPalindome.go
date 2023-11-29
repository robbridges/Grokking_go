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
// most efficent to handle a single word that is not a paldindome, assume no edge cases of non letter character.
func isPalindome(s string) bool {
	
    
	left, right := 0, len(s) - 1
	for left < right {
		leftRune := rune(s[left])
        rightRune := rune(s[right])
		if unicode.IsUpper(leftRune) {
            leftRune = unicode.ToLower(leftRune)
        }
        if unicode.IsUpper(rightRune) {
            rightRune = unicode.ToLower(rightRune)
        }
		if leftRune != rightRune {
            return false
        }
		left++
		right--
	}
	return true
}


