package main

import "sort"

// O(n log(n)) time and O(1) space

func TwoNumberSumArray_sol2(arr []int, target int) (result []int) {
	sort.Ints(arr)

	left, right := 0, len(arr)-1

	for right > left {
		current_sum := arr[left] + arr[right]
		if current_sum == target {
			result = []int{arr[left], arr[right]}
		}
		if current_sum > target {
			right--
		} else {
			left++
		}
	}
}
