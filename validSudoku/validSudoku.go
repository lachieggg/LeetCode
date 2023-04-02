package main

import "fmt"

const boardSize = 9
const squareSize = 3
const empty = 46

// Coords
type Coords struct {
	X int
	Y int
}

// containsDuplicate determines whether a set of
// numbers contains a duplicate
func containsDuplicate(arr []byte) bool {
	for first, _ := range arr {
		for second := first + 1; second < boardSize; second++ {
			if arr[first] == empty {
				continue
			}
			if arr[first] == arr[second] {
				fmt.Printf("Duplicate, %d\n", arr[first])
				return true
			}
		}
	}
	return false
}

// getColumn
func getColumn(number int, board [][]byte) []byte {
	if number > boardSize-1 {
		return nil
	}
	var column []byte
	for i := 0; i < boardSize; i++ {
		column = append(column, board[i][number])
	}
	return column
}

// getRow
func getRow(number int, board [][]byte) []byte {
	if number > boardSize-1 {
		return nil
	}
	return board[number]
}

// getSquare
func getSquare(coords Coords, board [][]byte) []byte {
	var sq []byte
	for x := coords.X * squareSize; x < squareSize*(coords.X+1); x++ {
		for y := coords.Y * squareSize; y < squareSize*(coords.Y+1); y++ {
			sq = append(sq, board[y][x])
		}
	}
	return sq
}

// isValidSudoku
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < boardSize; i++ {
		if containsDuplicate(getRow(i, board)) {
			fmt.Printf("Bad row %d\n", i)
			return false
		}
		if containsDuplicate(getColumn(i, board)) {
			fmt.Printf("Bad column %d\n", i)
			return false
		}
	}

	for x := 0; x < squareSize; x++ {
		for y := 0; y < squareSize; y++ {
			coords := Coords{X: x, Y: y}
			if containsDuplicate(getSquare(coords, board)) {
				fmt.Printf("Bad square x=%d y=%d\n", x, y)
				return false
			}
		}
	}
	return true
}

// byteArrayToStringArray
func byteArrayToStringArray(byteArr []byte) []string {
	strArr := make([]string, len(byteArr))
	for i, b := range byteArr {
		strArr[i] = string(b)
	}
	return strArr
}

// byteArray2DToString2DArray
func byteArray2DToString2DArray(byteArr [][]byte) [][]string {
	strArr := make([][]string, len(byteArr))
	for i, row := range byteArr {
		strArr[i] = byteArrayToStringArray(row)
	}
	return strArr
}

// stringArrayToByteArray
func stringArrayToByteArray(strArr []string) []byte {
	byteArr := make([]byte, len(strArr))
	for i, str := range strArr {
		byteArr[i] = str[0]
	}
	return byteArr
}

// string2DArrayToByteArray2D
func string2DArrayToByteArray2D(strArr [][]string) [][]byte {
	byteArr := make([][]byte, len(strArr))
	for i, row := range strArr {
		byteArr[i] = stringArrayToByteArray(row)
	}
	return byteArr
}

func main() {
	var board = [][]byte{
		{53, 51, 46, 46, 55, 46, 46, 46, 46},
		{54, 46, 46, 49, 57, 53, 46, 46, 46},
		{46, 57, 56, 46, 46, 46, 46, 54, 46},
		{56, 46, 46, 46, 54, 46, 46, 46, 51},
		{52, 46, 46, 56, 46, 51, 46, 46, 49},
		{55, 46, 46, 46, 50, 46, 46, 46, 54},
		{46, 54, 46, 46, 46, 46, 50, 56, 46},
		{46, 46, 46, 52, 49, 57, 46, 46, 53},
		{46, 46, 46, 46, 56, 46, 46, 55, 57},
	}

	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		{56, 51, 46, 46, 55, 46, 46, 46, 46},
		{54, 46, 46, 49, 57, 53, 46, 46, 46},
		{46, 57, 56, 46, 46, 46, 46, 54, 46},
		{56, 46, 46, 46, 54, 46, 46, 46, 51},
		{52, 46, 46, 56, 46, 51, 46, 46, 49},
		{55, 46, 46, 46, 50, 46, 46, 46, 54},
		{46, 54, 46, 46, 46, 46, 50, 56, 46},
		{46, 46, 46, 52, 49, 57, 46, 46, 53},
		{46, 46, 46, 46, 56, 46, 46, 55, 57},
	}

	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		{46, 50, 49, 46, 46, 46, 46, 46, 46},
		{46, 46, 46, 46, 54, 46, 46, 46, 46},
		{46, 46, 46, 46, 46, 46, 55, 46, 46},
		{46, 46, 46, 46, 53, 46, 46, 46, 46},
		{46, 46, 53, 46, 46, 46, 46, 46, 46},
		{46, 46, 46, 46, 46, 46, 51, 46, 46},
		{46, 46, 46, 46, 46, 46, 46, 46, 46},
		{51, 46, 46, 46, 56, 46, 49, 46, 46},
		{46, 46, 46, 46, 46, 46, 46, 46, 56},
	}

	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		{46, 46, 46, 46, 46, 46, 53, 46, 46},
		{46, 46, 46, 46, 46, 46, 46, 46, 46},
		{46, 46, 46, 46, 46, 46, 46, 46, 46},
		{57, 51, 46, 46, 50, 46, 52, 46, 46},
		{46, 46, 55, 46, 46, 46, 51, 46, 46},
		{46, 46, 46, 46, 46, 46, 46, 46, 46},
		{46, 46, 46, 51, 52, 46, 46, 46, 46},
		{46, 46, 46, 46, 46, 51, 46, 46, 46},
		{46, 46, 46, 46, 46, 53, 50, 46, 46},
	}

	fmt.Println(isValidSudoku(board))

	fmt.Println("Done")
}
