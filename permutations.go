package main

import (
	"fmt"
)

type PermutationGenerator struct {
	nums   []int
	result [][]int
}

func (pg *PermutationGenerator) permuteHelper(current []int) {
	if len(current) == len(pg.nums) {
		// Make a copy of the current slice, as slices are reference types
		temp := make([]int, len(current))
		copy(temp, current)
		pg.result = append(pg.result, temp)
		return
	}

	for _, num := range pg.nums {
		if contains(current, num) {
			continue
		}
		// Choose
		newCurrent := append(current, num)
		// Explore
		pg.permuteHelper(newCurrent)
	}
}

func main() {
	scale := 6
	var nums []int

	for i := range scale {
		nums = append(nums, i+1)
	}

	pg := PermutationGenerator{nums: nums}
	pg.permuteHelper([]int{}) // pass in an empty slice to be used as the local 'current' tracker

	for _, p := range pg.result {
		fmt.Println(p)
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
