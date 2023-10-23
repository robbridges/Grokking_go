package main

func RunningSum(nums []int) [] int {
	// total := nums[0]
	// result := []int{}
	// result = append(result, nums[0])
	// for i := 1; i < len(nums) -1; i++ {
		
	// 	total = total + nums[i]
	// 	result = append(result, total)
	// 	fmt.Println(total)
	// }

	// return result

	for i:=1;i<len(nums);i++{
		nums[i]=nums[i]+nums[i-1]
		}
		return nums
}