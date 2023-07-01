package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestRemoveInPlaceElem
func TestRemoveInPlaceElem(t *testing.T) {
	var nums []int = []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 4, 5, 5}
	x := removeElement(nums, 3)
	assert.Equal(t, 8, x)

	nums = []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5}
	x = removeElement(nums, 3)
	assert.Equal(t, 8, x)

	nums = []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5}
	x = removeElement(nums, 2)
	assert.Equal(t, 12, x)

	nums = []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5, 6, 6, 6}
	x = removeElement(nums, 2)
	assert.Equal(t, 15, x)

	nums = []int{3, 2, 2, 3}
	x = removeElement(nums, 3)
	assert.Equal(t, 2, x)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	x = removeElement(nums, 3)
	assert.Equal(t, 7, x)
}
