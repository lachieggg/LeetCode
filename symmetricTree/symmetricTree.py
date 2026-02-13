#!/usr/bin/env python3

DEBUG = False

class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetricArray(self, array):
        """Determine whether an array is symmetric"""
        if(len(array) % 2 != 0):
            return False
        middle = len(array)//2
        arrayOne = array[0:middle]
        arrayTwo = array[middle:]
        if(DEBUG): print("Array one is {}, array two is {}".format(arrayOne, arrayTwo))
        
        if(arrayOne == arrayTwo[::-1]):
            if(DEBUG): print("Symmetric array")
            return True
        
        if(DEBUG): print("Non-symmetric array")
        if(DEBUG): print("Tree is not a mirror of itself")
        return False


    def isSymmetric(self, tree):
        """
        Determine whether a binary tree is symmetric
        """
        depth = self.treeDepth(tree)
        self.getAllBinaryStrings(depth)
        treeArray = self.convertTreeToArray(tree)
        n = 1
        index = 0
        array = []
        while(1):
            index += n
            n = n*2
            if(index + n > len(treeArray)):
                return True
            if(DEBUG): print("Index is {}, n is {}".format(index, n))
            array = treeArray[index:index+n]
            if(DEBUG): print("Array is {}".format(array))
            if(not self.isSymmetricArray(array)):
                if(DEBUG): print("Array {} is not symmetric".format(treeArray))
                return False
                
    def treeDepth(self, tree):
        """Determine the depth of a given binary tree"""
        if(not tree):
            return 0

        return 1 + max(self.treeDepth(tree.left), self.treeDepth(tree.right))
    
    def convertTreeToArray(self, tree):
        """Convert a tree into an array"""
        array = []
        array.append(tree.val)

        leftArray = []
        rightArray = []

        # Example
        #
        #
        # root
        # left                # (0)
        # right               # (1)
        # left left           # (00)
        # left right          # (01)
        # right left          # (10)
        # right right         # (11)
        # left left left      # (000)
        # left left right     # (001)
        # right left left     # (100)
        # right left right    # (101)
        # right right right   # (111)

        treeDepth = self.treeDepth(tree)
        instructionSets = self.binaryStrings
        if(DEBUG): print(instructionSets)
        arr = self.traverse(tree, instructionSets)
        if(DEBUG): print("Converted tree to array, {}".format(arr))
        return arr
    
    def convertArrayToTree(self, array, tree, startIndex, length):
        """
        Convert an array to a tree

        Each layer we will be multiplying the starting index by 2, since
        when you go one layer deeper, you have double the number of nodes
        from the last layer

        The length is the length of the list and is the measure for when
        we terminate, since the process will be complete when we each the end of the list

        (Sourced from https://www.geeksforgeeks.org/construct-complete-binary-tree-given-array/)
        """
        if(startIndex < length):
            # Create a node with value equal to the value in the array at index 'startIndex'
            temp = TreeNode(array[startIndex], None, None)
            tree = temp

            tree.left = self.convertArrayToTree(array, tree.left, 2 * startIndex + 1, length)

            tree.right = self.convertArrayToTree(array, tree.right, 2 * startIndex + 2, length)
        
        return root
    
    def getNode(self, tree, instructionSet):
        """
        Returns a particular node that is reached by executing a given traversal instruction set
        """
        if(not tree):
            return None

        if(len(instructionSet) == 0):
            return tree.val

        # Right node
        if(eval(instructionSet[0])): 
            if(not tree.right):
                return None
            return self.getNode(tree.right, instructionSet[1:])
        
        if(not tree.left): 
            return None

        # Left node
        return self.getNode(tree.left, instructionSet[1:])
        
        
    def traverse(self, tree, instructionSets):
        """
        Traverse a binary tree using an instruction set

        An instruction set consists of a set of 1s and 0s, of length equal to the depth of the tree

        Zero indicates going down the left path of the tree, and one indicates going down the right path
        """
        array = []
        array.append(tree.val)
        if(len(instructionSets) == 0):
            return []
        for instructionSet in instructionSets:
            node = self.getNode(tree, instructionSet)
            array.append(node)

        return array

    def getAllBinaryStrings(self, n, i=1):
        """Efficiently generates the full set of binary strings of length <= n"""
        self.binaryStrings = ['0', '1']
        length = len(max(self.binaryStrings, key=len))
        for s in self.binaryStrings:
            if(len(s) >= n):
                return
            self.binaryStrings.append('0' + s)
            self.binaryStrings.append('1' + s)

        self.getAllBinaryStringsNew(n, i+1)

    def isPowerOfTwo(self, k):
        """Determines if a number k is a power of two, for k <= 2^50 ~= 1.12 quadrillion"""
        for x in range(50):
            if(2**x == k):
                return True
        return False
    
    def formatArrayAsTree(self, arr):
        """Formats an array of nodes in a tree pattern"""
        for i, x in enumerate(arr):
            if(self.isPowerOfTwo(i+1)):  
                print('\n')
            print(x, end=", ")

