package main

import (
	"fmt"
	"math"
)

const notFound = -1

// searchRange
func searchRange(nums []int, target int) []int {
	start := findStartEnd(nums, target, -1)
	end := findStartEnd(nums, target, +1)
	return []int{start, end}
}

// main
func main() {
	var nums []int = []int{5, 7, 7, 8, 8, 10}
	s := searchRange(nums, 7)
	fmt.Println(s)
	nums = []int{5, 7}
	s = searchRange(nums, 7)
	fmt.Println(s)
	nums = []int{5, 7, 7, 8, 8, 10}
	s = searchRange(nums, 8)
	fmt.Println(s)
	nums = []int{2, 2}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{2, 2, 2}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{1, 2, 2}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{1, 2, 3}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{1, 1, 2, 4}
	s = searchRange(nums, 2)
	fmt.Println(s)
	nums = []int{1, 1, 2}
	s = searchRange(nums, 1)
	fmt.Println(s)
	nums = []int{1, 1, 2, 2}
	s = searchRange(nums, 1)
	fmt.Println(s)
	nums = []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	s = searchRange(nums, 3)
	fmt.Println(s)
}

// findStartAndEnd finds the start or end index of the
// occurrences of a particular number. Takes
// an offset as an argument. Pass in -1 for beginning
// or +1 for end
func findStartEnd(nums []int, target int, offset int) int {
	var length int = len(nums)
	var overshot bool = false
	var position int
	var direction int = 1
	var velocity int = 0
	var exponent int = 0

	// Edge cases
	if length == 0 {
		return notFound
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return notFound
	}
	if offset < 0 && nums[0] == target && nums[0] == nums[length-1] {
		return 0
	}
	if offset > 0 && nums[0] == target && nums[0] == nums[length-1] {
		return length - 1
	}
	if length > 2 && nums[1] == target && nums[0] != target && nums[2] != target {
		return 1
	}

	for {
		if position > length-1 {
			if overshot {
				return notFound
			}
			overshot = true
			position = length - 1
			direction = -direction
			velocity = 1
			continue
		}
		if position < 0 {
			if overshot {
				return notFound
			}
			overshot = true
			position = 0
			direction = -direction
			velocity = 1
			continue
		}

		if nums[position] == target {
			// found
			plusCheckable := position+offset >= 0 && position+offset <= length-1
			minusCheckable := position-offset >= 0 && position-offset <= length-1
			bothCheckable := plusCheckable && minusCheckable

			// either at the beginning or end of the array
			if !bothCheckable {
				// end of the array
				if position == length-1 && offset > 0 {
					return position
				}
				// start of the array
				if position == 0 && offset < 0 {
					return position
				}
			}
			// only one element
			if bothCheckable && nums[position+offset] != target && nums[position-offset] != target {
				return position
			}
			// end/start with multiple elements
			if plusCheckable && nums[position+offset] != target {
				return position
			}
			// overshot, reset in opposite direction
			direction = -1
			position += direction
			velocity = 0
			exponent = 0
			continue
		} else if nums[position] < target {
			// accelerate
			direction = 1
			velocity += power(2, exponent)
		} else if nums[position] > target {
			// accelerate
			direction = -1
			velocity += power(2, exponent)
		}
		// update
		position += velocity * direction
	}
}

// power
func power(base, exponent int) int {
	result := int(math.Pow(float64(base), float64(exponent)))
	return result
}
