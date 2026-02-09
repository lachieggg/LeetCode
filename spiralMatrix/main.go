package main

import (
	"fmt"
)

const (
	debug = true
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	printMatrix(matrix)
	print(spiralOrder(matrix)) // [1,2,3,6,9,8,7,4,5] 3×3

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	printMatrix(matrix)
	print(spiralOrder(matrix)) // [1,2,3,4,8,12,11,10,9,5,6,7] 3×4
}

func spiralOrder(matrix [][]int) []int {
	elems := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	var x, y int

	// bounds
	var yStart, xStart, yEnd, xEnd = 0, 0, m - 1, n - 1

	// velocity
	var horizontal, vertical = 1, 0

	for len(elems) < m*n {
		elems = append(elems, matrix[y][x])

		if y == yStart && x == xEnd && horizontal == 1 {
			// top right corner
			yStart += 1
			vertical = 1
			horizontal = 0
		} else if y == yEnd && x == xStart && horizontal == -1 {
			// bottom left corner
			yEnd -= 1
			vertical = -1
			horizontal = 0
		} else if x == xStart && y == yStart && vertical == -1 {
			// top left corner
			xStart += 1
			horizontal = 1
			vertical = 0
		} else if x == xEnd && y == yEnd && vertical == 1 {
			// bottom right corner
			xEnd -= 1
			horizontal = -1
			vertical = 0
		}

		moveCursor(&x, &y, horizontal, vertical)
	}

	return elems
}

func moveCursor(x *int, y *int, horizontal int, vertical int) {
	*x = *x + horizontal
	*y = *y + vertical
}

// printMatrix
// fancy print matrix function
// [Sonnet 4.5]
func printMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	n := len(matrix)

	for col := 0; col < n; col++ {
		if col == 0 {
			fmt.Print("┌")
			for j := 0; j < n; j++ {
				fmt.Print("────")
				if j < n-1 {
					fmt.Print("┬")
				}
			}
			fmt.Println("┐")
		}

		fmt.Print("│")
		for row := 0; row < n; row++ {
			// Swap row & columns for transpose
			fmt.Printf("%3d │", matrix[col][row])
		}
		fmt.Println()

		if col < n-1 {
			fmt.Print("├")
			for j := 0; j < n; j++ {
				fmt.Print("────")
				if j < n-1 {
					fmt.Print("┼")
				}
			}
			fmt.Println("┤")
		}
	}

	fmt.Print("└")
	for j := 0; j < n; j++ {
		fmt.Print("────")
		if j < n-1 {
			fmt.Print("┴")
		}
	}
	fmt.Println("┘")
}

func print(s any) {
	if debug {
		fmt.Printf("%v\n", s)
	}
}
