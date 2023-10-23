package main

func groupAnagrams(strs []string) [][]string {
	// make a map of an array key to value of strings
	strsTable := make(map[[26]int][]string)
	// go through each string
	for _, s := range strs {
		var count [26]int
		// get the asci value of the character of that string
		for _, ch := range s {
			// add the value of the ascii character to the array
			count[ch - 'a']++
		}
		//add the key as the count array, and for the result append to it this string. All strings that share the same ascii value array will be grouped together
		strsTable[count] = append(strsTable[count], s)
	}

	var result [][]string
	// return all values from the map, we are no longer interested in the ascii keys or what it was, that was just to group them 
	for _, group := range strsTable {
		// append the sub array to the giant array
		result = append(result, group)
	}
	// return it
	return result
}