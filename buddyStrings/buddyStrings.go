package main

import "fmt"

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

func contains(s string, arr []string) bool {
	for _, g := range arr {
		if s == g { 
			return true
		}
	}
	return false
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	s1 := []rune(s)
	s2 := []rune(goal)

	// Contains all the previous
	// chracters in s from the current
	// index j
	var prevContents []string
	
	// Contains all the differences
	// between the two arrays
	// (if any)
	//
	// Must be of exactly length 4 
	// in order to be a buddy
	var differences []string

	if(s == goal) {
		for _, c := range s {
			if(contains(string(c), prevContents)) {
				return true
			}
			prevContents = append(prevContents, string(c))
		}
	} else {
		for i, c := range s1 {
			if c != s2[i] {
				differences = append(differences, string(c))
				differences = append(differences, string(s2[i]))
			}
		}
		fmt.Println(differences)
		// For a buddy string there should be exactly 4 elements in differences,
		// two elements for each difference, i.e. 2 discrepencies in total
		//
		// In a buddy string the first character from array A that is different is 
		// the same as the second character from array B that is different, and the 
		// second character from array A that is different is the same as the first 
		// character from array B. 
		if(len(differences) == 4 && differences[0] == differences[3] && differences[1] == differences[2]) {
			return true
		}
	}
	return false
}

func main() {
	test()
}

func test() {
	x := buddyStrings("1234", "1324")
	fmt.Println(x)
}