package main

func grayCode(n int) []int {
	seen := make([]bool, 1<<n)
	seen[0] = true
	var ret []int
	var helper func(now []int) bool
	helper = func(now []int) bool {
		if len(now) == len(seen) {
			ret = now
			return true
		}
		val := now[len(now)-1]
		for i := 0; i < n; i++ {
			flipped := (1 << i) ^ val
			if seen[flipped] {
				continue
			}
			seen[flipped] = true
			if helper(append(now, flipped)) {
				return true
			}
			seen[flipped] = false
		}
		return false
	}
	helper([]int{0})
	return ret
}
