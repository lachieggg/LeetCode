package main

import (
	"fmt"
)

func main() {
	var w, x, y, z int
	w = majorityElement([]int{6, 7, 7})
	x = majorityElement([]int{6, 6, 7, 7, 7})
	y = majorityElement([]int{6, 6, 7, 7, 7, 7})
	z = majorityElement([]int{6, 6, 6, 7, 7, 7, 7})
	fmt.Println(w, x, y, z)
}

func majorityElementNaive(nums []int) int {
	var l int
	m := make(map[int]int, l)

	for _, n := range nums {
		// calculate the length in the loop
		// rather than using len(nums)
		l += 1
		m[n] += 1
	}

	for _, n := range nums {
		if m[n] > l/2 {
			return n
		}
	}

	return -1
}

// majorityElement optimised
// O(n) time, O(1) space complexity
// Boyer-Moore Majority Vote Algorithm
func majorityElement(nums []int) int {
	var majority, candidate int

	for _, n := range nums {
		if majority == 0 {
			candidate = n
		}

		if candidate == n {
			majority++
		} else {
			majority--
		}
	}

	return candidate
}
