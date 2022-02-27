#!/usr/bin/env python3

def getSubArray(arr, startX, startY):
    subArray = [[0,0,0],[0,0,0],[0,0,0]]
    for y in range(startY, startY + 3):
            for x in range(startX, startX + 3):
                subArray[y-startY][x-startX] = arr[y][x]
    return subArray

def hourglassSumForSubarray(arr):
    # Takes a 3x3 array
    # Returns the hourglass sum
    sum = 0
    for x in range(3):
        for y in range(3):
            if(y == 1 and x == 0 or y == 1 and x == 2):
                continue
            sum += arr[y][x]

    return sum

def hourglassSum(arr):
    WIDTH = 6
    HEIGHT = 6

    subArrays = []
    # Get a list of all the arrays
    # that form the hourglasses
    startX = 0
    startY = 0
    while(startX <= 3):
        startY = 0
        while(startY <= 3):
            subArray = getSubArray(arr, startX, startY)
            subArrays.append(subArray)
            startY += 1
        startX += 1

    maxHourGlassSum = None
    for subArray in subArrays:
        hourGlassResult = hourglassSumForSubarray(subArray)
        print("Subarray = {}".format(str(subArray)))
        print("Result = {}".format(hourGlassResult))
        if(not maxHourGlassSum and not maxHourGlassSum == 0):
            maxHourGlassSum = hourGlassResult

        if(hourGlassResult > maxHourGlassSum):
            maxHourGlassSum = hourGlassResult

    return maxHourGlassSum




arr = [[1,1,1,0,0,0], [0,1,0,0,0,0], [1,1,1,0,0,0],
       [0,0,2,4,4,0], [0,0,0,2,0,0], [0,0,1,2,4,0]]


arr =   [[-1, -1,  0, -9, -2, -2],
        [-2, -1, -6, -8, -2, -5],
        [-1, -1, -1, -2, -3, -4],
        [-1, -9, -2, -4, -4, -5],
        [-7, -3, -3, -2, -9, -9],
        [-1, -3, -1, -2, -4, -5]]

arr =  [[-1, 1, -1, 0, 0, 0],
        [0, -1,  0, 0, 0, 0],
        [-1, -1, -1, 0, 0, 0],
        [0, -9, 2, -4, -4, 0],
        [-7, 0, 0, -2, 0, 0],
        [0, 0, -1, -2, -4, 0]]


maxHourGlassSum = hourglassSum(arr)
print(maxHourGlassSum)
