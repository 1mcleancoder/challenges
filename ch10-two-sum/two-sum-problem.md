Given an array of integers, return the indices of the two numbers that add up to a given target.

[1, 3, 7, 9, 2] 11
Here 9 and 2 add to the target value 11

Indexes of 9 and 2 are:

Given:
    Array:              [1, 3,  7,  9,  2]
            Indices:    0   1   2   3   4
    11

So the function should return the indexes of 9 and 2 i.e. [3, 4]

Constraints and edge cases:
- are all these +ve numbers? Yes, all numbers are +ve
- are all the numbers be unique or there will be duplicates? No duplicates.
- Will the array always contain the numbers whose sum will add upto a given target? if not then what to return in that case? No, return empty array
- Can there be more than two sets of number whose sum will add upto a given target? if yes, then what to return in that case? Yes, return any set of number's indexes
- will the array be always contain > than 2 numbers? if not then what to return in that case ? no, return empty array.