s = Solution()

# 1
# 2 4
# 3 1 36 7
#
# Not symmetric
# False

leftTree = TreeNode(2, TreeNode(3, None, None), TreeNode(1, None, None))
rightTree = TreeNode(4, TreeNode(36, None, None), TreeNode(7, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)

# 1
# 2 2
# 3 4 4 3
# 
# Symmetric
# True


leftTree = TreeNode(2, TreeNode(3, None, None), TreeNode(4, None, None))
rightTree = TreeNode(2, TreeNode(4, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)

# 1
# 2 2
# None 3 None 3
#
# Not symmetric
# False

leftTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
rightTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

result = s.isSymmetric(root)
print(result)

arr = [
    1, 
    74, 74, 
    68, -76, -76, 68, 
    -32, None, None, -85, None, -85, None, -32, 
    None, -69, 53, -80, -80, 53, -69, None, None, 96, -73, None, 62, None, None, 62, 
    None, -73, 96, None, 76, None, None, -49, None, -14, -14, None, -49, None, None, 76, 11, None, -93, None, None, 98, 98, None, None, -93, 11, None, 43, None, -60, -75, 
    -58, 46, 46, -58, -75, -60, None, 43, 69, None, None, 15, None, -12, None, -69, -54, None, None, -54, -69, None, -12, None, 15, None, None, 69, -35, None, 31, None, None, -57, None, 99, -85, None, None, -85, 99, 1, None, 20, None, None, 42, -12, None, None, -12, 42, None, None, 20, None, 1, None, -83, 62, None, None, 54, None, 
    24, 94, None, None, 75, -81, None, -69, None, 81, 38, None, 52, 56, None, 95, 94, 94, 95, 1, None, 56, 52, None, 38, 81, None, -69, None, -81, 75, None, None, 94, None, 95, -95, None, -13, None, 32, None, 80, None, -50, None, None, -68, None, 51, None, 5, -13, None, None, -13, 5, None, 51, None, -68, None, None, -50, None, 80, None, 32, None, -13, None, -95, 95, None, None, 12, None, 77, 13, None, None, -90, None, -92, None, 59, 20, None, -11, -12, -98, None, -47, None, None, -47, None, -98, -12, -11, None, 20, 59, None, -92, None, -90, None, None, 13, 77, None, 12,
]

tree = s.convertArrayToTree(arr, None, 0, len(arr))

result = s.isSymmetric(tree)
print(result)

arr = [
    9,
    14, 14, 
    74, None, None, 74, 
    None, 12, 12, None, 63, None, None, 63, 
    -8, None, None, -8, -53, None, None, -53, None, -96, -96, None, -65, None, None, -65, 
    98, None, None, 98, 50, None, None, 50, None, 91, 91, None, 41, -30, -30, 41, None, 86, None, -36, -36, None, 86, None, -78, None, 9, None, None, 9, None, -78, 
    47, None, 48, -79, -79, 48, None, 47, -100, -86, None, 47, None, 67, 67, None, 47, None, -86, -100, -28, 11, None, 56, None, 30, None, 64, 64, None, 30, None, 56, None, 11, -28, 43, 54, None, -50, 44, -58, 63, None, None, -43, -43, None, None, 63, -58, 44, -50, None, 54, 43
]

tree = s.convertArrayToTree(arr, None, 0, len(arr))
result = s.isSymmetric(tree)

print(result)