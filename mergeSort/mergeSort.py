#!/usr/bin/env python

from math import floor
import random

LENGTH = 33
MAX_VALUE = 100
CONSOLE_OUTPUT = False

def merge(A, B):
	if(len(A) == 0):
		return B
	if(len(B) == 0):
		return A

	if(A[0] <= B[0]):
		return [A[0]] + merge(A[1:], B)
	else:
		return [B[0]] + merge(A, B[1:])

def initArray():
	arr = [x for x in range(LENGTH)]
	for i in range(LENGTH):
		arr[i] = random.randint(0, MAX_VALUE)

	print(arr)
	return arr

def mergeSort(arr):
	if(CONSOLE_OUTPUT): print("Input array {}".format(arr))
	if(len(arr) > 1):
		n = len(arr)
		A = arr[0:floor(n/2)]
		B = arr[floor(n/2):n]
		return merge(mergeSort(A), mergeSort(B))
	else:
		return arr

def test():
	arr = initArray()
	sorted = mergeSort(arr)
	print(sorted)


if __name__ == "__main__":
	test()
