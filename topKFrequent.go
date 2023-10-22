package main


func TopFrequent(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}

	// we need 2 maps, one for our freq, one for our count and values. and a slice for the answers we'll return
	m:= make(map[int]int)
	freq := make(map[int][]int)
	res := []int{}

	//populate our map
	for _, v := range nums {
		m[v]++
	}

	// we need to go through our map and add any numbers that appear within that frequency
	for key, v := range m {
		freq[v] = append(freq[v], key)
	}

	// start at the endo f the map and keep going until our response is hte length of k, moving towards the start
	for i := len(nums); len(res) != k; i-- {
		// check the top most element to see if any elements exist
		if _, exists := freq[i]; exists {
			// append it to our response, any elements that are within that frequency should be returned
			res = append(res, freq[i]...)
		}
	}
	
	

	return res
}

