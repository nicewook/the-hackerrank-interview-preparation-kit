# Warm-up Challenges

links: https://www.hackerrank.com/interview/interview-preparation-kit/warmup/challenges

## Sock Merchant

links: https://www.hackerrank.com/challenges/sock-merchant

- find how many pairs of socks in there
- input
    - totoal sock number
    - slice of socks (int represent color of socke)
- output
    - total pairs of the right socks
- **point**
    - when meet the color first time, set map true
    - the second time, set map false and count increase


## Counting Valleys

links: https://www.hackerrank.com/challenges/counting-valleys

- find valley numbers
- input
    - total up/down step number
    - slice of ups and downs
- output
    - number of valleys
- **point**
    - if 'U' and it's sea level, then it was the valley


## Jumping on the Clouds

links: https://www.hackerrank.com/challenges/jumping-on-the-clouds

- move forward index + 1 or + 2
- cannot move to slice value is 1. only can move to slice value is 0
- find the shortest move to go to the end
- input
	- total number of clouds
	- slice of binary ints(0 or 1)
- output
	- number of jumps
- **point**
	- trying to go +2 
	- if not possible, then go +1

## Repeating String

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
