package main

var Parentheses = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func isValid(s string) bool {
	stack := []rune{}
    for _, char := range s {
        v, ok := Parentheses[char]; if ok {
            stack = append(stack, v)
            continue
        }
        if len(stack) == 0 {
            return false
        } else if stack[len(stack) - 1] != char {
            return false
        }
        stack = stack[:len(stack) -1]
    }
    return len(stack) == 0
}