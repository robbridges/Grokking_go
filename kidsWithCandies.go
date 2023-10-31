package main

func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, num := range candies {
		if num > max {
			max = num
		}
	}

	result := make([]bool, len(candies))
	for i, candy := range candies {
		if candy + extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}