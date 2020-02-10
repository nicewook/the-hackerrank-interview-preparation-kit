# Arrays

links: https://www.hackerrank.com/interview/interview-preparation-kit/arrays/challenges

## 2D Array - DS

links: https://www.hackerrank.com/challenges/2d-array

- 2d slices, and there is a subset of it which shape like hourglass
- get the sum of the value of hourglass shape
- Constraints
    - value range is -9 to 9
    - slice size is 6x6
- input
    - 2d slice
- output
    - maxSum
- **point**
    - there could be 16 cases
    - consider i, j or j, i (confusing) https://play.golang.org/p/8K0c6IkiI2T

## Left Rotation

links: https://www.hackerrank.com/challenges/ctci-array-left-rotation

- find valley numbers
- input
    - array a
    - array size n
    - rotation number d
- output
    - rotation result
- **point**
    - just used built-in "append" function

## New Year Chaos

links: https://www.hackerrank.com/challenges/new-year-chaos

- find valley numbers
- input
    - total up/down step number
    - slice of ups and downs
- output
    - number of valleys
- **point**
    - if 'U' and it's sea level, then it was the valley


## Minimum Swaps 2

links: https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

- move forward index + 1 or + 2
- cannot move to slice value is 1. only can move to slice value is 0
- find the shortest move to go to the end
- input
	- total number of clouds
	- slice of binary ints(0 or 1)
- output
	- number of jumps
- **point**
    -  


## Array Manipulation

links: https://www.hackerrank.com/challenges/repeated-string

- given string s, and make long string with repeat s
- count 'a' in len(repeated s string) is n
- input
    - given string s
    - length of repeated string
- output
    - count of 'a' in repeated string that the length of it is n
- **point** 
    - make slice each value of index contain count 'a' from index 0 to i
    - calc string number in length n and multiply total 'a' in string s
    - calc rest part of the string and count 'a'
    - sum and return 
