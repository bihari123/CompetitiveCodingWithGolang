package main

// O(n) time and O(n) space
func TwoNumberSumArray_sol1(arr []int, target int) (result []int) {
	numbers := map[int]bool{}
	for _, val := range arr {
		potential := target - val
		if _, found := numbers[potential]; found {
			result = []int{val, potential}
		}
		numbers[val] = true

	}

	return
}
