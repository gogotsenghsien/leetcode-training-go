package main

import "sort"

func sortArray(nums []int) []int {
	n := 1 // choose the case you want

	switch n {
	case 1:
		// 1. built-in sort method
		return builtInSort(nums)
	case 2:
		// 2. bubble sort
		return bubbleSort(nums)
	case 4:
		// 4. merge sort
		return MergeSort(nums)
	}

	return []int{}
}

func builtInSort(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func bubbleSort(nums []int) []int {
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

func MergeSort(nums []int) []int {
	if len(nums) > 1 {
		mid := len(nums) / 2
		l := nums[:mid]
		r := nums[mid:]
		l = sortArray(l)
		r = sortArray(r)
		i, j := 0, 0
		res := make([]int, 0)

		for i < len(l) && j < len(r) {
			if l[i] < r[j] {
				res = append(res, l[i])
				i++
			} else {
				res = append(res, r[j])
				j++
			}
		}

		for i < len(l) {
			res = append(res, l[i])
			i++
		}

		for j < len(r) {
			res = append(res, r[j])
			j++
		}
		return res
	}
	return nums
}
