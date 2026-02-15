package main

import (
	"fmt"
	"sort"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	var m map[string][][]int = make(map[string][][]int)
	return helper(nums, &m)
}

func helper(nums []int, m *map[string][][]int) [][]int {
	l := len(nums)

	if l == 0 {
		return [][]int{}
	}

	if l == 1 {
		return [][]int{{nums[0]}, {}}
	}

	var ret [][]int
	// base
	ret = append(ret, nums)

	for i := 0; i < l; i++ {
		w := without(nums, i)
		sort.Ints(w)

		subresult, ok := (*m)[hash(w)]
		if ok {
			// already
			continue
		}
		subresult = helper(w, m)
		(*m)[hash(w)] = subresult
		ret = append(ret, subresult...)
	}

	var res [][]int = [][]int{{}}
	// remove dupes
	for _, ss := range ret {
		if len(ss) == 0 {
			continue
		}
		res = append(res, ss)
	}

	return res
}

func without(nums []int, index int) []int {
	var withoutIndex []int
	withoutIndex = append(withoutIndex, nums[0:index]...)
	withoutIndex = append(withoutIndex, nums[index+1:]...)
	return withoutIndex
}

func hash(arr []int) string {
	var s strings.Builder
	for _, x := range arr {
		fmt.Fprint(&s, x)
	}
	return s.String()
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 1}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{1, 2}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 3}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 4}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 4, 5}))
	fmt.Println("----------------------------")
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
