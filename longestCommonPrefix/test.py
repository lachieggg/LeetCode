#!/usr/bin/env python3

# Test code for breaking out of a 
# meta loop in Python
# 
for s in range(1,5):
	for j in range(3,5):
		print(str(s) + " " + str(j))
		if(s == j):
			# Break out of meta loop
			break
	else:
		continue
	break
