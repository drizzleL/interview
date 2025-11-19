package main

import "strconv"

func permute(n, k int) int {
	res := 1
	for i := 0; i < k; i++ {
		res *= (n - i)
	}
	return res
}

func countSpecialNumbers(n int) int {
	S := strconv.Itoa(n)
	L := len(S)
	count := 0
	for k := 1; k < L; k++ {
		count += 9 * permute(9, k-1)
	}
	seen := make(map[int]bool)
	for i := 0; i < L; i++ {
		digit := int(S[i] - '0')
		start := 0
		if i == 0 {
			start = 1
		}
		for j := start; j < digit; j++ {
			if !seen[j] {
				remainingDigits := 10 - (i + 1)
				remainingPositions := L - 1 - i
				count += permute(remainingDigits, remainingPositions)
			}
		}
		if seen[digit] {
			return count
		}
		seen[digit] = true
	}
	count++
	return count
}
