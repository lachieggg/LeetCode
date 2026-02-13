package main

import (
	"fmt"
	"math"
)

// RunCase
func RunCase(nums []int, target int, expected []int) {
	fmt.Println(nums)
	result := searchRange(nums, target)
	for i, v := range result {
		if expected[i] != v {
			fmt.Println("❌")
		}
	}
	fmt.Println("✅")
}

const (
	DNE = -1
)

// searchRange
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	fnd := binarysearch(nums, target)
	if fnd == DNE {
		return []int{-1, -1}
	}
	start := findStart(nums, target)
	end := findEnd(nums, target)
	return []int{start, end}
}

// findStart
func findStart(nums []int, target int) int {
	var length int = len(nums)
	var displacement, exponent int
	var velocity int = 1

	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return DNE
	}

	for {
		previousCheckable := displacement-1 >= 0
		nextCheckable := displacement < length-1

		if displacement < 0 || displacement > length-1 {
			exponent = 0
			if displacement < 0 {
				displacement = 0
				velocity = 1
			} else {
				displacement = length - 1
				velocity = -1
			}
			continue
		}

		if nextCheckable && nums[displacement] < target && nums[displacement+1] > target {
			return DNE
		}

		if nums[displacement] < target {
			velocity = 1
			exponent = 0
		} else if nums[displacement] > target {
			velocity = -1
			exponent = 0
		} else {
			// found
			if previousCheckable && nums[displacement-1] != target || !previousCheckable {
				return displacement
			}
			for displacement >= 0 && nums[displacement] == target {
				displacement -= 1
			}
			return displacement + 1
		}

		switch displacement {
		case 0:
			displacement += 1
			velocity = 1
			exponent = 0
		case length - 1:
			displacement -= 1
			velocity = -1
			exponent = 0
		default:
			displacement += int(math.Pow(float64(2), float64(exponent))) * velocity
			exponent += 1
		}
	}
}

// findStart
func findEnd(nums []int, target int) int {
	var length int = len(nums)
	var displacement, exponent int
	var velocity int = 1

	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return DNE
	}

	for {
		nextCheckable := displacement < length-1

		if displacement < 0 || displacement > length-1 {
			velocity = -velocity
			exponent = 0
			if displacement < 0 {
				displacement = 0
				velocity = 1
			} else {
				displacement = length - 1
				velocity = -1
			}
			continue
		}

		if nextCheckable && nums[displacement] < target && nums[displacement+1] > target {
			return DNE
		}

		if nums[displacement] < target {
			velocity = 1
			exponent = 0
		} else if nums[displacement] > target {
			velocity = -1
			exponent = 0
		} else {
			if nextCheckable && nums[displacement+1] != target || !nextCheckable {
				return displacement
			}
			for displacement <= length-1 && nums[displacement] == target {
				displacement += 1
			}
			return displacement - 1
		}

		switch displacement {
		case 0:
			displacement += 1
			exponent = 0
		case length - 1:
			displacement -= 1
			exponent = 0
		default:
			displacement += int(math.Pow(float64(2), float64(exponent))) * velocity
			exponent += 1
		}
	}
}

// binarysearch
func binarysearch(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return DNE
	}
	splitIndex := int(math.Floor((float64(length)) / 2))
	f := binarysearch(nums[0:splitIndex], target)
	if f != DNE {
		return f
	}
	f = binarysearch(nums[splitIndex:], target)
	if f != DNE {
		return f + splitIndex
	}

	return DNE
}

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
