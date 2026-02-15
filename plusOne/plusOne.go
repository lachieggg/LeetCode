package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	l := len(digits)

	if l == 0 {
		return []int{}
	}

	if digits[l-1] != 9 {
		digits[l-1] += 1
		return digits
	}

	for i := l - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] != 10 {
			break
		}
	}

	all := true
	for _, d := range digits {
		if d != 10 {
			all = false
		}
	}

	zero(digits)

	if all {
		digits[0] = 1
		digits = append(digits, 0)
	}

	return digits
}

// zero turns all 10s to 0s
func zero(digits []int) {
	for i := 0; i < len(digits); i++ {
		if digits[i] == 10 {
			digits[i] = 0
		}
	}
}

func main() {
	var x []int
	x = plusOne([]int{9}) // 10
	fmt.Println(x)

	x = plusOne([]int{1, 9, 9, 9}) // 2000
	fmt.Println(x)

	x = plusOne([]int{1, 2, 3, 4}) // 1235
	fmt.Println(x)

	x = plusOne([]int{9, 9, 9, 9}) // 10000
	fmt.Println(x)

	x = plusOne([]int{9, 8, 9}) // 990
	fmt.Println(x)

	x = plusOne([]int{2, 2, 2, 9}) // 2230
	fmt.Println(x)

	x = plusOne([]int{3, 3, 9, 9}) // 3400
	fmt.Println(x)

	x = plusOne([]int{9, 9, 9, 1}) // 9992
	fmt.Println(x)

	x = plusOne([]int{1, 9, 1, 9, 1, 9}) // 191920
	fmt.Println(x)
}
