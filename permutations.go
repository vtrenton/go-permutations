package main

import (
	"fmt"
)

type PermutationGenerator struct {
	nums    []int
	current []int
	result  [][]int
}

func (pg *PermutationGenerator) Generate() [][]int {
	pg.permuteHelper()
	return pg.result
}

func (pg *PermutationGenerator) permuteHelper() {
	if len(pg.current) == len(pg.nums) {
		// Make a copy of the current slice, as slices are reference types
		temp := make([]int, len(pg.current))
		copy(temp, pg.current)
		pg.result = append(pg.result, temp)
		return
	}

	for _, num := range pg.nums {
		if contains(pg.current, num) {
			continue
		}
		// Choose
		pg.current = append(pg.current, num)
		// Explore
		pg.permuteHelper()
		// Unchoose (backtrack)
		pg.current = pg.current[:len(pg.current)-1]
	}
}

func main() {
	scale := 6
	var nums []int
	for i := range scale {
		nums = append(nums, i+1)
	}

	pg := PermutationGenerator{nums: nums}
	permutations := pg.Generate()
	for _, p := range permutations {
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
