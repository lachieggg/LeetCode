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
		sleep(0.05)
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

def old():
	for j,curr in enumerate(sorted):
		print("Curr = {}".format(curr))
		# Between
		if(x >= prev and x <= curr):
			print("Found")
			sorted = sorted.insert(j, x)
			print(sorted)
			continue
		#if(x > sorted[]):

		prev = curr



def insertionSort(A):
	print(A)
	first = A[0]
	sorted = [first]

	for i,x in enumerate(A[1:]):
		print("x = {}".format(x))
		print("Sorted = {}".format(sorted))
	
		j = 0
		while(j <= len(sorted)):
			prev = sorted[j-1]

			# Start position
			if(j == 0):
				if(x < sorted[j]):
					sorted.insert(j, x)
					break

			if(j == len(sorted)):
				sorted.insert(j, x)
				break

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
