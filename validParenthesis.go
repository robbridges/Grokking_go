package main

var Parentheses = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func IsValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		v, ok := Parentheses[c]; if ok {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		} else if stack[len(stack) -1] != c {
			return false
		}
		stack = stack[:len(stack) -1]
	}
	return len(stack) == 0
}