# Rotate Array

This is a recursive solution. It works, however, there are some optimisations that are possible to get it to pass with the correct time complexity.

Firstly, if the number of rotations is greater than the length of the array, then the required rotations is just `k % len(nums)`.

This is because rotating `len(nums)` times will just result in the same array.

It is also possible to just do the copy step with all the numbers in one swoop, which obviates the need to make the recursive call.




