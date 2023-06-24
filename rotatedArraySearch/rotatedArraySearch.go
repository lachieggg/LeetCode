package main

import (
	"fmt"
	"math"
)

const notFound = -1

// main
func main() {
	nums := []int{4, 5, 6, 7, 1, 2, 3}
	s := search(nums, 1)
	fmt.Println(s)
}

// search
func search(nums []int, target int) int {
	length := len(nums)
	if length == 2 {
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
		return notFound
	}

	if nums[0] > nums[length-1] {
		// rotated
		rotationIndex := findOutOfOrder(nums)
		if rotationIndex == notFound {
			return notFound
		}
		arr := nums[0 : rotationIndex+1]
		fnd := bsearch(arr, target)
		if fnd != notFound {
			return fnd
		}
		arr = nums[rotationIndex:]
		fnd = bsearch(arr, target)
		if fnd != notFound {
			return fnd + rotationIndex
		}

		return notFound
	}

	found := bsearch(nums, target)
	return found
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

// findOutOfOrder find a pair of numbers that are out of order
// and then returns the index of the first number in the pair.
// Returns index of x, where x is left adjacent to y, and x > y
// Returns -1 if no such pair exists.
func findOutOfOrder(nums []int) int {
	var length int = len(nums)
	var overshot bool = false
	var position int
	var direction int = 1
	var velocity int = 1

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

		if nums[0] > nums[position] {
			if position-1 >= 0 && nums[position-1] > nums[position] {
				return position - 1 // found
			}
			// overshot, reset in opposite direction
			direction = -1
			velocity = 1
		} else {
			// accelerate
			velocity *= 2
		}
		// update
		position += velocity * direction
	}
}

// midpoint
func midpoint(num float64) int {
	return int(math.Floor(num / 2))
}
