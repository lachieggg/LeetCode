package main

import "fmt"

func main() {
	var x bool

	x = isValid("()")
	fmt.Println(x)

	x = isValid("[([]])")
	fmt.Println(x)

}

const NOT_FOUND =      -1
const EMPTY         =  "EMPTY"
const OPENING_PAREN =  "("
const CLOSING_PAREN =  ")"
const OPENING_BRACE =  "{"
const CLOSING_BRACE =  "}"
const OPENING_SQ =     "["
const CLOSING_SQ =     "]"

func find(s string, l []string) int {
	for i, c := range l {
		if s == c {
			return i
		}
	}
	return NOT_FOUND
}

func closing(c string) bool {
	return c == CLOSING_BRACE || c == CLOSING_PAREN || c == CLOSING_SQ
}

func isValid(s string) bool {
	openings := []string{
		OPENING_PAREN,
		OPENING_BRACE, 
		OPENING_SQ,
	}

	closings := []string{
		CLOSING_PAREN, 
		CLOSING_BRACE,
		CLOSING_SQ,
	}

    m := make(map[string]int)

	for _, b := range openings {
		m[b] = 0
	}


	prev := EMPTY

	// Iterate over the 
	// string and work out the 
	// totals for each bracket
	//
	// Also work out whether there are any
	// issues with ordering 
	//
	for _, c := range []rune(s) {
		z := string(c)
		if(closing(z)) {
			if prev == EMPTY {
				return false
			} else if find(prev, closings) != NOT_FOUND { 
				// If it is a closing
				// we can skip checking
			} else if find(z, closings) != find(prev, openings) {
				return false
			}
			index := openings[find(z, closings)]
			m[index] -= 1
			if m[index] < 0 {
				return false
			}
		} else {
			m[z] += 1
			if m[z] < 0 {
				return false
			}
		}

		prev = string(c)
	}

	if  m[string(openings[0])] - m[string(closings[0])] != 0 || 
	    m[string(openings[1])] - m[string(closings[1])] != 0 ||
		m[string(openings[2])] - m[string(closings[2])] != 0 {
		// fmt.Println("Sum is wrong")
		return false
	}



	return true
}