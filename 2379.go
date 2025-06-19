package main

func minimumRecolors(blocks string, k int) int {
	var diff int
	var ret int
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			diff += 1
		}
	}
	ret = diff
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			diff += 1
		}
		if blocks[i-k] == 'W' {
			diff -= 1
		}
		ret = min(ret, diff)
	}
	return ret
}
