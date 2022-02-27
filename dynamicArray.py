#!/usr/bin/env python3

import math
import os
import random
import re
import sys

#
# Complete the 'dynamicArray' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER n
#  2. 2D_INTEGER_ARRAY queries
#

# Constants
# 
# Query constants
FIRST_QUERY_INT = 1
SECOND_QUERY_INT = 2

# Query index constants
QUERY_TYPE_INDEX = 0
X_INDEX = 1
Y_INDEX = 2


def xor(x, y):
    print('{0:08b}'.format(x))
    print('{0:08b}'.format(y))

def dynamicArray(n, queries):
    # Construct the 2D array
    arr = []
    for index in range(n):
        arr.append([])
    
    lastAnswer = 0
    print(arr)
    x = queries[X_INDEX]
    y = queries[Y_INDEX]

    xor(x, y)
    if(queries[QUERY_TYPE_INDEX] == FIRST_QUERY_INT):
        # Execute first query algorithm
        pass
    elif(queries[QUERY_TYPE_INDEX] == SECOND_QUERY_INT):
        # Execute second query algorithm
        pass
    else:
        pass
    

dynamicArray(10, [1, 3, 5])
exit()



def main():
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    q = int(first_multiple_input[1])

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    result = dynamicArray(n, queries)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()

if __name__ == '__main__':
    main()
