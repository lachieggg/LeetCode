package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	var l = len(s)

	if l == 1 {
		return s
	}

	if l == 2 && s[0] != s[1] {
		return s[0:1]
	}

	if l == 2 && s[0] == s[1] {
		return s
	}

	var best string
	for i := 0; i < l-1; i++ {
		for j := i + 1; j <= l; j++ {
			if j-i <= len(best) {
				continue
			}
			if isPalindrome(s[i:j]) {
				best = s[i:j]
			}
		}
	}

	return best
}

func isPalindrome(s string) bool {
	l := len(s)

	var i int
	if l%2 != 0 {
		// odd
		for i = 0; i < l/2+1; i++ {
			if s[l/2-i] != s[l/2+i] {
				return false
			}
		}
	} else {
		// even
		for i = 0; i < l/2; i++ {
			if s[l/2-i-1] != s[l/2+i] {
				return false
			}
		}
	}
	return true
}

func main() {
	inputs := []struct {
		input    string
		expected string
	}{
		{
			input:    "b",
			expected: "b",
		},
		{
			input:    "ab",
			expected: "a",
		},
		{
			input:    "aca",
			expected: "aca",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "aabbbaa",
			expected: "aabbbaa",
		},
		{
			input:    "acadaca",
			expected: "acadaca",
		},
		{
			input:    "sacacas",
			expected: "sacacas",
		},
		{
			input:    "ababababa",
			expected: "ababababa",
		},
		{
			input:    "acacacacaca",
			expected: "acacacacaca",
		},
		{
			input:    "ccccccccccccccc",
			expected: "ccccccccccccccc",
		},
		{
			input:    "cccccccccccccccccccccccccccccc",
			expected: "cccccccccccccccccccccccccccccc",
		},
		{
			input:    "ccccccccccccccccccccccccccccccccccccccccccccc",
			expected: "ccccccccccccccccccccccccccccccccccccccccccccc",
		},
	}

	for _, v := range inputs {
		if longestPalindrome(v.input) == v.expected {
			fmt.Println("✅")
		} else {
			fmt.Println("❌")
		}
		fmt.Printf("'%v' and '%v'\n", longestPalindrome(v.input), v.expected)
	}
}
