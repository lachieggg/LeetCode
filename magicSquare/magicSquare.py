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

EMPTY_MATRIX = [[0,0,0],[0,0,0],[0,0,0]]
ROW = 'R'
COLUMN = 'C'
DIAGONAL = 'D'

# Count the number of occurrences
# of a value in a dictionary
def count(dict, value):
    total = 0
    for v in list(dict.values()):
        if(v == value):
            total += 1
    return total

def transverse(matrix):
    newMatrix = EMPTY_MATRIX
    for x in range(3):
        for y in range(3):
            newMatrix[x][y] = matrix[y][x]
    return newMatrix

def formingMagicSquare(s):
    sets = coords.values()                   # Sets
    coords = getSquareCoordsSet(s)           # Map of coordinates onto sets
    sums = getSetSums(sets)                  # Map of sets onto sum totals
    occurrences = getOccurrencesOfSums(sums) # Map of sums onto occurrences of those sums

    print(coords)
    print(sets)
    print(sums)
    print(occurrences)

    
def getOccurrencesOfSums(sums):
    occurrences = {}
    for _sum in sums.values():
        if(not occurrences.get(_sum)): 
            occurrences[_sum] = 1
            continue
        occurrences[_sum] += 1

    return occurrences

def getSetSums(sets):
    sums = {}
    for _set in sets:
        sums[tuple(_set)] = sum(_set)

    return sums

def getSquareCoordsSet(s):
    coords = {}
    for index, row in enumerate(s):
        key = ((0, index), (1, index), (2, index))
        coords[key] = row

    for index, column in enumerate(transverse(s)):
        key = ((index, 0), (index, 1), (index, 2))
        coords[key] = column

    diagonalOne = [s[0][0], s[1][1], s[2][2]]
    diagonalTwo = [s[0][2], s[1][1], s[2][0]]

    key = ((0, 0), (1, 1), (2,2))
    coords[key] = diagonalOne

    key = ((2,2), (1,1), (0,0))
    coords[key] = diagonalTwo

    return coords

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
