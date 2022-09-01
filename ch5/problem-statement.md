5. Longest Palindromic Substring
   Medium

16324

961

Add to List

Share
Given a string s, return the longest palindromic substring in s.



Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
Accepted
1,718,947
Submissions
5,423,583
Seen this question in a real interview before?

Yes

No
0 ~ 6 months6 months ~ 1 year1 year ~ 2 years

Amazon
|
32

Microsoft
|
18

Adobe
|
12

Google
|
6

Apple
|
6

Facebook
|
5

Oracle
|
5

tiktok
|
5

Bloomberg
|
4

ByteDance
|
4

Infosys
|
3

PayTM
|
3

Expedia
|
2

Uber
|
2

LinkedIn
|
2

Goldman Sachs
|
2

Salesforce
|
2

Wayfair
|
2

Samsung
|
2

Intel
|
2
How can we reuse a previously computed palindrome to compute a larger palindrome?
If “aba” is a palindrome, is “xabax” a palindrome? Similarly is “xabay” a palindrome?
Complexity based hint:
If we use brute-force and check whether for every start and end position a substring is a palindrome we have O(n^2) start - end pairs and O(n) palindromic checks. Can we reduce the time for palindromic checks to O(1) by reusing some previous computation.