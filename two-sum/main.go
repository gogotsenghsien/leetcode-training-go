package main

func twoSum(nums []int, target int) []int {
	records := map[int]int{}
	for idx, val := range nums {
		if v, ok := records[target-val]; ok {
			return []int{v, idx}
		}
		records[val] = idx
	}
	return []int{-1, -1}
}
