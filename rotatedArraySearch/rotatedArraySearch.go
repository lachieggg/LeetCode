
// search
func Search(nums []int, target int) int {
	prev := nums[0]

	if nums[0] > nums[-1] {
		// rotated
		rotationIndex := findOutOfOrder()
		// split into two arrays and do
		// bsearch on each

		// return

		return -1
	}

	bsearch(nums, target)

}

// bsearch
func bsearch(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return notFound
	}
	splitIndex := midpoint(float64(length))
	fnd := bsearch(nums[0:splitIndex], target)
	if fnd != notFound {
		return fnd
	}
	fnd = bsearch(nums[splitIndex:], target)
	if fnd != notFound {
		return fnd + splitIndex
	}

	return notFound
}

// findOutOfOrder find a pair of numbers that are out of order
// and then returns the index of the first number in the pair.
// Returns index of x, where x is left adjacent to y, and x > y
// Returns -1 if no such pair exists.
func findOutOfOrder(nums []int) int {
	var index int
	var direction int = 1
	var exponent int = 1
	var length int = len(nums)
	var overshot bool = false
	var velocity int = 1

	if debug {
		fmt.Printf("length = %d\n", length)
	}

	for {
		velocity = direction * powInt(2, exponent)
		index = index + velocity
		if index > length-1 {
			if overshot {
				return -1
			}
			overshot = true
			if debug {
				fmt.Printf("past end of array %d\n", index)
			}
			index = length - 1
			direction = -direction
			exponent = 1
		}
		if index < 0 {
			if overshot {
				return -1
			}
			overshot = true
			if debug {
				fmt.Printf("before beginning of array, %d\n", index)
			}
			index = 0
			direction = -direction
			exponent = 1
		}
		if debug {
			fmt.Printf("index = %d\n", index)
		}
		if nums[0] > nums[index] {
			// overshot or found?
			if nums[index-1] > nums[index] {
				return index - 1 // found
			}
			// overshot, reset in opposite direction
			direction = -direction
			exponent = 1
		} else {
			// accelerate
			exponent += 1
		}
	}

}

// powInt
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
