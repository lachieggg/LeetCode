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

# To run in terminal, execute:
#
# export OUTPUT_PATH='/home/lachlan/Desktop/Git/LeetCode/output'
#
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

lastAnswers = [0]

def getLastAnswer():
    print("Last answer = {}".format(lastAnswers[len(lastAnswers)-1]))
    return lastAnswers[len(lastAnswers)-1]

def setLatestAnswer(latestAnswer):
    print("Appending last answer = {}".format(latestAnswer))
    lastAnswers.append(latestAnswer)

def binary_print(x, y):
    print('{0:08b}'.format(x))
    print('{0:08b}'.format(y))

def dynamicArray(n, queries):
    # Construct the 2D array
    arr = []
    for index in range(n):
        arr.append([])
    
    for qindex, query in enumerate(queries):
        print(arr)
        x = queries[qindex][X_INDEX]
        y = queries[qindex][Y_INDEX]

        print("x = {}".format(x))
        print("y = {}".format(y))

        idx = ((x ^ getLastAnswer()) % n)
        print("idx = {}".format(idx))

        if(queries[qindex][QUERY_TYPE_INDEX] == FIRST_QUERY_INT):
            print("Query type = 1")
            # Execute first query algorithm
            arr[idx].append(y)
        elif(queries[qindex][QUERY_TYPE_INDEX] == SECOND_QUERY_INT):
            # Execute second query algorithm
            print("Query type = 2")
            arrIndex = y % len(arr[idx])
            setLatestAnswer(arr[idx][arrIndex])
        else:
            pass
    
    returnArr = []
    for queryAnswer in arr[::-1]: # Reverse
        if(len(queryAnswer) == 0):
            continue
        returnArr.append(queryAnswer[len(queryAnswer)-1])
    return returnArr
        

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
