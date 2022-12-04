package main

import "fmt"
import "strings"

const NOT_FOUND =      -1
const EMPTY         =  " "
const OPENING_PAREN =  "("
const CLOSING_PAREN =  ")"
const OPENING_BRACE =  "{"
const CLOSING_BRACE =  "}"
const OPENING_SQ =     "["
const CLOSING_SQ =     "]"


var open = []string{OPENING_PAREN,
	OPENING_BRACE, 
	OPENING_SQ,
}

var closed = []string{CLOSING_PAREN,
	CLOSING_BRACE, 
	CLOSING_SQ,
}

var opening = make(map[string]string)

func find(s string, l []string) int {
	for i, c := range l {
		if s == c {
			return i
		}
	}
	return 0
}

func isClosing(c string) bool {
	return c == CLOSING_BRACE || c == CLOSING_PAREN || c == CLOSING_SQ
}

// delete the character
// at i
func delete(s string, i int) string {
	var y []string
	for j, c := range s {
		if i != j {
			y = append(y, string(c))
		}
	}
	return strings.Join(y, "")
}

func subcall(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) == 2 {
		if opening[string(s[1])] == string(s[0]) && s[1] != s[0] {
			return true
		}
		return false
	}

	prev := EMPTY
	i := 0
	for _, e := range s {
		if isClosing(string(e)) {
			if prev != EMPTY {
				if opening[string(e)] == prev {
					// found opening brace
					// directly followed by 
					// corresponding closing brace
					// remove
					break
				} else {
					return false
				}
			}
		}
		i += 1
		prev = string(e)
	}

	x := delete(delete(s, i), i-1)

	return subcall(x)
}

func isValid(s string) bool {
	if len(s) %2 != 0 {
		return false
	}

	// maps open brackets
	// to their closing 
	// counterparts
	for i, b := range closed {
		opening[b] = open[i]
	}

	return subcall(s)
}


func main() {
	fmt.Println(isValid("[([]])"))  // false
	fmt.Println(isValid("{{}}"))    // true
	fmt.Println(isValid("{{}]}"))   // false
	fmt.Println(isValid("{[{}}"))   // false 
	fmt.Println(isValid("{({})}"))  // true
	fmt.Println(isValid("{{()}}"))  // true
	fmt.Println(isValid("(("))      // false

}