package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RunCase
func RunCase(nums []int, target int, expected []int) {
	fmt.Println(nums)
	result := searchRange(nums, target)
	assert.Equal(&testing.T{}, result, expected)
}

const notFound = -1

// searchRange
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	fnd := bsearch(nums, target)
	if fnd == notFound {
		return []int{-1, -1}
	}
	start := findStart(nums, target)
	end := findEnd(nums, target)
	return []int{start, end}
}

// findStart
func findStart(nums []int, target int) int {
	var length int = len(nums)
	var position int
	var direction int = 1
	var exponent int = 0

	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return notFound
	}

	for {
		minusCheckable := position-1 >= 0
		plusCheckable := position+1 <= length-1
		currCheckable := position >= 0 && position <= length-1

		if !currCheckable {
			exponent = 0
			if position < 0 {
				position = 0
				direction = 1
			} else {
				position = length - 1
				direction = -1
			}
			continue
		}

		if plusCheckable && nums[position] < target && nums[position+1] > target {
			return notFound
		}

		if nums[position] == target {
			// found
			if minusCheckable && nums[position-1] != target {
				// first one
				return position
			} else if !minusCheckable {
				// start of array
				return position
			}
			for position >= 0 && nums[position] == target {
				position -= 1
			}
			return position + 1
		} else if nums[position] < target {
			direction = 1
			exponent = 0
		} else if nums[position] > target {
			direction = -1
			exponent = 0
		}

		switch position {
		case 0:
			position += 1
			direction = 1
			exponent = 0
		case length - 1:
			position -= 1
			direction = -1
			exponent = 0
		default:
			position += power(2, exponent) * direction
			exponent += 1
		}
	}
}

// findStart
func findEnd(nums []int, target int) int {
	var length int = len(nums)
	var position int
	var direction int = 1
	var exponent int = 0

	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return notFound
	}

	for {
		plusCheckable := position+1 <= length-1
		currCheckable := position >= 0 && position <= length-1

		if !currCheckable {
			direction = -direction
			exponent = 0
			if position < 0 {
				position = 0
				direction = 1
			} else {
				position = length - 1
				direction = -1
			}
			continue
		}

		if plusCheckable && nums[position] < target && nums[position+1] > target {
			return notFound
		}

		if nums[position] == target {
			// found
			if plusCheckable && nums[position+1] != target {
				// last one
				return position
			} else if !plusCheckable {
				// end of array
				return position
			}
			for position <= length-1 && nums[position] == target {
				position += 1
			}
			return position - 1
		} else if nums[position] < target {
			direction = 1
			exponent = 0
		} else if nums[position] > target {
			direction = -1
			exponent = 0
		}

		switch position {
		case 0:
			position += 1
			exponent = 0
		case length - 1:
			position -= 1
			exponent = 0
		default:
			position += power(2, exponent) * direction
			exponent += 1
		}
	}
}

// power
func power(base, exponent int) int {
	result := int(math.Pow(float64(base), float64(exponent)))
	return result
}

// bsearch
func bsearch(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return notFound
	}
	splitIndex := midpoint(float64(length))
	fnd := bsearch(nums[0:splitIndex], target)
	if fnd != notFound {
		return fnd
	}
	fnd = bsearch(nums[splitIndex:], target)
	if fnd != notFound {
		return fnd + splitIndex
	}

	return notFound
}

// midpoint
func midpoint(num float64) int {
	return int(math.Floor(num / 2))
}

// main
func main() {
	RunCase([]int{}, 2, []int{-1, -1})
	RunCase([]int{2, 2}, 2, []int{0, 1})
	RunCase([]int{5, 7}, 7, []int{1, 1})
	RunCase([]int{2, 2}, 5, []int{-1, -1})
	RunCase([]int{2, 2, 2}, 2, []int{0, 2})
	RunCase([]int{1, 2, 2}, 2, []int{1, 2})
	RunCase([]int{1, 1, 2}, 1, []int{0, 1})
	RunCase([]int{1, 2, 3}, 2, []int{1, 1})
	RunCase([]int{1, 1, 2, 4}, 2, []int{2, 2})
	RunCase([]int{1, 1, 2, 2}, 1, []int{0, 1})
	RunCase([]int{5, 7, 7, 8, 8, 10}, 7, []int{1, 2})
	RunCase([]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4})
	RunCase([]int{0, 0, 1, 2, 3, 3, 4}, 2, []int{3, 3})
	RunCase([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3, []int{2, 5})
	RunCase([]int{0, 0, 1, 1, 1, 2, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 10}, 4, []int{7, 8})
}
