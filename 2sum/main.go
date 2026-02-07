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
func twoSumCut(nums []int, target int) []int {
	// max index
	max := len(nums)
	ln := max
	// sorted array
	sorted := make([]int, max)
	copy(sorted, nums)
	sort.Ints(sorted)

	if sorted[0] >= 0 {
		// all values are positive
		for i, num := range sorted {
			if num > target {
				max = i
			}
		}
	} else if sorted[max-1] < 0 {
		// all values are negative
		for i, num := range sorted {
			if num < target {
				max = i
			}
		}
	}

	// whether we trimmed the array (or not)
	// the same as whether max has been altered
	var trimmed = (max != ln)

	var result []int
	var i, j int
	if trimmed {
		i, j, result = findIndexes(max, sorted, target)
	} else {
		i, j, _ = findIndexes(max, nums, target)
	}

	if trimmed && result[0] != result[1] {
		// trim occurred, so need to find the original indexes
		return []int{
			find(nums, result[0]),
			find(nums, result[1])}
	}

	if trimmed && result[0] == result[1] {
		// cut occurred, duplicate number
		// find in reverse to prevent using the same index twice
		return []int{
			find(nums, result[0]),
			findReverse(nums, result[1]),
		}
	}

	return []int{i, j}
}

// findIndexes
func findIndexes(max int, arr []int, target int) (int, int, []int) {
	for i := 0; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			if i == j {
				continue
			}
			if arr[i]+arr[j] == target {
				// return the numbers, lookup indexes afterward
				// if needed
				return i, j, []int{arr[i], arr[j]}
			}
		}
	}
	return -1, -1, []int{}
}

// find returns the index of the target in arr, if it exists
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
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))        // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))             // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))                // [0 1]
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))        // [1 2]
	fmt.Println(twoSum([]int{-10, -1, -18, -19}, -19)) // [1 2]
	fmt.Println(twoSum([]int{0, 3, -3, 4, -1}, -1))    // [4 0]
	fmt.Println(twoSum([]int{3, 2, 95, 4, -3}, 92))    // [2 4]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))             // [1 2]
}

func twoSum(nums []int, target int) []int {
	m := MultiMap{}
	m.FromArray(nums)

	for k, indexes := range m {
		if len(indexes) >= 2 && target == 2*k {
			return indexes
		}
		for j, secondary := range m {
			if j == k {
				continue
			}

			if j+k == target {
				return []int{indexes[0], secondary[0]}
			}
		}
	}

	return []int{}
}

type MultiMap map[int][]int

func (m MultiMap) FromArray(arr []int) {
	for i, v := range arr {
		m[v] = append(m[v], i)
	}
}


func (m MultiMap) Add(key, value int) {
	m[key] = append(m[key], value)
}

func (m MultiMap) Pop(key int) (int, bool) {
	values := m[key]
	if len(values) == 0 {
		return 0, false
	}

	last := values[len(values)-1]
	m[key] = values[:len(values)-1]
	return last, true
}

func (m MultiMap) Remove(key, value int) {
	values := m[key]
	for i, v := range values {
		if v == value {
			m[key] = append(values[:i], values[i+1:]...)
			return
		}
	}
}