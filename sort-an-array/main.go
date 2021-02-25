package main

func sortArray(nums []int) []int {
	// 1. built-in sort method
	//sort.Ints(nums)
	//return nums

	// 2. bubble sort
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
