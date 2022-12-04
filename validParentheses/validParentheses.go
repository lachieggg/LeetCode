package main

import "fmt"
import "strings"

var open = []string{
	"(",
	"{", 
	"[",
}

var closed = []string{
	")",
	"}", 
	"]",
}

var opening = make(map[string]string)

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

	prev := ""
	i := 0
	for _, e := range s {
		c := string(e)
		if c == "}" || c == ")" || c == "]" {
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
	fmt.Println(isValid("()[]{}"))  // true
}