package main

func kthFactor(n int, k int) int {
	divide := func(x int) []int {
		ret := []int{1}
		for i := 2; i*i <= x; i++ {
			if x%i != 0 {
				continue
			}
			ret = append(ret, i)
		}
		for i := len(ret) - 1; i >= 0; i-- {
			v := x / i
			if i == v {
				continue
			}
			ret = append(ret, v)
		}
		return ret
	}
	vals := divide(n)
	if len(vals) >= k {
		return vals[k-1]
	}
	return -1
}
