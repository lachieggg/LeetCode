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

// twoSumNaive checks every pair of numbers to see if they sum to the target.
// O(n^2)
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
// O(n^2)
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
// O(n^2)
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

// findClosest finds the key in the MultiMap that is closest to the target value
// Returns: the closest key, its associated indices slice, and whether a match was found
func findClosest(target int, m MultiMap) (int, []int, bool) {
	// Return early if map is empty
	if len(m) == 0 {
		return 0, nil, false
	}

	// Extract all keys from the map into a slice
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// Sort keys in ascending order for binary-like search
	sort.Ints(keys)

	// Iterate through sorted keys to find closest match
	for i, n := range keys {
		// Found exact match - return immediately
		if n == target {
			return n, m[n], true
		}

		// Found first key greater than target
		if n > target {
			// If it's the first key, return it (no smaller option)
			if i == 0 {
				return n, m[n], true
			}
			// Compare distances: previous key vs current key
			// Return whichever is closer to the target
			prev := keys[i-1]
			if abs(target-prev) <= abs(n-target) {
				return prev, m[prev], true
			}
			return n, m[n], true
		}
	}

	// All keys are less than target, return the largest key
	lastKey := keys[len(keys)-1]
	return lastKey, m[lastKey], true
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// twoSum signature
func twoSum(arr []int, target int) []int {
	return arr
}

func main() {
	if !slicesEqual(twoSum([]int{}, 0), []int{}) {
		fmt.Println("[], 0 ==",
			twoSum([]int{}, 0),
			"expected []") // [0 0]
	}
	if !slicesEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		fmt.Println("[2, 7, 11, 15], 9 ==",
			twoSum([]int{2, 7, 11, 15}, 9),
			"expected [0, 1]") // [0 1]
	}
	if !slicesEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}) {
		fmt.Println("[3, 2, 4], 6 ==",
			twoSum([]int{3, 2, 4}, 6),
			"expected [1, 2]") // [1 2]
	}
	if !slicesEqual(twoSum([]int{3, 3}, 6), []int{0, 1}) {
		fmt.Println("[3, 3], 6 ==",
			twoSum([]int{3, 3}, 6),
			"expected [0, 1]") // [0 1]
	}
	if !slicesEqual(twoSum([]int{2, 5, 5, 11}, 10), []int{1, 2}) {
		fmt.Println("[2, 5, 5, 11], 10 ==",
			twoSum([]int{2, 5, 5, 11}, 10),
			"expected [1, 2]") // [1 2]
	}
	if !slicesEqual(twoSum([]int{-10, -1, -18, -19}, -19), []int{1, 2}) {
		fmt.Println("[-10, -1, -18, -19], -19 ==",
			twoSum([]int{-10, -1, -18, -19}, -19),
			"expected [1, 2]") // [1 2]
	}
	if !slicesEqual(twoSum([]int{0, 3, -3, 4, -1}, -1), []int{4, 0}) {
		fmt.Println("[0, 3, -3, 4, -1], -1 ==",
			twoSum([]int{0, 3, -3, 4, -1}, -1),
			"expected [4, 0]") // [4 0]
	}
	if !slicesEqual(twoSum([]int{3, 2, 95, 4, -3}, 92), []int{2, 4}) {
		fmt.Println("[3, 2, 95, 4, -3], 92 ==",
			twoSum([]int{3, 2, 95, 4, -3}, 92),
			"expected [2, 4]") // [2 4]
	}
	if !slicesEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}) {
		fmt.Println("[3, 2, 4], 6 ==",
			twoSum([]int{3, 2, 4}, 6),
			"expected [1, 2]") // [1 2]
	}
	if !slicesEqual(twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3}) {
		fmt.Println("[0, 4, 3, 0], 0 ==",
			twoSum([]int{0, 4, 3, 0}, 0),
			"expected [0, 3]") // [0 3]
	}
	if !slicesEqual(twoSum([]int{-11, 7, 3, 2, 1, 7, -10, 11, 21, 3}, 11), []int{6, 8}) {
		fmt.Println("[-11,7,3,2,1,7,-10,11,21,3], 11 ==",
			twoSum([]int{-11, 7, 3, 2, 1, 7, -10, 11, 21, 3}, 11),
			"expected [6, 8]") // [6 8]
	}

}

// -------------------
// Helper functions
// -------------------

func slicesEqual(x, y []int) bool {
	sort.Ints(x)
	sort.Ints(y)
	if len(x) != len(y) {
		fmt.Println("❌ fail, wrong length")
		return false
	}

	for i, v := range x {
		if v != y[i] {
			fmt.Println("❌ fail, not equal")
			return false
		}
	}
	fmt.Println("✅ pass")
	return true
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

type MultiMap map[int][]int

func (m MultiMap) FromArray(arr []int) int {
	var length = 0
	for i, v := range arr {
		m[v] = append(m[v], i)
		length += 1
	}
	return length
}

func (m MultiMap) Print() {
	fmt.Printf("MultiMap contents:\n")
	for key, indices := range m {
		fmt.Printf("  %d: %v\n", key, indices)
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

func (m MultiMap) Remove(key int) {
	m[key] = []int{}
}

// findIndexes returns the indexes of numbers which add to the target
// returns the numbers that sum to the target, and the indexes.
// O(n^2)
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
// O(n)
func find(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // return index
		}
	}
	return -1 // not found
}

// findReverse returns the index of the target in arr, if it exists, starting from the end
// O(n)
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
