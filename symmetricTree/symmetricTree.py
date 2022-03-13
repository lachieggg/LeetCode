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
        Determine whether a binary tree is symmetric
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
    
    def convertArrayToTree(self, array, tree, startIndex, length):
        """
        Convert an array to a tree

        Algorithm sourced from GeeksToGeeks

        Each layer we will be multiplying the starting index by 2, since
        when you go one layer deeper, you have double the number of nodes
        from the last layer

        The length is the length of the list and is the measure for when
        we terminate, since the process will be complete when we each the end of the list
        """
        if(startIndex < length):
            # Create a node with value equal to the value in the array at index 'startIndex'
            temp = TreeNode(array[startIndex], None, None)
            tree = temp

            tree.left = self.convertArrayToTree(array, tree.left, 2 * startIndex + 1, length)

            tree.right = self.convertArrayToTree(array, tree.right, 2 * startIndex + 2, length)
        
        return root

    def printTree(self, node, level=0):
        # Sourced from StackOverflow
        if node != None:
            self.printTree(node.left, level + 1)
            print(' ' * 4 * level + '-> ' + str(node.val))
            self.printTree(node.right, level + 1)
    
    def getNode(self, tree, instructionSet):
        if(not tree):
            return None

        if(len(instructionSet) == 0):
            return tree.val

        # Evalulate the right node
        if(eval(instructionSet[0])): 
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

#result = s.isSymmetric(root)
#print(result)

# 1
# 2 2
# 3 4 4 3

leftTree = TreeNode(2, TreeNode(3, None, None), TreeNode(4, None, None))
rightTree = TreeNode(2, TreeNode(4, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

#result = s.isSymmetric(root)
#print(result)

# 1
# 2 2
# None 3 None 3

leftTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
rightTree = TreeNode(2, TreeNode(None, None, None), TreeNode(3, None, None))
root = TreeNode(1, leftTree, rightTree)

#result = s.isSymmetric(root)



arr = [
    1,74,74,68,-76,-76,68,-32,None,None,
    -85,None,-85,None,-32,None,-69,53,-80,-80,53,
    -69,None,None,96,-73,None,62,None,None,62,None,
    -73,96,None,76,None,None,-49,None,-14,-14,None,
    -49,None,None,76,11,None,-93,None,None,98,98,None,
    None,-93,11,None,43,None,-60,-75,-58,46,46,-58,-75,
    -60,None,43,69,None,None,15,None,-12,None,-69,-54,
    None,None,-54,-69,None,-12,None,15,None,None,69,-35,
    None,31,None,None,-57,None,99,-85,None,None,-85,99,
    1,None,20,None,None,42,-12,None,None,-12,42,None,None,
    20,None,1,None,-83,62,None,None,54,None,24,94,None,None,
    75,-81,None,-69,None,81,38,None,52,56,None,95,94,94,95,
    None,56,52,None,38,81,None,-69,None,-81,75,None,None,94,
    None,95,-95,None,-13,None,32,None,80,None,-50,None,None,
    -68,None,51,None,5,-13,None,None,-13,5,None,51,None,-68,
    None,None,-50,None,80,None,32,None,-13,None,-95,95,None,
    None,12,None,77,13,None,None,-90,None,-92,None,59,20,None,
    -11,-12,-98,None,-47,None,None,-47,None,-98,-12,-11,None,
    20,59,None,-92,None,-90,None,None,13,77,None,12
]

tree = s.convertArrayToTree(arr, None, 0, len(arr))
result = s.isSymmetric(tree)
print(result)

arr = [
    9,14,14,74,None,None,74,None,12,12,None,63,None,None,63,-8,None,
    None,-8,-53,None,None,-53,None,-96,-96,None,-65,None,None,-65,98,None,
    None,98,50,None,None,50,None,91,91,None,41,-30,-30,41,None,86,None,-36,
    -36,None,86,None,-78,None,9,None,None,9,None,-78,47,None,48,
    -79,-79,48,None,47,-100,-86,None,47,None,67,67,None,47,None,-86,-100,
    -28,11,None,56,None,30,None,64,64,None,30,None,56,None,11,-28,43,54,None,
    -50,44,-58,63,None,None,-43,-43,None,None,63,-58,44,-50,None,54,43
]


tree = s.convertArrayToTree(arr, None, 0, len(arr))
result = s.isSymmetric(tree)
print(result)