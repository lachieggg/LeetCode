package main

import (
	"fmt"
)

// generateMatrix
func generateMatrix(n int) [][]int {
	// starting pos
	var x, y = 0, 0

	// initial value
	var k = 1

	// bounds
	var yStart, xStart, yEnd, xEnd = 0, 0, n - 1, n - 1

	// velocity
	var horizontal, vertical = 1, 0

	// create matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for k <= n*n {
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

		matrix[y][x] = k
		x += horizontal
		y += vertical
		k += 1
	}
	return matrix
}

func main() {
	var matrix [][]int
	for i := range 25 {
		matrix = generateMatrix(i)
		printMatrix(matrix)
	}
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

const debug = false

func print(s string) {
	if debug {
		fmt.Printf(s)
	}
}
