package main

// prints
func prints(nums []int) {
	for _, x := range nums {
		print(x)
	}
}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	length := len(nums)
	arr := make([]int, length)
	copy(arr, nums)
	nums[0] = nums[length-1]
	copy(nums[1:], arr[0:length-1])
	rotate(nums, k-1)
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	rotate(x, 1)
	prints(x)
}
