package main

import (
	"fmt"
)

// lengthOfLastWord_
// O(N) time complexity
func lengthOfLastWord_(s string) int {
	var curr rune
	l := 0

	for _, v := range s {
		if v == ' ' {
			curr = v
			continue
		}

		if curr != ' ' {
			l += 1
		}

		if curr == ' ' {
			l = 1
		}

		curr = v
	}

	return l
}

// lengthOfLastWord
// O(n) time, O(1) space
// [Sonnet 4.5] optimised solution
func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	// Skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Count the last word
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	fmt.Printf("%d\n", lengthOfLastWord("Hello World"))
	fmt.Printf("%d\n", lengthOfLastWord("luffy is still a joyboy"))
	fmt.Printf("%d\n", lengthOfLastWord("   fly me   to    the moon  "))
}
