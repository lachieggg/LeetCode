#!/usr/bin/env python3

def treeDepth(self, tree):
	if(not tree):
		return 0

	depth = 1 + max(treeDepth(tree.left), treeDepth(tree.right))1~def treeDepth(self, tree):
	if(not tree):
		return 0

	depth = 1 + max(treeDepth(tree.left), treeDepth(tree.right))1~ def treeDepth(self, tree):
	if(not tree):
		return 0

	depth = 1 + max(treeDepth(tree.left), treeDepth(tree.right))
