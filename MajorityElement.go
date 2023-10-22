package main




func MajorityElement(nums []int) int {
    // Create a map to store the count of each element
    numMap := make(map[int]int)

    // Iterate through the array
    for _, num := range nums {
        // Check if the element exists in the map
        if _, exists := numMap[num]; exists {
            // If it exists, increment the count
            numMap[num]++
        } else {
            // If it doesn't exist, add it to the map with a count of 1
            numMap[num] = 1
        }
    }

    for k, v := range numMap {
		if v >= len(nums) /2 {
		return k
		} 
	}
	return 0
}