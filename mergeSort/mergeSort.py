#!/usr/bin/env python

from math import floor

def merge(A, B):
	if(len(A) == 0):
		return B
	if(len(B) == 0):
		return A

	if(A[0] <= B[0]):
		return [A[0]] + merge(A[1:], B)
	else:
		return [A[0]] + merge(A, B[1:])

def mergeSort(arr):
	print("Input array {}".format(arr))
	if(len(arr) > 1):
		n = len(arr)
		A = arr[0:floor(n/2)]
		B = arr[floor(n/2):n]
		return merge(mergeSort(A), mergeSort(B))
	else:
		return arr

def test():
	arr = [3,6,1,4,7,1,6]
	sorted = mergeSort(arr)
	print(sorted)


if __name__ == "__main__":
	test()
