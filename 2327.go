package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
	curr := 1
	add := make([]int, delay+1)
	del := make([]int, forget+1)
	add[delay] = 1
	del[forget] = 1
	for i := 1; i < n; i++ {
		for i := 1; i < len(del); i++ {
			del[i-1] = del[i]
		}
		curr -= del[0]
		if curr < 0 {
			curr += 1e9 + 7
		}
		add[0] -= del[0]
		if add[0] < 0 {
			add[0] += 1e9 + 7
		}
		tmp := add[0]
		for i := 1; i < len(add); i++ {
			add[i-1] = add[i]
		}
		add[0] += tmp
		add[0] %= 1e9 + 7
		del[forget] = add[0]
		curr += add[0]
		curr %= 1e9 + 7
		add[delay] = add[0]
	}
	return curr
}
