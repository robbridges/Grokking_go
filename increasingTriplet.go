package main

import "math"

func increasingTriplet(nums []int) bool {
    first, second := math.MaxInt32, math.MaxInt32
    for _, num := range nums {
        if num <= first {
            first = num
        }
        if num > first && num <= second {
            second = num
        }
        if num > second {
            return true
        }
    }
    return false
}