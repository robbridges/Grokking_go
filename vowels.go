package main

import "strings"

func ReverseVowels(s string) string {
    // To reverse (manupulate a string in any way), convert it to runes as strings are immutable in golang
    runes := []rune(s)
    n := len(runes)
    i := 0
    j := n-1 

    for i < j {
        if isVowel(runes[i]) && isVowel(runes[j]) {
            runes[i], runes[j] = runes[j], runes[i]
            i++
            j--
        }else{ 
            if !isVowel(runes[i]) {
                i++
            }
            if !isVowel(runes[j]) {
                j--
            }
        }   

    }

    return string(runes)
}

func isVowel(r rune) bool {
    vowels := "aeiouAEIOU"
    return strings.ContainsRune(vowels, r)
}

