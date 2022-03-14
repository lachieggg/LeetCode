#!/usr/bin/env python3

import random
from time import sleep

LENGTH = 8
MAX_VALUE = 100
VISUALIZE = False

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


def bubbleSort(arr):
	for j in range(len(arr)):
		if(VISUALIZE): visualize(arr)
		for i in range(1, len(arr)):
			if(arr[i-1] > arr[i]):
				arr = swap(arr, i, i-1)		
	
	return arr

def initArray():
	arr = [x for x in range(LENGTH)]
	for i in range(LENGTH):
		arr[i] = random.randint(0, MAX_VALUE)
	
	return arr

def main():
	arr = initArray()
	print(arr)
	arr = bubbleSort(arr)
	print(arr)

if(__name__ == "__main__"):
	print('\n')
	main()
	print('\nDone.')
