package main

import (
	"fmt"
)

var best int

func lengthOfLongestSubstring(s string) int {
	fmt.Println("---------------------------------")
	fmt.Println(s)
	best = 0
	var used = make(map[byte]int)

	i := llss(s, 0, used, 0)
	fmt.Println("---------------------------------")
	return i
}

func llss(s string, currentIndex int, used map[byte]int, startIndex int) int {
	fmt.Printf("s = [%v]\n", s[startIndex:currentIndex])
	if currentIndex == len(s) {
		return currentIndex - startIndex
	}
	lastUsedAtIndex, ok := lookup(used, s, currentIndex)
	used[s[currentIndex]] = currentIndex
	if ok {
		updateBest(currentIndex - startIndex)
		updateStartIndex(lastUsedAtIndex+1, &startIndex)
		// found a previous instance
	}

	k := llss(s, currentIndex+1, used, startIndex)
	updateBest(k)
	return best
}

func updateStartIndex(potential int, currStart *int) {
	if potential > *currStart {
		*currStart = potential
	}
}

func updateBest(potential int) {
	if potential > best {
		best = potential
	}
}

func lookup(used map[byte]int, s string, index int) (int, bool) {
	lastUsed, ok := used[s[index]]
	return lastUsed, ok
}

func printMap(m map[byte]int) {
	for k, v := range m {
		fmt.Printf("%c -> %d\n", k, v)
	}
}

func highest(m map[string]int) int {
	highest := 0
	for _, k := range m {
		if k > highest {
			highest = k
		}
	}
	return highest
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("a"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("aa"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("aaa"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("aaaa"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("aaaaa"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("aaaaab"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("abba"))
}
