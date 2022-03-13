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
        n = 1
        index = 0
        array = []
        while(1):
            index += n
            n = n*2
            if(index + n > len(tree)):
                return True
            print("Index is {}, n is {}".format(index, n))
            array = tree[index:index+n]
            print("Array is {}".format(array))
            if(not self.isSymmetricArray(array)):
                return False
            
            print("\n\n")
