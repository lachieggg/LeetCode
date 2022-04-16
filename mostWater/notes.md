Container With Most Water

[Link to problem](https://leetcode.com/problems/container-with-most-water/)


The water contained by two sides of a container $s1$ and $s2$ is equal
to the difference in the values of the indexes of both sides
multiplied by the minimum height out of both sides.

Define $h(s)$ to be the height of a particular side.

Idea: first define a function to compute the water contained by two sides. Pairs of columns are what we are looking for. The height of the highest column is insignificant, the important value is the height of the lowest column.

Perhaps we can go through each column, and then find the highest column that the column could be matched with, but then we need to consider that at each moment the width is increasing.

At each column, the amount of water grows in proportion to the minimum of the two columns.