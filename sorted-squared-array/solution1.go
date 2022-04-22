package main

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	new_array := []int{}
	for _, val := range array {
		new_array = append(new_array, val*val)
	}
	return new_array
}
