#!/usr/bin/env python3

def getBinaryStrings(length):
  """Returns all binary strings of particular length"""
  print("Length is {}".format(length))
  strings = []
  if(length == 1):
    return ['0', '1']
  for s in getBinaryStrings(length-1):
    strings.append("0" + s)
    strings.append("1" + s)
  return sorted(strings)

def getAllBinaryStrings(n):
  """Returns all binary strings of length less than n"""
  strings = []
  for l in range(1, n+1):
    print("l is {}".format(l))
    for string in getBinaryStrings(l):
      strings.append(string)
  return strings

s = getAllBinaryStrings(5)
print(s)
