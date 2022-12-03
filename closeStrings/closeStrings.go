package main

import (
	"fmt"
	"strings"
)

const REPLACE_ALL = -1

// Keeps track of the strings that
// have already been checked
var checked []string

// contains
func contains(arr []string, s string) bool {
	for _, z := range arr {
		if z == s {
			return true
		}
	}
	return false
}

// recurse
// tries to turn s2 into s1
// by use of transformations
func recurse(s1 string, s2 string) bool {

	if s1 == s2 { 
		return true
	}
	checked = append(checked, s2)

	// Generate the set of possible
	// strings that can be created
	// at only this step, using 
	// transformations, and iteratively
	// check whether any are a solution
	var toCheck []string
	for i, _ := range s1 {
		for j, _ := range s1[0:i] {
			fmt.Printf("swapping %d with %d  ", i, j)
			n := swap(s2, i, j)
			fmt.Println(n)
			if n == s1 {
				return true
			}
			toCheck = append(toCheck, n)
		}
	}

	for i, _ := range s1 {
		for j, _ := range s1[0:i] {
			fmt.Printf("replacing %s with %s  ", string(s2[i]), string(s2[j]))
			n := replace(s2, string(s2[i]), string(s2[j]))
			fmt.Println(n)
			if n == s1 {
				return true
			}
			toCheck = append(toCheck, n)
		}
	}

	// Check the child strings of
	// each permutation at this step
	for _, s := range toCheck {
		//fmt.Println(s)
		if contains(checked, s) {
			fmt.Println("already checked")
			continue
		}
		checked = append(checked, s)
		if recurse(s1, s) {
			return true
		}
	}

	return false
}

// closeStrings
func closeStrings(s1 string, s2 string) bool {
	if len(s1) != len(s2) { 
		return false
	}

	return recurse(s1, s2)
}

// swap
// Given a string s and two indexes
// corresponding to characters swap
// two characters in s with one another 
func swap(s string, index1 int, index2 int) string {
	r := []rune(s)
	tmp := r[index1]
	r[index1] = r[index2]
	r[index2] = tmp
	return string(r)
}

// replace
// Given a string s and two strings s1 and s2,
// replace all instances of s1 with s2 
// and all instances of s2 with s1
func replace(s string, s1 string, s2 string) string {
	r := strings.NewReplacer(s1, s2, s2, s1)
	s = r.Replace(s)
	return s
}

// test
func test() {

	a := "123456"
	b := "653421"
	x := closeStrings(a, b)
	fmt.Println(x)
}

// main
func main() {
	test()
}
