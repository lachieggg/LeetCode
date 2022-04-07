#!/usr/bin/env python3

### PAYPALISHIRING
###
### P   A   H   N
### A P L S I I G
### Y   I   R


### ABCDE
###
### A D C 
### B   E


### ABC
###
### A C
### B  

### ABCDEFGH
### 
### A C E G
### B D F H


DEBUG = False

class Solution:
	def __init__(self, s=None, numRows=None):
		self.s = s
		self.numRows = numRows
	
	def getNumberOfColumns(self):
		"""Calculate the number of columns in an output given an input"""
		cols, count, skip, length = 1, 0, 0, len(self.s)
		for c in self.s:
			if(skip):
				skip -= 1
				continue
			count += 1
			if(count == self.numRows):
				count = 0
				cols += 1 + self.numRows - 2
				skip = self.numRows - 2

		#print(cols)
		return cols
	
	def convertSolutionRowsToString(self):
		"""Converts a solution row to a string"""
		string = ''
		for rowIndex, row in enumerate(self.rows):
			for columnIndex in range(len(row)):
				if(self.rows[rowIndex][columnIndex] == '_'):
					continue
				string += self.rows[rowIndex][columnIndex]
		
		return string
	
	def convert(self, s=None, numRows=None):
		if(not s):
			s = self.s 
		if(not numRows):
			numRows = self.numRows
		
		self.s = s
		self.numRows = numRows

		if(DEBUG): print(s)
		self.algorithm(s, numRows)
		if(DEBUG): print(self.rows)
		s = self.convertSolutionRowsToString()
		return s
	
	def algorithm(self, s, numRows):
		numCols = self.getNumberOfColumns()
		rows = [['_' for x in range(numCols)] for x in range(numRows)]

		index = 0
		startHeight = numRows - 2
		zag = False
		#print(s)

		if(len(s) == 1):
			self.rows = [[s[0]]]
			return
		
		if(numRows == 1):
			self.rows = [x for x in s]
			print(self.rows)
			return
		
		if(numRows == 2):
			first, second = [], []
			for index in range(0, len(s), 2):
				first.append(s[index])
			for index in range(1, len(s), 2):
				second.append(s[index])

			self.rows = [first, second]
			return
		
		height = startHeight

		for col in range(0, numCols, 1):
			self.rows = rows
			if(zag):
				if(DEBUG): print("Zag")
				if(height == 1):
					if(DEBUG): print("Height == 1")
					zag = False
					try:
						rows[height][col] = s[index]
						if(DEBUG): print(rows[height][col])
					except IndexError as e:
						if(DEBUG): print("Index Error")
						if(DEBUG): print("Col = {}".format(col))						
						self.rows = rows
						return
					index += 1
					height = startHeight
					continue
				try:
					rows[height][col] = s[index]
					if(DEBUG): 	print(rows[height][col])
				except IndexError as e:
					if(DEBUG): print("Index Error")
					if(DEBUG): print("Col = {}".format(col))
					self.rows = rows
					return
				index += 1
				height -= 1
			else:
				if(DEBUG): print("Zig")
				for r in range(0, numRows, 1):
					try:
						rows[r][col] = s[index]
						if(DEBUG): print(rows[r][col])
					except IndexError as e:
						if(DEBUG): print("Index Error")
						if(DEBUG): print("Col = {}".format(col))
						if(DEBUG): print("Row = {}".format(r))
						if(DEBUG): print("Index = {}".format(index))
						self.rows = rows
						return
					index += 1
					zag = True
					height = startHeight

s = Solution("ABC", 2)
k = s.getNumberOfColumns()
print(k)
m =  s.convert()
print(m)
