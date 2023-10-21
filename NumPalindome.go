package main



func isNumPalindome(x int ) bool {
	original := x
    reversed := 0

    for x > 0 {
        digit := x % 10
        reversed = reversed*10 + digit
        x /= 10
    }

    return original == reversed
}