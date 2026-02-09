package main

import (
	"fmt"
	"math"
)

const (
	IGNORE    = -math.MaxInt16
	NOT_FOUND = -math.MaxInt8
)

func main() {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{4, 3, 2, 1, 0, 4, 0, 0}, false},
		{[]int{4, 3, 3, 3, 1, 1, 0, 4}, false},
		{[]int{4, 3, 3, 1, 0, 4}, true},
		{[]int{3, 3, 1, 0, 4}, true},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{0, 3, 1, 1, 4}, false},
		{[]int{1}, true},
		{[]int{0}, true},
		{[]int{2, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, true},
		{[]int{5, 9, 3, 8, 1, 0, 2, 3, 3, 1, 0, 0, 0}, false},
		{[]int{1, 1, 1, 1, 1}, true},
		{[]int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, true},
	}
	runTests(tests)
}

func canJump(nums []int) bool {
	fmt.Printf("%v\n", nums)
	length := len(nums)
	if length <= 1 || nums[0] >= length {
		return true
	}
	if nums[0] <= 0 {
		return false
	}

	// scroll through any positive numbers from the end
	i := length - 1
	if nums[i-1] > 0 {
		return canJump(nums[0:i])
	}

	var need int = 1
	i = length - 1
	for i > 0 {
		i--
		if nums[i] >= need && need != 0 {
			return canJump(nums[0 : i+1])
		}
		need += 1
	}

	return false
}

// runTests
// [Sonnet 4.5]
func runTests(tests []struct {
	input    []int
	expected bool
}) {
	passed := 0
	failed := 0

	for _, test := range tests {
		fmt.Printf("\n---------------------------\n")
		result := canJump(test.input)
		if result == test.expected {
			passed++
		} else {
			failed++
			fmt.Printf("✗ canJump(%v) = %v (expected %v)\n", test.input, result, test.expected)
		}
	}

	fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	if failed == 0 {
		fmt.Println("✓ All tests passed!")
	}
}
