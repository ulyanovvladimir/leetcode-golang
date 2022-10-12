package main

import (
	"strconv"
)

/*
https://leetcode.com/problems/palindrome-number/

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.


Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

*/

func isPalindrome(x int) bool {
	var chars = []rune(strconv.Itoa(x))
	n := len(chars)
	for i := 0; i < n && i <= n-i-1; i++ {
		if chars[i] != chars[n-i-1] {
			return false
		}
	}
	return true
}
