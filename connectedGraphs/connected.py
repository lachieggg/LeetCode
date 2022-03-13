#!/usr/bin/env python3

NOT_CONNECTED = 0

import os

class Graph:
	def __init__(self, vertices, order, size):
		self.vertices = vertices
		self.size = size
		self.order = order

	def reachable(self):
		pass

class Vertex:
	def __init__(self, label, Id, In=[], Out=[]):
		self.label = label
		self.Id = Id
		self.In = In
		self.Out = Out

def explore(vertex, visited, node_count):
	node_count -= 1;
	len_out = len(vertex.Out)
	visited[vertex.Id] = 1

	if(node_count == 0):
		print("Node count is 0")
		return visited
	else:
		for i in range(len_out):
			if not visited[i]:
				return explore(vertex.Out[i], visited, node_count)

	return NOT_CONNECTED

# Load and define the vertices and the edges
#
def load_vertices(lines, order):
	vertices = [Vertex(str(x), x) for x in range(order)]
	for line in lines[1:]:
		# (u, v)
		line = line.split(" ")
		uId = int(line[0])
		vId = int(line[1])
		# Add the out edge to u
		vertices[uId].Out.append(vertices[vId])
		# Add the in edge to v
		vertices[vId].In.append(vertices[uId])
	return vertices

# Load a graph of vertices and edges G(V, E)
#
# Order is the number of vertices
# Size is the number of edges
# Order = |V|
# Size = |E|
def load_graph(infile):
	lines = infile.readlines()
	order = int(lines[0])
	size = len(lines[order+1:])
	graph = Graph(load_vertices(lines, order), order, size)
	return graph

def main(graph):
	visited = [0] * graph.order
	for i in range(graph.order):
		result = explore(graph.vertices[i], visited, graph.order)
		if result == NOT_CONNECTED:
			print("Vertex {} is not connected to all other vertices".format(i))
			print("The graph is not connected.")
			return False

		print("Vertex {} is connected to all other vertices!".format(i))

	print("The graph is connected.")
	return True

if __name__ == "__main__":
	infile = open(os.getcwd() + "/dag.txt",'r')
	outfile = open(os.getcwd() + "/out.txt", 'w')
	graph = load_graph(infile)
	main(graph)
