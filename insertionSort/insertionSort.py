#!/usr/bin/env python3

import random
from time import sleep

LENGTH = 8
MAX_VALUE = 100

def swap(arr, i, j):
	temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr

def visualize(arr):
	"""
	Visualize an array of values
	"""
	for x in arr:
		print(x, end=' ')
		# Spacing for single value characters
		if(len(str(x)) == 1): print(' ', end='')
		for y in range(x):
			print("|", end='')
		print("\n")
	print("\n")

def insert(A, elem, index):
	B = A[0:index]
	B.append(elem)
	for x in A[index:]:
		B.append(x)
	return B

def insertionSort(L):
	print(L)
	first = L[0]
	sorted = [first]

	for x in L[1:]:	
		print(sorted)

		# Start position
		index = 0
		if(x < sorted[0]):
			sorted.insert(0, x)
			continue
	
		# End position
		index = len(sorted)
		if(x > sorted[index-1] and index > 0):
			sorted.insert(index, x)
			continue
		
		# Middle position(s)
		j = 0
		while(j <= len(sorted)):
			prev = sorted[j-1]
			curr = sorted[j]

			if(x >= prev and x <= curr):
				sorted.insert(j, x)
				break

			j += 1
	
	return sorted


def initArray():
	arr = [x for x in range(LENGTH)]
	for i in range(LENGTH):
		arr[i] = random.randint(0, MAX_VALUE)
	
	return arr

def main():
	arr = initArray()
	print(arr)
	arr = insertionSort(arr)
	print(arr)

if(__name__ == "__main__"):
	print('\n')
	main()
	print('\nDone.')
