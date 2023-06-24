package main

import (
	"fmt"
	"math"
)

const debug = false
const notFound = -1

// main
func main() {
	nums := []int{3, 4, 5, 6, 1, 2}
	s := findOutOfOrder(nums)
	fmt.Println(s)
	// s := search(nums, 1)
	// fmt.Printf("result %d", s)
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
		fmt.Printf("rot at %d\n", rotationIndex)
		if rotationIndex == -1 {
			fmt.Println("err...")
			return -1
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
	var index int
	var direction int = 1
	var exponent int = 1
	var length int = len(nums)
	var overshot bool = false
	var velocity int = 1

	if nums[0] > nums[1] {
		return 0
	}

	for {
		fmt.Printf("index = %d\n", index)
		fmt.Printf("velocity = %d\n", velocity)
		velocity = direction * powInt(2, exponent)
		index = index + velocity
		if index > length-1 {
			if overshot {
				return -1
			}
			overshot = true
			fmt.Println("overshot")
			index = length - 1
			direction = -direction
			exponent = 0
		}
		if index < 0 {
			if overshot {
				return -1
			}
			overshot = true
			fmt.Println("undershot")
			index = 0
			direction = -direction
			exponent = 0
		}
		if nums[0] > nums[index] {
			if index-1 >= 0 && nums[index-1] > nums[index] {
				return index - 1 // found
			}
			if index+1 <= length-1 && nums[index] > nums[index+1] {
				return index // found
			}
			// overshot, reset in opposite direction
			direction = -direction
			exponent = 1
		} else {
			// accelerate
			exponent += 1
		}
	}
}

// powInt
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// midpoint
func midpoint(num float64) int {
	return int(math.Floor(num / 2))
}
