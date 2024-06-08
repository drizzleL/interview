package main

func largestPalindrome(n int, k int) string {
	return ""
}

// You are given two positive integers n and k.

// An integer x is called k-palindromic if:

//     x is a
//     palindrome
//     .
//     x is divisible by k.

// Return the largest integer having n digits (as a string) that is k-palindromic.

// Note that the integer must not have leading zeros.

// 1: 99999999999999
// 2: 899999998
// 3: 9999999999999
// 4: 888888 xxx 88888888

// 8888888888888888888
// 5: 59999999999995
// 6: 88    mod==1:2
// 7:
// 8: 888888888888888888
// 9: 999999999999999999999

// Example 1:

// Input: n = 3, k = 5

// Output: "595"

// Explanation:

// 595 is the largest k-palindromic integer with 3 digits.

// Example 2:
