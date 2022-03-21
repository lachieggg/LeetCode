#!/usr/bin/env python3

from random import randint 

MATRIX_SIZE = 6
MAX_VALUE = 10

class Node:
	def __init__(self, id=None, val=None, connections=[]):
		self.id = id
		self.val = val
		self.connections = connections
	
	def represent(self):
		connectedNodeIds = []
		for connection in self.connections:
			connectedNodeIds.append(connection.id)
		return "Node {} is connected to {}".format(self.id, connectedNodeIds)

	def __repr__(self):
		return str(self.id)

class Graph:
	def __init__(self, matrix=None, nodes=None):
		if(not matrix):
			self.matrix = self.generateMatrix()
		if(not nodes):
			self.nodes = self.generateNodes()
		
		self.visited = []
		
		print("Matrix validity: {}".format(self.validate(self.matrix)))

	def initializeNodes(self):
		for index, row in self.nodes:
			node = Node(None, transformRow(row))
	
	def transformRow(self, row):
		"""Turn a row into a set of node ids"""
		r = []
		for index, y in enumerate(row):
			if(not y):
				continue
			r.append(index)
		
		return r
	
	def explore(self, node):
		if(node.id in self.visited):
			return
		self.visited.append(node.id)
		for connection in node.connections:
			self.explore(connection)

	
	def getConnections(self, nodeId):
		"""Get the node connections as a list of ids"""
		row = self.matrix[nodeId]
		connections = self.transformRow(row)
		return connections
	
	def validate(self, matrix):
		# Every row must be of length N
		for row in matrix:
			if(not len(matrix) == len(row)):
				return False
		
		# Every node must be reachable from itself
		for index, row in enumerate(matrix):
			if(not row[index] == 1):
				return False
		
		# Every edge must have a twin edge
		for x in range(len(matrix)):
			for y in range(len(matrix)):
				if(matrix[x][y] and not matrix[y][x]):
					return False
		
		return True

	def representGraph(self):
		# Print the matrix
		for row in self.matrix:
			for element in row:
				print(element, end=' ')
			print()
		print()
		
		# Print the values
		for nodeId in self.nodes.keys():
			print("Node {} connections: {}".format(nodeId, self.nodes.get(nodeId).represent()))
		
	def generateMatrix(self):
		m = [[0 for x in range(MATRIX_SIZE)] for y in range(MATRIX_SIZE)]

		for x in range(MATRIX_SIZE):
			for y in range(MATRIX_SIZE):
				if(x == y):
					m[x][y] = 1
				if(randint(0,1)):
					m[x][y] = 1
					m[y][x] = 1

		return m

	def generateNodes(self):
		n = {}
		for nodeId in range(MATRIX_SIZE):
			val = randint(0, MAX_VALUE)
			n[nodeId] = Node(nodeId, val, [])

		for nodeId in n.keys():
			for connectedNodeId in self.getConnections(nodeId):
				n[nodeId].connections.append(n[connectedNodeId])

		return n

if(__name__ == "__main__"):
	g = Graph()
	g.representGraph()
	node = g.nodes[0]
	g.explore(node)
	print(g.visited)