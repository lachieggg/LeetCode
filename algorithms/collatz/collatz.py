#!/usr/bin/env python

def collatz(n):
	"""Collatz conjecture for some positive integer n"""
	print(n)

	if n == 1:
		return
	elif n%2 == 0:
		collatz(n//2)
	else:
		collatz(3*n+1)

def main():
	collatz(1010101010)

if(__name__ == "__main__"):
	main()
