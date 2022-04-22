package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	sqindex := 0

	for _, val := range array {
		if sqindex == len(sequence) {
			break
		}
		if val == sequence[sqindex] {
			sqindex++
		}
	}

	return sqindex == len(sequence)
}
