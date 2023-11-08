package main

func containsNearbyDuplicate(nums []int, k int) bool {
    arrayMap := make(map[int]int)

    for i, v := range nums {
        if prevIndex, ok := arrayMap[v]; ok {
            if i - prevIndex <= k {
                return true
            }
        }
        arrayMap[v] = i
    }
    return false
}