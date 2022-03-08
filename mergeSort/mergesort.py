#!/usr/bin/env python

from math import floor

def merge(arr1, arr2):
	if(len(arr1) == 0):
		return arr2
	if(len(arr2) == 0):
		return arr1

	if(arr1[0] <= arr2[0]):
		return [arr1[0]] + merge(arr1[1:], arr2)
	else:
		return [arr2[0]] + merge(arr1, arr2[1:])

def mergeSort(arr):
	print("Input array {}".format(arr))
	if(len(arr) > 1):
		n = len(arr)
		arr1 = arr[0:floor(n/2)]
		arr2 = arr[floor(n/2):n]
		return merge(mergeSort(arr1), mergeSort(arr2))
	else:
		return arr

def test():
	arr = [3,6,1,4,7,1,6]
	sorted = mergeSort(arr)
	print(sorted)


if __name__ == "__main__":
	test()
