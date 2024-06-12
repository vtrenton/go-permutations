package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	permuteHelper(nums, []int{}, &result)
	return result
}

func permuteHelper(nums []int, current []int, result *[][]int) {
	if len(current) == len(nums) {
		// Make a copy of the current slice, as slices are reference types
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for _, num := range nums {
		if contains(current, num) {
			continue
		}
		// Choose
		current = append(current, num)
		// Explore
		permuteHelper(nums, current, result)
		// Unchoose (backtrack)
		current = current[:len(current)-1]
	}
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	scale := 6
	var nums []int
	// 1.22 for loop
	// note the +1 on the append
	for i := range scale {
		nums = append(nums, i+1)
	}
	permutations := permute(nums)
	for _, p := range permutations {
		fmt.Println(p)
	}
}
