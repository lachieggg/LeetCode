#!/usr/bin/env python

from sys import argv

LOGGING = False

def collatz(n):
	"""Test the collatz conjecture for some positive integer n"""
	if(LOGGING): print(n)
	if n == 1:
		print("Complete.")
		return
	if(n%2 == 0): # even
		collatz(n//2)
	else: # odd
		collatz(3*n+1)

def main():
	if(len(argv) != 2):
		n = int(input("Enter a number: "))
	else:
		n = int(argv[1])
	if(LOGGING): print(n)
	collatz(n)

if(__name__ == "__main__"):
	main()
