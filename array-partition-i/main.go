package main

import (
	"math"
	"sort"
)

func arrayPairSum(nums []int) int {
	result := 0
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		sum := math.Min(float64(nums[i]), float64(nums[i+1]))
		result += int(sum)
	}

	return result
}
