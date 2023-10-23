package main

func MergeString(s1, s2 string) string {
	var result string
	shortest := len(s1)
	// redefine shortest if the second word is shorter
	if len(s2) < len(s1) {
		shortest = len(s2)
	}

	// use the shortest index so each string contains a value, and add this string
	for i := 0; i < shortest; i++ {
		result += string(s1[i]) + string(s2[i])
	}

	// we are adding any value after the shortest, this way the shortest is already added and anything in the longer string is added. This is Go's index before : index after trick We're telling it 
	// only start adding any character after the shortest
	result += s1[shortest:] + s2[shortest:]
	return result
}