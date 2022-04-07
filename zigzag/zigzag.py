#!/usr/bin/env python3

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

		self.algorithm(s, numRows)
		s = self.convertSolutionRowsToString()
		return s
	
	def algorithm(self, s, numRows):
		numCols = self.getNumberOfColumns()
		rows = [['_' for x in range(numCols)] for x in range(numRows)]

		index = 0
		startHeight = numRows - 2
		zag = False

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
				if(height == 1):
					zag = False
					try:
						rows[height][col] = s[index]
					except IndexError as e:
						self.rows = rows
						return
					index += 1
					height = startHeight
					continue
				try:
					rows[height][col] = s[index]
				except IndexError as e:
					self.rows = rows
					return
				index += 1
				height -= 1
			else:
				for r in range(0, numRows, 1):
					try:
						rows[r][col] = s[index]
					except IndexError as e:
						self.rows = rows
						return
					index += 1
					zag = True
					height = startHeight

s = Solution("ABC", 2)
k = s.getNumberOfColumns()
m =  s.convert()
print(m)
