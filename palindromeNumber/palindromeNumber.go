package main

import "strconv"


// strIsPalindrome
//
// Given a string y, return true
// if y is a palindrome, and false
// otherwise.
func strIsPalindrome(y string) bool {
	r := []rune(y)
	for i, _ := range y {
		if r[i] != r[len(y) - 1 - i] {
			return false
		}
	}
	return true
}

// isPalindrome
//
// Given an integer x, return true 
// if x is a palindrome, and false 
// otherwise.
func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return strIsPalindrome(s)
}

func main() {
	test()
}

func test() {
	isPalindrome(12)
}