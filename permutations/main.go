package main

import (
	"fmt"
)

// LeetCode
// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{{nums[0]}}
	}

	if n == 2 {
		return [][]int{
			{nums[1], nums[0]},
			{nums[0], nums[1]},
		}
	}

	var permutations [][]int
	var rows [][]int

	for i := 0; i < n; i++ {
		// n-1 permutations step
		// pick some number, remove it
		// get all permutations
		subarray := combine(
			nums[0:i],
			nums[i+1:],
		)
		// sub-problem
		permutations = permute(
			subarray,
		)

		// prepend nums[i] to all permutations
		for j := range len(permutations) {
			row := combine([]int{nums[i]}, permutations[j])
			rows = append(rows, row)
		}
	}

	return rows
}

// combine concatenates two arrays into one
// [Sonnet 4.5]
func combine(x []int, y []int) []int {
	result := make([]int, 0, len(x)+len(y))
	result = append(result, x...)
	result = append(result, y...)
	return result
}

func main() {
	// [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
	r := permute([]int{1, 2, 3})
	fmt.Println(r)

	// [[1, 0], [0, 1]]
	r = permute([]int{0, 1})
	fmt.Println(r)

	// [[1]]
	r = permute([]int{1})
	fmt.Println(r)

	// [[2 4 8 6] [2 4 6 8] [2 6 8 4] ...]
	r = permute([]int{2, 4, 6, 8})
	fmt.Println(r)

	// [[1 3 7 5] [1 3 5 7] [1 5 7 3] ...]
	r = permute([]int{1, 3, 5, 7})
	fmt.Println(r)

}

// printMatrix
// fancy print matrix function for rectangular matrices
// [Sonnet 4.5]
func printMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Top border
	fmt.Print("┌")
	for j := 0; j < cols; j++ {
		fmt.Print("────")
		if j < cols-1 {
			fmt.Print("┬")
		}
	}
	fmt.Println("┐")

	// Rows
	for row := 0; row < rows; row++ {
		fmt.Print("│")
		for col := 0; col < cols; col++ {
			fmt.Printf("%3d │", matrix[row][col])
		}
		fmt.Println()

		// Middle border (not after last row)
		if row < rows-1 {
			fmt.Print("├")
			for j := 0; j < cols; j++ {
				fmt.Print("────")
				if j < cols-1 {
					fmt.Print("┼")
				}
			}
			fmt.Println("┤")
		}
	}

	// Bottom border
	fmt.Print("└")
	for j := 0; j < cols; j++ {
		fmt.Print("────")
		if j < cols-1 {
			fmt.Print("┴")
		}
	}
	fmt.Println("┘")
}
