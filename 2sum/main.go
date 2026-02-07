package main

import (
	"fmt"
	"sort"
)

var checks = 0

const (
	debugFlag = true
)

func debug(nums []int, target int) {
	if debugFlag {
		fmt.Printf("%v, %v\n", nums, target)
	}
}

func twoSumNaive(nums []int, target int) []int {
	var i, j int
	l := len(nums)

	for i = 0; i < l-1; i++ {
		for j = i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// twoSumMirror creates a mirror of the input array by subtracting
// each element from the target.
func twoSumMirror(nums []int, target int) []int {
	mirror := []int{}
	for _, num := range nums {
		mirror = append(mirror, num-target)
	}

	for index, x := range mirror {
		for j, y := range nums[1:] {
			if j+1 != index && x+y == 0 {
				return []int{index, j + 1}
			}
		}
	}

	return []int{}
}

// twoSumCut sorts the input and cuts the search space by looking at the
// smallest and largest numbers.
func twoSum(nums []int, target int) []int {
	var orig []int
	max := len(nums)
	var sorted = true
    sort.Ints(nums)

	if nums[0] >= 0 {
		for i, num := range nums {
			if num > target {
				max = i
			}
		}
		orig = clone(nums)
	} else if nums[max-1] < 0 {
		for i, num := range nums {
			if num < target {
				max = i
			}
		}
		orig = clone(nums)
	} else {
		sorted = false
	}

	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("max: %v\n", max)

	var result []int
	for i := 0; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			fmt.Printf("%d %d\n", i, j)
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target && sorted {
				// return the numbers, lookup indexes afterward
				result = []int{nums[i], nums[j]}
				break
			}
			if nums[i]+nums[j] == target && !sorted {
				// return the indexes, done
				result = []int{i, j}
				break
			}
		}
	}

	if result[0] != result[1] && !sorted {
		// indexes preserved, done
		return result
	}

	if result[0] != result[1] && sorted {
		// need to find the original indexes
		return []int{
			find(orig, result[0]),
			find(orig, result[1])}
	}

	if result[0] == result[1] && sorted {
		// same number, sorted list
		// find in reverse to prevent using the same index twice
		return []int{
			find(orig, result[0]),
			findReverse(orig, result[0]),
		}
	}

	if result[0] == result[1] && !sorted {
		return []int{result[0], result[1]}
	}

	return []int{}
}

func clone(arr []int) []int {
	x := make([]int, len(arr))
	copy(x, arr)
	return x
}

func find(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // return index
		}
	}
	return -1 // not found
}

func findReverse(arr []int, target int) int {
	i := len(arr) - 1
	for i > 0 {
		if arr[i] == target {
			return i // return index
		}
		i--
	}
	return -1 // not found
}

func main() {
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))        // [0 1]
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))             // [1 2]
	// fmt.Println(twoSum([]int{3, 3}, 6))                // [0 1]
	// fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))        // [1 2]
	// fmt.Println(twoSum([]int{-10, -1, -18, -19}, -19)) // [1 2]
	// fmt.Println(twoSum([]int{0, 3, -3, 4, -1}, -1))    // [4 0]
	fmt.Println(twoSum([]int{3, 2, 95, 4, -3}, 92)) // [2 4]
}
