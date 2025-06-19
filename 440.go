package main

func findKthNumber0(n int, k int) int {
	calSteps := func(n, n1, n2 int) int {
		steps := 0
		for n1 <= n {
			steps += min(n+1, n2) - n1
			n1 *= 10
			n2 *= 10
		}
		return steps
	}
	curr := 1
	k = k - 1
	for k > 0 {
		steps := calSteps(n, curr, curr+1)
		if steps <= k {
			curr += 1
			k -= steps
		} else {
			curr *= 10
			k -= 1
		}
	}
	return curr
}
