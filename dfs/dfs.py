#!/usr/bin/env python3

MATRIX_SIZE = 6

class Node:
	def __init__(self, id=None, val=None, neighbours=[]):
		self.val = val
		self.neighbours = neighbours

class Graph:
	def __init__(self, matrix, values):
		if(not self.validate(matrix)):
			exit("Invalid Matrix")
		self.matrix = matrix
		self.nodes = []
		self.values = values

	
	def initializeNodes(self):
		# 0:    1 1 1 1 1 1
		# 1:    0 1 1 0 0 0
		# 2:    1 1 1 0 0 0
		# 3:    1 0 1 0 1 0
		# 4:    1 0 0 0 0 0
		# 5:    0 1 0 0 0 0
		for index, row in self.matrix:
			node = Node(None, transformRow(row))
	
	def transformRow(self, row):
		"""
		A row represents the nodes that a node is connected to
		This method will transform a row from a matrix into a set of ids that a the node is connected to 
		"""
		r = []
		for index, y in enumerate(row):
			if(not y):
				continue
			r.append(index)
		
		return r
	
	def validate(self, matrix):
		# Every row must be of length N
		for row in matrix:
			if(not len(matrix) == len(row)):
				return False
		
		# Every node must be reachable from itself
		for index, row in enumerate(matrix):
			if(not row[index] == 1):
				return False
		
		return True
	
def generateMatrix():
	m = []
	for x in range(MATRIX_SIZE):
		r = []
		for y in range(MATRIX_SIZE):
			if(y == x):
				r.append(1)
			else:
				r.append(0)
		m.append(r)
	return m

if(__name__ == "__main__"):
	m = generateMatrix()
	print(m)
