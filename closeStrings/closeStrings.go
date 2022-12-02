package main

import (
	"fmt"
	"strings"
)

const REPLACE_ALL = -1

// Keeps track of the strings that
// have already been checked
var checked []string

// recurse
func recurse(s1 string, s2 string) bool {
	checked = append(checked, s2)

	if(s1 == s2) { 
		return true
	}

	for i, _ := range s1 {
		fmt.Println(i)
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
	x := closeStrings("123", "1234")
	fmt.Println(x)

	y := closeStrings("123", "123")
	fmt.Println(y)

	a := "1112"
	b := "3332"
	fmt.Println(b)

	c := replace(a, "1", "2")

	fmt.Println(c)

	d := swap(a, 3, 1)
	fmt.Println(d)

	closeStrings(a, b)
}

// main
func main() {
	test()
}
