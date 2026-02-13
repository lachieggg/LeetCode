#!/usr/bin/env python3

import math
import os
import random
import re
import sys


# Constants
#
# Query constants
FIRST_QUERY_INT = 1
SECOND_QUERY_INT = 2

# Query index constants
QUERY_TYPE_INDEX = 0
X_INDEX = 1
Y_INDEX = 2

DEBUG = True

def dynamicArray(n, queries):
    global lastAnswer
    lastAnswer = 0

    global answerArray
    answerArray = []

    # Construct the 2D array
    arr = []
    for index in range(n):
        arr.append([])

    for qindex, query in enumerate(queries):
        if(DEBUG): print(arr)
        x = queries[qindex][X_INDEX]
        y = queries[qindex][Y_INDEX]

        if(DEBUG):
            print("x = {}".format(x))
            print("y = {}".format(y))

        idx = ((x ^ lastAnswer) % n)
        if(DEBUG): print("idx = {}".format(idx))

        if(queries[qindex][QUERY_TYPE_INDEX] == FIRST_QUERY_INT):
            if(DEBUG): print("Query type = 1")
            # Execute first query algorithm
            arr[idx].append(y)
        elif(queries[qindex][QUERY_TYPE_INDEX] == SECOND_QUERY_INT):
            # Execute second query algorithm
            if(DEBUG): print("Query type = 2")
            arrIndex = y % len(arr[idx])
            lastAnswer = arr[idx][arrIndex]
            answerArray.append(lastAnswer)

    return answerArray

def main():
    cwd = os.getcwd()
    fptr = open(cwd+ "/output", 'w')

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
