package main

func FindMiddleMaximumCapacity(capacity []int32) int32 {
	n := len(capacity)
	maxCapacity := int32(-1)
	// go through the list, we are using three pointers
	for a := 0; a < n-2; a++ {
		// if the capacity of b is divisible by the capacity of a
		for b := a+1; b < n-1; b++ {
			// if the capacity of c is divisible by the capacity of b
			if capacity[b] % capacity[a] == 0 {
				maxCapacity = max(maxCapacity, capacity[b])
				for c := b+1; c < n; c++ {
					if capacity[c] % capacity[b] == 0 {
						break
					}
				}
			}
		}
	}

	return maxCapacity
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
