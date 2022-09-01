Given an array of integers sorted in ascending order, return the starting and ending index of a given target value in an array, i.e. [x, y].

Your solution should run in O(logN) time.

input arr = [1, 3, 3, 5, 5, 5, 8, 9], target=5
indexes   = [0, 1, 2, 3, 4, 5, 6, 7] 

Return ttarting and ending index of 5: [3, 5]

Constraints:
- What do we return if the target is not found in the array?
  - Return -1, all values in the array are positive.
- 