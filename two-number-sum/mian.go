package main

import "fmt"

func main() {
	arr := []int{3, 4, -1, 11}
	target := 10
	result := TwoNumberSumArray_sol1(arr, target)
	fmt.Println(result)
}
