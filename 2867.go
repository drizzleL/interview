package main

func countPaths(n int, edges [][]int) int64 {
	sieve := make([]bool, n+1)
	sieve[0] = true
	sieve[1] = true
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		for j := 2; i*j <= n; j++ {
			sieve[i*j] = true
		}
	}
	dict := make([][]int, n+1)
	for _, edge := range edges {
		if sieve[edge[1]] {
			dict[edge[0]] = append(dict[edge[0]], edge[1])
		}
		if sieve[edge[0]] {
			dict[edge[1]] = append(dict[edge[1]], edge[0])
		}
	}
	members := map[int]int{}
	cache := map[[2]int]int{}
	var helper func(v int, lastVal int) int
	helper = func(v int, lastVal int) (ret int) {
		if c, ok := cache[[2]int{v, lastVal}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{v, lastVal}] = ret
		}()
		ret = 1
		for _, next := range dict[v] {
			if next == lastVal || !sieve[next] {
				continue
			}
			ret += helper(next, v)
		}
		return ret
	}
	for i, v := range sieve {
		if v {
			continue
		}
		for _, next := range dict[i] {
			members[next] = helper(next, 0)
		}
	}
	var ret int
	for i, v := range sieve {
		if v {
			continue
		}
		var mems []int
		for _, next := range dict[i] {
			if members[next] == 0 {
				continue
			}
			mems = append(mems, members[next])
		}
		var sum int
		for m := 0; m < len(mems); m++ {
			ret += sum * mems[m]
			sum += mems[m]
		}
		ret += sum
	}
	return int64(ret)
}
