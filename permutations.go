package main

import (
	"fmt"
	"slices"
)

type PermutationGenerator struct {
	nums   []int
	result [][]int
}

func (pg *PermutationGenerator) populateNums(scale int) {
	for i := range scale {
		pg.nums = append(pg.nums, i+1)
	}
}

func (pg *PermutationGenerator) generator(current []int) {
	if len(current) == len(pg.nums) {
		// Make a copy of the current slice, as slices are reference types
		temp := make([]int, len(current))
		copy(temp, current) // I hate this pattern but it's needed
		pg.result = append(pg.result, temp)
		return
	}

	for _, num := range pg.nums {
		if slices.Contains(current, num) {
			continue
		}
		// Choose
		newCurrent := append(current, num)
		// Explore
		pg.generator(newCurrent)
	}
}

func main() {
	scale := 6
	// create object to track global state
	var pg PermutationGenerator

	// give a list of input numbers
	pg.populateNums(scale)

	// generate all permutations in a list of lists
	pg.generator([]int{}) // pass in an empty slice to be used as the local 'current' tracker

	// print lists within the list
	for _, p := range pg.result {
		fmt.Println(p)
	}
}
