package main

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}