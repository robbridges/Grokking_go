package main

import (
	"fmt"
	"strconv"
)

func Reverse(x int) int {
    // Convert the integer to a string
    s := strconv.Itoa(x)

    // Initialize an empty string for the reversed result
    res := ""

    // Check if x is negative
    if x < 0 {
        // If x is negative, add the minus sign to the result
        res += "-"
        // Make x positive for the reversal
        s = s[1:]
    }

    // Iterate through the characters of the string s in reverse order
    for i := len(s) - 1; i >= 0; i-- {
        res += string(s[i])
    }

    // Convert the reversed string back to an integer
    reversedInt, err := strconv.Atoi(res)
    if err != nil {
        fmt.Println(err)
    }

    // Check for integer overflow
    if reversedInt < -2147483648 || reversedInt > 2147483647 {
        return 0
    }

    return reversedInt
}