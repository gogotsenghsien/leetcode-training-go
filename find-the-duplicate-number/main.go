package main

import "sort"

func findDuplicate(nums []int) int {
	sort.Ints(nums)

	for idx, num := range nums {
		if num == nums[idx+1] {
			return num
		}
	}

	return -1
}
