## Trapping Rain Water

#### Problem

Given an array of positive integers representing elevations on a 2 dimensional map calculate how many units of rain water can be trapped.

#### Example

Given array [ 0, 2, 0, 3, 1, 0, 2, 1, 0 ] the total units of water that can be trapped is 5. 


                         _
                     _  | |    _
                    | | | |_  | |_
                   _| |_|   |_|   |_
    
    elevations = [ 0 2 0 3 1 0 2 1 0 ]
    
#### Exercise

Write an algorithm that accepts an array (slice, list, etc.) of elevations as an argument and returns an integer representing the units of water that can be trapped. 

###### Case #1

    elevations = [ 1 0 1 ]
    answer: 1

###### Case #2

    elevations = [ 0 2 0 3 1 0 2 1 0 ]
    answer: 5

###### Case #3

    elevations = [ 1, 0, 2, 2, 4, 0, 1, 5, 2, 1, 6, 1 ]
    answer: 15
