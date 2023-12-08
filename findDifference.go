package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := addToMap(nums1)
	map2 := addToMap(nums2)
 
	res := [][]int{}
 
	res = append(res, compareMaps(map1, map2))
	res = append(res, compareMaps(map2, map1))
 
	return res
 }
 
 func addToMap(ints []int) map[int]bool {
   numMap := make(map[int]bool)
   for _, num := range ints {
	 if numMap[num] {
	   continue
	 }
	 numMap[num] = true
   }
   return numMap
 }
 
 func compareMaps(map1, map2 map[int]bool) []int {
   diff := []int{}
   for num, _ := range map1 {
	  if !map2[num] {
		diff = append(diff, num)
	  }
	}
	return diff
 }
 