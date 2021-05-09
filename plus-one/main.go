package main

func main() {
	plusOne([]int{1, 2, 3})
}

func plusOne(digits []int) []int {
	size := len(digits)
	result := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		digit := digits[i]
		if digit < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	result = append([]int{1}, digits...)
	return result
}
