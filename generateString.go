package main

import "strings"

func generateTheString(n int) string {
	if n% 2 == 0  {
        return "b" + strings.Repeat("a", n-1)
    }

    return strings.Repeat("a", n)


}