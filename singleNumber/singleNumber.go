package main

import "fmt"

const NF = -1

func print(nums map[int]int) {
	for a, x := range nums {
		fmt.Printf("%d : %d\n", a, x)
		// fmt.Println(x)
	}
}

func singleNumberI(nums []int) int {
    singles := make(map[int]int)

	for _, x := range nums {
		singles[x] += 1
		if singles[x] == 2 {
			delete(singles, x)
		}
	}

	var k int 

	for key, _ := range singles {
		k = key
	}

	return k
}


func singleNumberII(nums []int) int {
    singles := make(map[int]int)

	for _, x := range nums {
		singles[x] += 1
	}

	var k int 

	for key, _ := range singles {
		if singles[k] == 1 {
			k = key
			break
		}
	}

	return k
}



func singleNumberIII(nums []int) int {
    singles := make(map[int]int)

	for _, x := range nums {
		singles[x] += 1
	}

	for key, _ := range singles {
		if singles[key] == 1 {
			return key
		}
	}

	return NF
}

func main() {
	x := [...]int{1, 2, 2, 3, 4, 1, 3, 4, 5}
	y := singleNumberI(x[:])
	fmt.Println(y)
}