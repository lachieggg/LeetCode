package main

// removeElement returns the number of elements
// in nums with values not equal to val.
func removeElement(nums []int, val int) int {
	counter := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		index := i - counter
		if index == length-counter {
			break
		}
		num := nums[index]
		if num == val {
			counter += 1
			bubbleToEnd(&nums, index)
		}
	}
	return length - counter
}

// bubbleToEnd moves the element at the i'th
// index to the end of the array, retaining
// current order
func bubbleToEnd(arr *[]int, i int) {
	nums := *arr
	end := len(nums) - 1
	curr := i
	for curr != end {
		bubbleUp(arr, curr)
		curr += 1
	}
}

// bubbleUp
func bubbleUp(arr *[]int, index int) {
	nums := *arr
	nums[index], nums[index+1] = nums[index+1], nums[index]
}
