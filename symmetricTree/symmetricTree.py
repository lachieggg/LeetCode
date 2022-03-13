#!/usr/bin/env python3

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetricArray(self, array):
        if(len(array) % 2 != 0):
            return False
        middle = len(array)//2
        arrayOne = array[0:middle]
        arrayTwo = array[middle:]
        print("Array one is {}, array two is {}".format(arrayOne, arrayTwo))
        
        if(arrayOne == arrayTwo[::-1]):
            print("Symmetric array")
            return True
        
        print("Non-symmetric array")
        print("Tree is not a mirror of itself")
        return False


    def isSymmetric(self, tree):
        """
        Determine whether a binary tree, expressed as an array, is symmetric
        """
        treeArray = self.convertTreeToArray(tree)
        n = 1
        index = 0
        array = []
        while(1):
            index += n
            n = n*2
            if(index + n > len(treeArray)):
                return True
            print("Index is {}, n is {}".format(index, n))
            array = treeArray[index:index+n]
            print("Array is {}".format(array))
            if(not self.isSymmetricArray(array)):
                print("Array {} is not symmetric".format(treeArray))
                return False
                
    def treeDepth(self, tree):
        if(not tree):
            return 0

        return 1 + max(self.treeDepth(tree.left), self.treeDepth(tree.right))
    
    def convertTreeToArray(self, tree):
        array = []
        array.append(tree.val)

        leftArray = []
        rightArray = []

        # root
        # left                #(0)
        # right               #(1)
        # left left           #(00)
        # left right          #(01)
        # right left          #(10)
        # right right         #(11)
        # left left left      #(000)
        # left left right     #(001)
        # right left left     #(100)
        # right left right    #(101)
        # right right right   #(111)

        treeDepth = self.treeDepth(tree)
        instructionSets = self.getAllBinaryStrings(treeDepth)
        print(instructionSets)
        arr = self.traverse(tree, instructionSets)
        print("Converted tree to array, {}".format(arr))
        return arr
    
    def getNode(self, tree, instructionSet):
        if(not tree):
            return None

        if(len(instructionSet) == 0):
            return tree.val

        if(eval(instructionSet[0])): # Evalulate the right node
            if(not tree.right):
                return None
            return self.getNode(tree.right, instructionSet[1:])
        
        if(not tree.left): 
            return None

        # Evalulate the left node
        return self.getNode(tree.left, instructionSet[1:])
        
        
    def traverse(self, tree, instructionSets):
        array = []
        array.append(tree.val)
        if(len(instructionSets) == 0):
            return []
        for instructionSet in instructionSets:
            node = self.getNode(tree, instructionSet)
            array.append(node)

        return array

    def getBinaryStrings(self, length):
        """Returns all binary strings of particular length"""
        print("Length is {}".format(length))
        strings = []
        if(length == 1):
            return ['0', '1']
        for s in self.getBinaryStrings(length-1):
            strings.append("0" + s)
            strings.append("1" + s)
        return sorted(strings)

    def getAllBinaryStrings(self, n):
        """Returns all binary strings of length less than n"""
        strings = []
        for l in range(1, n+1):
            print("l is {}".format(l))
            for string in self.getBinaryStrings(l):
               strings.append(string)
        return strings


s = Solution()

# 1
# 2 4
# 3 1 36 7

leftTree = TreeNode(2, TreeNode(3, None, None), TreeNode(1, None, None))
rightTree = TreeNode(4, TreeNode(36, None, None), TreeNode(7, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)

# 1
# 2 2
# 3 4 4 3

leftTree = TreeNode(2, TreeNode(3, None, None), TreeNode(4, None, None))
rightTree = TreeNode(2, TreeNode(4, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)

# 1
# 2 2
# null 3 null 3

leftTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
rightTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)
