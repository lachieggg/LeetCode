#!/usr/bin/env python3

# Plan:
#
# Take the input square
# Compute all the current totals of each
# row/column/diagonal
#
# Sort the totals by number of occurrences
#
# Then manipulate the sqauare
# into fitting each of the magic numbers,
# beginning with the magic number that appears
# most frequently in the rows/columns/diagonals
# of S
#
#

import math
import os
import random
import re
import sys

#
# Complete the 'formingMagicSquare' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY s as parameter.
#

EMPTY_MATRIX = [[0,0,0],[0,0,0],[0,0,0]]

def transverse(matrix):
    newMatrix = EMPTY_MATRIX
    for x in range(3):
        for y in range(3):
            newMatrix[x][y] = matrix[y][x]
    return newMatrix

def formingMagicSquare(s):
    totals = []
    # Row totals
    for row in s:
        totals.append(sum(row))

    # Column totals
    for column in transverse(s):
        totals.append(sum(column))

    # Diagonal totals
    diagonalOne = [s[0][0], s[1][1], s[2][2]]
    diagonalTwo = [s[0][2], s[1][1], s[2][0]]

    totals.append(sum(diagonalOne))
    totals.append(sum(diagonalTwo))

    # Occurrences
    occurrences = {}
    for y in totals:
        occurrences[y] = totals.count(y)

def main():
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    s = []
    for _ in range(3):
        s.append(list(map(int, input().rstrip().split())))
    result = formingMagicSquare(s)
    fptr.write(str(result) + '\n')
    fptr.close()

if __name__ == '__main__':
    main()
