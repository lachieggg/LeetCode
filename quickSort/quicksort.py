#!/usr/bin/env python3

import random
from time import sleep
from math import floor 

LENGTH = 16
MAX_VALUE = 100
TIME_DELTA = 0.0125
VISUALIZE = True

def swap(arr, i, j):
	temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr

def visualize(arr):
	"""
	Visualize an array of values
	"""
	if(not VISUALIZE):
		return

	for x in arr:
		sleep(TIME_DELTA)
		print(x, end=' ')
		# Spacing for single value characters
		if(len(str(x)) == 1): print(' ', end='')
		for y in range(x):
			print("|", end='')
		print("\n")
	print("\n")
	

def initArray():
	arr = [x for x in range(LENGTH)]
	for i in range(LENGTH):
		arr[i] = random.randint(0, MAX_VALUE)
	
	return arr

def partition(A, lo, hi):
	"""
	Partition algorithm
	"""
	pivot = A[floor((hi + lo) / 2)]
	i = lo - 1
	j = hi + 1

	while(1):
		i = i + 1
		while(A[i] < pivot):
			i = i + 1
		
		j = j - 1
		while(A[j] > pivot):
			j = j - 1
		
		if(i >= j):
			return j
		
		swap(A, i, j)

def quickSort(A, lo, hi):
	"""
	Quicksort algorithm
	"""
	visualize(A)

	if(lo >= 0 and hi >= 0 and lo < hi):
		p = partition(A, lo, hi)
		A = quickSort(A, lo, p) 	# Sort left of pivot
		A = quickSort(A, p+1, hi)   # Sort right of pivot

	return A

def main():
	A = initArray()
	print(A)
	A = quickSort(A, 0, len(A) - 1)
	print(A)

if(__name__ == "__main__"):
	print('\n')
	main()
	print('\nDone.')
