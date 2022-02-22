#!/usr/bin/env python3

# Hourglass problem

def getSubArray(arr, startX, startY):
    # Compute the sub array from given
    # starting coordinates
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
    # Compute the maximum hour glass sum
    # from a 6x6 array

    subArrays = []
    # Try to get a list of all the 3x3 arrays
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

    print(subArrays)

    maxHourGlassSum = 0
    for subArray in subArrays:
        hourGlassResult = hourglassSumForSubarray(subArray)
        if(hourGlassResult > maxHourGlassSum):
            maxHourGlassSum = hourGlassResult

    print(maxHourGlassSum)




arr = [[1,1,1,0,0,0], [0,1,0,0,0,0], [1,1,1,0,0,0],
       [0,0,2,4,4,0], [0,0,0,2,0,0], [0,0,1,2,4,0]]

hourglassSum(arr)
