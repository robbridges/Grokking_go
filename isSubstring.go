package main 

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0

	for ; j < len(t); j++ {
		if i < len(s) && t[j] == s[i] {
			i++
		}
	}

	if i == len(s) {
		return true
	}

	return false

}